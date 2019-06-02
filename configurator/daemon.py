""" Configurator services """

from center import models
from configurator import api_pb2
from configurator import api_pb2_grpc

class Configurator(api_pb2_grpc.ConfiguratorServicer):
    def GetRadioCfg(self, request, context):
        if request.cluster != 1:
            raise RuntimeError("Invalid cluster ID received")

        config = models.RFConfig.objects.select_related('rf_profile').get(pk=request.cluster)

        return api_pb2.RadioCfgResponse(
                network=config.network_id,
                radio_config=config.config_bytes(),
                aes_key=bytes.fromhex(config.aes_key)
                )


def add_services(server):
    api_pb2_grpc.add_ConfiguratorServicer_to_server(Configurator(), server)
