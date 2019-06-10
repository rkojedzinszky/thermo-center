""" Configurator services """

from django.db import transaction
from django.utils.timezone import now
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

    def TaskAcquire(self, request, context):
        with transaction.atomic():
            task = models.ConfigureSensorTask.objects \
                .select_related('sensor') \
                .select_for_update() \
                .get(pk=request.task_id, started__isnull=True)

            task.started = now()
            task.save()

            return api_pb2.TaskDetails(
                    task_id=task.id,
                    sensor_id=task.sensor.id,
                    config=self.GetRadioCfg(api_pb2.RadioCfgRequest(cluster=1), context),
                    )

    def TaskDiscoveryReceived(self, request, context):
        with transaction.atomic():
            task = models.ConfigureSensorTask.objects \
                .select_for_update() \
                .get(pk=request.task_id, started__isnull=False, finished__isnull=True)

            task.last_discovery = now()
            if task.first_discovery is None:
                task.first_discovery = task.last_discovery

            task.save()

            return api_pb2.TaskUpdateResponse(success=True)

    def TaskFinished(self, request, context):
        with transaction.atomic():
            task = models.ConfigureSensorTask.objects \
                .select_related('sensor') \
                .select_for_update() \
                .get(pk=request.task_id, started__isnull=False, finished__isnull=True)

            task.finished = now()
            task.error = request.error

            task.save()

            if request.error == '':
                task.sensor.resync()

            return api_pb2.TaskUpdateResponse(success=True)


def add_services(server):
    api_pb2_grpc.add_ConfiguratorServicer_to_server(Configurator(), server)
