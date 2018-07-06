import os, sys
import asyncio
import select
import logging
from center.receiver.mqtt import MqttClient
from center import models
from django.db.backends import signals
from django.conf import settings
import spidev
from center.receiver import radio, gpio

logger = logging.getLogger(__name__)

class ConsoleClient(asyncio.Protocol):
    def __init__(self, main):
        self.main = main

    def connection_made(self, transport):
        self.transport = transport

    def data_received(self, data):
        data = data.decode().strip()

        if data == 'stop':
            self.main.loop.create_task(self.main.stop())
        elif data == 'configure':
            self.main.loop.create_task(self.main.startconfigurator())
            self.main.loop.call_later(1, self.main.startreceiver_sync)
        elif data == 'reload':
            self.main.loop.create_task(self.main.startreceiver())

        self.transport.close()

class Main:
    """ Main radio handler daemon """

    def run(self, daemonize=True):
        spi = spidev.SpiDev()
        spi.open(*settings.SPI_DEV)
        spi.mode = settings.SPI_MODE
        spi.max_speed_hz = settings.SPI_FREQ
        self._radio = radio.Radio(spi)

        partnum = self._radio.status(radio.Radio.StatusReg.PARTNUM)
        if partnum != 0x0:
            raise RuntimeError('CC1101 identification failed: PARTNUM={:x}'.format(partnum))
        version = self._radio.status(radio.Radio.StatusReg.VERSION)
        if version != 0x14:
            raise RuntimeError('CC1101 identification failed: VERSION={:x}'.format(version))

        self._gpio = gpio.GPIO(settings.INT_GPIO_DIR)
        self._gpio.input()

        self._loop = None
        self._pidmap = {}

        if daemonize:
            if os.fork() > 0:
                sys.exit()

            os.chdir('/')
            os.setsid()

            if os.fork() > 0:
                sys.exit()

            with open('/dev/null', 'r') as fh:
                os.dup2(fh.fileno(), sys.stdin.fileno())
            with open('/dev/null', 'a+') as fh:
                os.dup2(fh.fileno(), sys.stdout.fileno())
                os.dup2(fh.fileno(), sys.stderr.fileno())

        signals.connection_created.connect(self._set_sync_commit_to_off)

        self.loop = asyncio.get_event_loop()

        self._mqtt_setup()

        self.start_console()

        self.loop.create_task(self.startreceiver())

        self.loop.run_forever()

        self.console.close()

        self._mqtt_teardown()

    def _mqtt_setup(self):
        if hasattr(settings, 'MQTT_HOST'):
            self._mqtt = MqttClient(self.loop, (settings.MQTT_HOST, settings.MQTT_PORT))
            self._mqtt.start()
        else:
            self._mqtt = None

    def _mqtt_teardown(self):
        if self._mqtt:
            self.loop.run_until_complete(self._mqtt.stop())
            self._mqtt = None

    def start_console(self):
        umask = os.umask(0o077)
        self.console = self.loop.run_until_complete(self.loop.create_unix_server(self.create_console_client, path=settings.RECEIVER_SOCKET))
        os.umask(umask)

    def create_console_client(self):
        return ConsoleClient(self)

    def _set_sync_commit_to_off(self, sender, connection, **kwargs):
        try:
            with connection.cursor() as c:
                c.execute('set synchronous_commit to off')
        except:
            pass

    def startreceiver_sync(self):
        self.loop.create_task(self.startreceiver())

    async def startreceiver(self):
        await self._setloop(receiver.Receiver)
        self._loop.setpidmap(self._pidmap)

    async def startconfigurator(self):
        await self._setloop(configurator.Configurator)

    async def _setloop(self, cls):
        if self._loop:
            await self._loop.stop()
        if cls:
            self._loop = cls(self.loop, self._radio, self._gpio, self._mqtt)
            self._loop.start()
        else:
            self._loop = None

    async def stop(self):
        await self._setloop(None)
        self.loop.stop()

class RadioBase:
    """ Base class for receiver and configurator mode """

    class GPIOInterruptHandler:
        def __init__(self, loop, gpio):
            self._loop = loop
            self._poller = select.epoll()
            self._gpio = gpio
            self._gpio.interrupt('rising')
            self._poller.register(gpio.fileno(), select.POLLPRI)

        def enable(self):
            self._loop.add_reader(self._poller.fileno(), self._onread)

        def disable(self):
            self._loop.remove_reader(self._poller.fileno())

        async def waitforinterrupt(self):
            while True:
                if self._gpio.value():
                    return
                await asyncio.sleep(0)
                self.interrupt = self._loop.create_future()
                self.enable()
                try:
                    await self.interrupt
                finally:
                    self.disable()

        def _onread(self):
            self._poller.poll()
            self.interrupt.set_result(True)

    def __init__(self, loop, radio, gpio, mqtt):
        self.loop = loop
        self._radio = radio
        self._radio.reset()
        self._ih = RadioBase.GPIOInterruptHandler(loop, gpio)
        self._mqtt = mqtt
        self._config = models.RFConfig.objects.select_related('rf_profile').get(pk=1)

    def start(self):
        logger.info('{} starting'.format(self.name))
        self.task = self.loop.create_task(self.main())

    async def stop(self):
        logger.info('{} stopping'.format(self.name))
        self.task.cancel()
        await asyncio.wait([self.task])
        logger.info('{} finished'.format(self.name))

    async def waitforinterrupt(self):
        await self._ih.waitforinterrupt()

from center.receiver import receiver
from center.receiver import configurator
