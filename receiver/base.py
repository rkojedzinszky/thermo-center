
import logging
import grpc
from lib import aiothread
from lib import cc1101
from receiver import (
        api_pb2, api_pb2_grpc,
        gpio, config,
)
from configurator import api_pb2 as cfg_msg, api_pb2_grpc as cfg_grpc


logger = logging.getLogger(__name__)


class Base(aiothread.AIOThread):
    """ Base class for radio tasks """
    def __init__(self, args, radio):
        super().__init__()

        self.args = args
        self.radio = radio
        self.grpcserver_channel = None
        self.configurator = None

    def init(self):
        """ Set up gpio interrupt handler """
        self.gpio = gpio.InterruptHandler(loop=self.loop, gpiopath=self.args.gpio_dir)

        self.grpcserver_channel = grpc.insecure_channel('{}:{}'.format(self.args.grpcserver_host, self.args.grpcserver_port),
                (
                    ('grpc.keepalive_time_ms', 10000),
                    ('grpc.keepalive_timeout_ms', 1000),
                )
            )
        self.configurator = cfg_grpc.ConfiguratorStub(self.grpcserver_channel)

    def deinit(self):
        self.configurator = None
        self.grpcserver_channel.close()

    def _read_config(self):
        """ Read configuration from configurator """
        return config.Config(self.configurator.GetRadioCfg(cfg_msg.RadioCfgRequest(cluster=1)))

    def __str__(self):
        return self.name
