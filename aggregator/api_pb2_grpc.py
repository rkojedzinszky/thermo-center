# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from aggregator import api_pb2 as aggregator_dot_api__pb2


class AggregatorStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.FeedSensorPacket = channel.unary_unary(
                '/aggregator.Aggregator/FeedSensorPacket',
                request_serializer=aggregator_dot_api__pb2.SensorPacket.SerializeToString,
                response_deserializer=aggregator_dot_api__pb2.FeedResponse.FromString,
                )


class AggregatorServicer(object):
    """Missing associated documentation comment in .proto file."""

    def FeedSensorPacket(self, request, context):
        """Feed packet to processor
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_AggregatorServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'FeedSensorPacket': grpc.unary_unary_rpc_method_handler(
                    servicer.FeedSensorPacket,
                    request_deserializer=aggregator_dot_api__pb2.SensorPacket.FromString,
                    response_serializer=aggregator_dot_api__pb2.FeedResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'aggregator.Aggregator', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Aggregator(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def FeedSensorPacket(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/aggregator.Aggregator/FeedSensorPacket',
            aggregator_dot_api__pb2.SensorPacket.SerializeToString,
            aggregator_dot_api__pb2.FeedResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
