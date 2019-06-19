
import re
import asyncio
import websockets
import logging
import signal
from hbmqtt.client import MQTTClient, ClientException, QOS_0

logger = logging.getLogger(__name__)

MQTT_PREFIX = 'thsensor/'

class MqttClient:
    T_RE = re.compile(MQTT_PREFIX + '(.*?)/report')

    def __init__(self, loop, wsserver):
        self.loop = loop
        self.wsserver = wsserver

    async def _run(self):
        cl = MQTTClient(config={'auto_reconnect':False})
        logger.debug('Connecting to MQTT broker')
        await cl.connect('mqtt://{}:{}'.format(self.wsserver.args.mqtt_host, self.wsserver.args.mqtt_port))
        logger.info('Connected to MQTT broker')
        await cl.subscribe([(MQTT_PREFIX + '+/report', QOS_0)])

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
                logger.warning(e, exc_info=True)

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
            logger.warning('%s: exiting', self, exc_info=True)

    async def done(self):
        self._task = self.loop.create_task(self._loop())
        await self._task

    def send(self, payload):
        try:
            self._q.put_nowait(payload)
        except asyncio.QueueFull:
            logger.warning('%s: buffer overrun, exiting', self)
            self._task.cancel()

    def remote(self):
        h = self.ws.request_headers
        return \
                h.get('X-Forwarder-For', None) or \
                h.get('X-Real-IP', None) or \
                self.ws.remote_address[0]

    def __str__(self):
        return self.remote()

class Main:
    def __init__(self, args, loop):
        self.args = args
        self.loop = loop
        self.mqtt = MqttClient(self.loop, self)
        self.stopserver = asyncio.Event(loop=loop)
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
        mqtt_task = self.loop.create_task(self.mqtt.run())
        async with websockets.serve(self._handler, '0.0.0.0', args.ws_port):
            await self.stopserver.wait()
        mqtt_task.cancel()
        await mqtt_task

    def shutdown(self, *args, **kwargs):
        self.stopserver.set()


if __name__ == '__main__':
    import argparse
    import os
    import sys

    parser = argparse.ArgumentParser(description='Thermo center websocket daemon')
    parser.add_argument('--mqtt-host', default=os.environ.get('MQTT_HOST', 'mqtt'),
            help='MQTT hostname/address')
    parser.add_argument('--mqtt-port', type=int, default=int(os.environ.get('MQTT_PORT', '1883')),
            help='MQTT port')
    parser.add_argument('--ws-port', type=int, default=int(os.environ.get('WS_PORT', '8081')),
            help='Websocket port')

    logging.basicConfig(stream=sys.stderr, level=logging.INFO)

    args = parser.parse_args()

    loop = asyncio.get_event_loop()

    daemon = Main(args=args, loop=loop)

    signal.signal(signal.SIGTERM, daemon.shutdown)

    loop.run_until_complete(daemon.run())
