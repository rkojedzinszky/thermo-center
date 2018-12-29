
import re
import asyncio
import websockets
import logging
from django.conf import settings
from hbmqtt.client import MQTTClient, ClientException, QOS_0
from django.utils.functional import cached_property

logger = logging.getLogger(__name__)

class MqttClient:
    T_RE = re.compile(settings.MQTT_PREFIX + '(.*?)/report')

    def __init__(self, loop, wsserver):
        self.loop = loop
        self.wsserver = wsserver

    async def _run(self):
        cl = MQTTClient(config={'auto_reconnect':False})
        logger.debug('Connecting to MQTT broker')
        await cl.connect('mqtt://{}:{}'.format(settings.MQTT_HOST, settings.MQTT_PORT))
        logger.info('Connected to MQTT broker')
        await cl.subscribe([(settings.MQTT_PREFIX + '+/report', QOS_0)])

        while True:
            message = await cl.deliver_message()
            topic = message.publish_packet.variable_header.topic_name

            m = self.T_RE.match(topic)
            if m:
                h = '{}'.format(int(m.group(1), base=16))
                self.wsserver.report(h)

    async def run(self):
        while True:
            try:
                await self._run()
            except asyncio.CancelledError:
                return
            except Exception as e:
                logger.warn(e, exc_info=True)

            await asyncio.sleep(1)

class WsClient:
    def __init__(self, loop, ws):
        self.loop = loop
        self.ws = ws
        self._task = None
        self._q = asyncio.Queue(maxsize=10, loop=loop)

    async def _loop(self):
        logger.info('%s: starting', self)
        try:
            while True:
                el = await self._q.get()
                logger.debug('%s: sending %s', self, el)
                await self.ws.send(el)

        except Exception as e:
            logger.warn('%s: exiting', self, exc_info=True)

    async def done(self):
        self._task = self.loop.create_task(self._loop())
        await self._task

    def send(self, payload):
        try:
            self._q.put_nowait(payload)
        except asyncio.QueueFull:
            logger.warn('%s: buffer overrun, exiting', self)
            self._task.cancel()

    @cached_property
    def remote(self):
        h = self.ws.request_headers
        return \
                h.get('X-Forwarder-For', None) or \
                h.get('X-Real-IP', None) or \
                self.ws.remote_address[0]

    def __str__(self):
        return self.remote

class Main:
    def __init__(self, loop=None):
        self.loop = loop or asyncio.get_event_loop()
        self.mqtt = MqttClient(self.loop, self)
        self.clients = set()

    async def _handler(self, ws, path):
        cl = WsClient(self.loop, ws)

        self.clients.add(cl)
        try:
            await cl.done()
        finally:
            self.clients.remove(cl)

    def report(self, id_):
        for cl in self.clients:
            cl.send(id_)

    async def run(self):
        self.loop.create_task(self.mqtt.run())
        await websockets.serve(self._handler, '0.0.0.0', 8081)
