
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

    def init(self):
        """ Set up gpio interrupt handler """
        self.gpio = gpio.InterruptHandler(loop=self.loop, gpiopath=self.args.gpio_dir)

    def _read_config(self):
        """ Read configuration from configurator """
        channel = grpc.insecure_channel('{}:{}'.format(self.args.configurator_host, self.args.configurator_port),
                                        (
                    ('grpc.keepalive_time_ms', 10000),
                    ('grpc.keepalive_timeout_ms', 1000),
                )
            )
        stub = cfg_grpc.ConfiguratorStub(channel)

        cfg = config.Config(
            stub.GetRadioCfg(cfg_msg.RadioCfgRequest(cluster=1)))

        channel.close()

        return cfg

    def __str__(self):
        return self.name
