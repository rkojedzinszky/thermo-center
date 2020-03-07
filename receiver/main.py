import sys
import os
import spidev
import logging
import time
import signal
import concurrent.futures
import grpc
from receiver import (
        api_pb2, api_pb2_grpc,
        radio, config,
        receiver,
        configurator,
)

logger = logging.getLogger(__name__)


class ReceiverServicer(api_pb2_grpc.ReceiverServicer):
    def __init__(self, daemon):
        super().__init__()

        self.daemon = daemon

    def HandleTask(self, request, context):
        """ Enter the daemon into sensor configurator mode """
        logger.info('Received configurator request for task %d', request.task_id)
        self.daemon.ConfigureSensor(request)

        return api_pb2.HandleResponse(success=True)

class Daemon:
    """ Main receiver class """
    def __init__(self, args):
        self.args = args
        self.radio = None
        self.gpio = None
        self.task = None  # the currently running task
        self.grpcserver = None

    def _init_radio(self):
        dev = spidev.SpiDev()
        dev.open(self.args.spi_bus, self.args.spi_cs)
        dev.mode = self.args.spi_mode
        dev.max_speed_hz = self.args.spi_freq
        self.radio = radio.Radio(dev)

    def _start_new_task(self, task):
        old = self.task
        self.task = task
        if old:
            old.cancel()

    def shutdown(self, *args, **kwargs):
        logger.warning('shutdown request received')

        task = self.task
        self.task = None
        if task:
            task.cancel()

    def ConfigureSensor(self, task):
        """ Set configurator mode """
        self._start_new_task(configurator.Configurator(args=self.args, radio=self.radio, task=task))

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

    def _start_grpcserver(self):
        self.grpcserver = grpc.server(concurrent.futures.ThreadPoolExecutor(max_workers=1), options=(
            ('grpc.keepalive_time_ms', 60000),
            ('grpc.keepalive_timeout_ms', 1000),
            ('grpc.keepalive_permit_without_calls', True),
            ('grpc.http2.max_pings_without_data', 0),
            ('grpc.http2.min_time_between_pings_ms', 10000),
            ('grpc.http2.min_ping_interval_without_data_ms', 1000))
        )
        api_pb2_grpc.add_ReceiverServicer_to_server(ReceiverServicer(self), self.grpcserver)
        self.grpcserver.add_insecure_port('0.0.0.0:{}'.format(self.args.receiver_port))
        self.grpcserver.start()

    def _stop_grpcserver(self):
        self.grpcserver.stop(None)

    def run(self):
        """ Main entrypoint """

        os.chdir('/')

        if self.args.daemonize:
            self.daemonize()

        self._init_radio()

        self._start_grpcserver()

        task = None
        while True:
            # if self.task has changed, that means that is
            # to be run. Else, schedule a new receiver instance
            if task is self.task:
                self.task = receiver.Receiver(args=self.args, radio=self.radio)

            task = self.task

            logger.info('%s: starting', task)
            task.start()
            task.join()
            logger.info('%s: finished', task)

            if not self.task:  # requested loop end
                break

            time.sleep(1)

        self._stop_grpcserver()


if __name__ == '__main__':
    import argparse
    import os

    parser = argparse.ArgumentParser(description='Thermo center receiver daemon')
    parser.add_argument('--grpcserver-host', default=os.environ.get('GRPCSERVER_HOST', 'grpcserver'),
            help='Grpcserver hostname/address')
    parser.add_argument('--grpcserver-port', type=int, default=int(os.environ.get('GRPCSERVER_PORT', '8079')),
            help='Grpcserver port')
    parser.add_argument('--receiver-port', type=int, default=int(os.environ.get('RECEIVER_PORT', '8079')),
            help='Receiver port')
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

    logging.basicConfig(stream=sys.stderr, level=logging.INFO)

    config = parser.parse_args()

    daemon = Daemon(config)

    signal.signal(signal.SIGTERM, daemon.shutdown)

    daemon.run()
