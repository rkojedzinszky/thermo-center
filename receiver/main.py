import sys
import os
import spidev
import logging
import time
from . import (
    radio, config, receiver
)

logger = logging.getLogger(__name__)

class Daemon:
    """ Main receiver class """
    def __init__(self, args):
        self.args = args
        self.radio = None
        self.gpio = None
        self.task = None
        self.task_class = receiver.Receiver

    def _init_radio(self):
        dev = spidev.SpiDev()
        dev.open(self.args.spi_bus, self.args.spi_cs)
        dev.mode = self.args.spi_mode
        dev.max_speed_hz = self.args.spi_freq
        self.radio = radio.Radio(dev)

    async def _start_receiver(self):
        """ Set receive mode """
        await self._set_radio_task(receiver.Receiver)

    async def _start_discoverer(self):
        """ Set discover mode """
        pass

    async def _set_radio_task(self, cls):
        if self.radio_task:
            await self.radio_task.stop()
        if cls:
            self.radio_task = cls(
                    loop=self.loop,
                    config=self.config,
                    radio=self.radio,
                    gpio=self.gpio,
                    mqtt=self.mqtt,
                    )
            await self.radio_task.start()
        else:
            self.radio_task = None

    def daemonize(self):
        """ Become a daemon """
        pid = os.fork()
        if pid > 0:
            sys.exit(0)

        devnullfd = os.open('/dev/null', os.RDWR)
        os.dup2(devnullfd, sys.stdin.fileno())
        os.dup2(devnullfd, sys.stdout.fileno())
        os.dup2(devnullfd, sys.stderr.fileno())
        os.close(devnullfd)

    def run(self):
        """ Main entrypoint """

        os.chdir('/')

        if self.args.daemonize:
            self.daemonize()

        self._init_radio()

        while True:
            self.task = self.task_class(args=self.args, radio=self.radio)
            logger.info('%s: starting', self.task)
            self.task.start()
            self.task.join()
            logger.info('%s: finished', self.task)
            time.sleep(1)


if __name__ == '__main__':
    import argparse
    import os

    parser = argparse.ArgumentParser(description='Thermo center receiver daemon')
    parser.add_argument('--aggregator-host', default=os.environ.get('AGGREGATOR_HOST', 'aggregator'),
            help='Aggregator hostname/address')
    parser.add_argument('--aggregator-port', type=int, default=int(os.environ.get('AGGREGATOR_PORT', '8079')),
            help='Aggregator port')
    parser.add_argument('--configurator-host', default=os.environ.get('CONFIGURATOR_HOST', 'configurator'),
            help='Configurator hostname/address')
    parser.add_argument('--configurator-port', type=int, default=int(os.environ.get('CONFIGURATOR_PORT', '8079')),
            help='Configurator port')
    parser.add_argument('--spi-bus', type=int, default=int(os.environ.get('SPI_BUS_NUM', '0')),
            help='SPI Bus number')
    parser.add_argument('--spi-cs', type=int, default=int(os.environ.get('SPI_CS_NUM', '0')),
            help='SPI Chip-select number')
    parser.add_argument('--spi-mode', type=int, default=int(os.environ.get('SPI_MODE', '0')),
            help='SPI mode')
    parser.add_argument('--spi-freq', type=int, default=int(os.environ.get('SPI_FREQ', '100000')),
            help='SPI frequency')
    parser.add_argument('--gpio-dir', default=os.environ.get('GPIO_DIR', '/gpio'),
            help='GPIO dir for interrupt')
    parser.add_argument('--daemonize', type=bool, default=False,
            help='Become a daemon')

    logging.basicConfig(stream=sys.stderr, level=logging.WARNING)

    logging.getLogger('receiver').setLevel(logging.INFO)

    config = parser.parse_args()

    Daemon(config).run()
