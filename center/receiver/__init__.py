import os, sys
import select
from twisted.internet import reactor, interfaces, protocol
from twisted.protocols import basic
from zope.interface import implementer
from center import models

from django.conf import settings
import spidev
from center.receiver import radio, gpio

class Console(basic.LineOnlyReceiver):
    def setMain(self, main):
        self._main = main
        return self

    def lineReceived(self, line):
        if line == 'stop':
            self.sendLine('exiting')
            self._main.stop()
        elif line == 'configure':
            self._main.startconfigurator()
            self.sendLine('entered sensor configuration mode')
        elif line == 'reload':
            self._main.startreceiver()
            self.sendLine('reloaded receiver mode')

class ConsoleFactory(protocol.ServerFactory):
    def setMain(self, main):
        self._main = main
        return self

    def buildProtocol(self, addr):
        return Console().setMain(self._main)

class Main(object):
    """ Main radio handler daemon """

    pidfile = '%s/receiver.pid' % settings.BASE_DIR

    def run(self, daemonize=True):
        spi = spidev.SpiDev()
        spi.open(settings.SPI_BUS, 0)
        spi.mode = 3
        spi.max_speed_hz = 1000000
        self._radio = radio.Radio(spi)

        self._gpio = gpio.GPIO(settings.GPIO_SYS_DIR)
        self._gpio.input()

        self._loop = None

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

        self._listen = reactor.listenTCP(1234, ConsoleFactory().setMain(self))

        self.startreceiver()

        reactor.run()

    def startreceiver(self):
        from center.receiver import receiver
        self._setloop(receiver.Receiver)

    def startconfigurator(self):
        from center.receiver import configurator
        self._setloop(configurator.Configurator)

    def _setloop(self, cls):
        if self._loop:
            self._loop.finish()
        if cls:
            self._loop = cls(self._radio, self._gpio)
            self._loop._main = self
            self._loop.run()
        else:
            self._loop = None

    def stop(self):
        self._listen.stopListening()
        reactor.stop()

class RadioBase(object):
    """ Base class for receiver and configurator mode """

    @implementer(interfaces.IReadDescriptor)
    class GPIOInterruptHandler(object):
        def __init__(self, gpio, cb):
            self._poller = select.epoll()
            self._gpio = gpio
            self._gpio.interrupt('rising')
            self._poller.register(gpio.fileno(), select.POLLPRI)
            self._cb = cb

        def fileno(self):
            return self._poller.fileno()

        def doRead(self):
            self._poller.poll()
            while self._gpio.value():
                self._cb()

        def logPrefix(self):
            return 'GPIOInterruptHandler'

        def connectionLost(self, reason):
            pass

    def __init__(self, radio, gpio):
        self._radio = radio

        self._radio.reset()
        self._ih = RadioBase.GPIOInterruptHandler(gpio, self.oninterrupt)
        self._config = models.RFConfig.objects.select_related('rf_profile').get(pk=1)

    def oninterrupt(self, cb):
        raise NotImplementedError()

    def enable_interrupt(self):
        reactor.addReader(self._ih)

    def disable_interrupt(self):
        reactor.removeReader(self._ih)

    def finish(self):
        self._radio.sidle()
        self.disable_interrupt()
        del self._ih
