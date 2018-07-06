import asyncio
import logging
from hbmqtt.client import MQTTClient, ClientException

logger = logging.getLogger(__name__)

class DictQueue:
    """ A dictionary that also works as a queue """
    def __init__(self, loop):
        self._dict = dict()
        self._loop = loop
        self._has_elems = asyncio.Event(loop=loop)

    def put(self, key, data):
        self._dict[key] = data
        self._has_elems.set()

    async def get(self):
        while True:
            if len(self._dict) > 0:
                return self._dict.popitem()
            self._has_elems.clear()
            await self._has_elems.wait()

class MqttClient:
    """ MQTT client which reconnects infinitely, and publishes valid updates to topics """
    def __init__(self, loop, address, validity=1.0):
        self._loop = loop
        self._address = address
        self._validity = validity
        self._dq = DictQueue(loop)
        self._task = None

    async def _get(self):
        while True:
            key, data = await self._dq.get()
            expiry, payload = data
            if self._loop.time() <= expiry:
                return key, payload

    async def _main(self):
        while True:
            try:
                cl = MQTTClient(config={'auto_reconnect':False})
                logger.debug('Connecting to MQTT broker')
                await cl.connect('mqtt://{}:{}'.format(*self._address))
                logger.info('Connected to MQTT broker')

                while True:
                    topic, payload = await self._get()
                    await cl.publish(topic, payload)

            except asyncio.CancelledError:
                break
            except ClientException as e:
                logger.warn('MQTT Exception: {}'.format(e))

            try:
                await asyncio.sleep(1)
            except asyncio.CancelledError:
                break

        try:
            await cl.disconnect()
        except ClientException:
            pass

    def start(self):
        self._task = self._loop.create_task(self._main())

    async def stop(self):
        self._task.cancel()
        await asyncio.wait_for(self._task, None, loop=self._loop)

    def publish(self, topic, payload, validity=None):
        expiry = self._loop.time() + (validity or self._validity)
        self._dq.put(topic, (expiry, payload))

