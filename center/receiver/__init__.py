import os, sys
import asyncio
import select
import logging
import paho.mqtt.client as mqtt
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
            self.main.stop()
        elif data == 'configure':
            self.main.startconfigurator()
            self.main.loop.call_later(15, self.main.startreceiver)
        elif data == 'reload':
            self.main.startreceiver()

        self.transport.close()

class Main(object):
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

        self._mqtt_setup()

        self.loop = asyncio.get_event_loop()

        self.start_console()

        self.startreceiver()

        self.loop.run_forever()

        self.console.close()

        self._mqtt_teardown()

    def _mqtt_setup(self):
        if hasattr(settings, 'MQTT_HOST'):
            self._mqtt = mqtt.Client()
            self._mqtt.loop_start()
            self._mqtt.connect(settings.MQTT_HOST, settings.MQTT_PORT)
        else:
            self._mqtt = None

    def _mqtt_teardown(self):
        if hasattr(self, '_mqtt') and self._mqtt:
            self._mqtt.loop_stop()

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

    def startreceiver(self):
        self._setloop(receiver.Receiver)
        self._loop.setpidmap(self._pidmap)

    def startconfigurator(self):
        self._setloop(configurator.Configurator)

    def _setloop(self, cls):
        if self._loop:
            self._loop._stop()
        if cls:
            self._loop = cls(self.loop, self._radio, self._gpio, self._mqtt)
            self._loop._start()
        else:
            self._loop = None

    def stop(self):
        self._setloop(None)
        self.loop.stop()

class RadioBase(object):
    """ Base class for receiver and configurator mode """

    class GPIOInterruptHandler(object):
        def __init__(self, gpio, cb):
            self._poller = select.epoll()
            self._gpio = gpio
            self._gpio.interrupt('rising')
            self._poller.register(gpio.fileno(), select.POLLPRI)
            self._cb = cb

        def enable(self):
            asyncio.get_event_loop().add_reader(self._poller.fileno(), self._onread)

        def disable(self):
            asyncio.get_event_loop().remove_reader(self._poller.fileno())

        def _onread(self):
            self._poller.poll()
            while self._gpio.value():
                self._cb()

    def __init__(self, loop, radio, gpio, mqtt):
        self.loop = loop
        self._radio = radio
        self._radio.reset()
        self._ih = RadioBase.GPIOInterruptHandler(gpio, self.oninterrupt)
        self._mqtt = mqtt
        self._config = models.RFConfig.objects.select_related('rf_profile').get(pk=1)

    def oninterrupt(self, cb):
        raise NotImplementedError()

    def enable_interrupt(self):
        self._ih.enable()

    def disable_interrupt(self):
        self._ih.disable()

    def _start(self):
        logger.info('%s starting' % self.name)
        self.start()
        logger.info('%s started' % self.name)

    def start(self):
        """ Override this in subclasses """

    def _stop(self):
        logger.info('%s stopping' % self.name)
        self._radio.sidle()
        self.disable_interrupt()
        del self._ih
        self.stop()
        logger.info('%s stopped' % self.name)

    def stop(self):
        """ Override this in subclasses """

from center.receiver import receiver
from center.receiver import configurator
