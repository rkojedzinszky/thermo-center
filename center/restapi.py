""" API for center """

import datetime
import time
import grpc
from django.utils import timezone
from django.dispatch import receiver
from django.db.models.signals import post_save
from django.core.cache import cache
from django.conf import settings
from application import restapi
from application.resource import ResourceMetaCommon
from center.models import Sensor, SensorResync, ConfigureSensorTask
from tastypie import resources
from tastypie.authentication import Authentication
from tastypie.authorization import ReadOnlyAuthorization, Authorization as RWAuthorization
from tastypie import fields
from receiver import api_pb2_grpc
from configurator import api_pb2 as cfg_pb2

class SensorResource(resources.ModelResource):
    valid = fields.BooleanField(null=True, readonly=True, help_text='Recent update status')
    vcc = fields.FloatField(null=True, help_text='Power in battery in sensor')
    rssi = fields.FloatField(null=True, help_text='RSSI of sensor')
    lqi = fields.FloatField(null=True, help_text='Line Quality Indicator of sensor')
    interval = fields.FloatField(null=True, help_text='Elapsed time since last report')
    age = fields.FloatField(null=True, readonly=True)
    server_time = fields.DateTimeField(readonly=True)

    sensor_resync = fields.ForeignKey('center.restapi.SensorResyncResource', '', readonly=True, null=True)
    thsensor = fields.ForeignKey('center.restapi.THSensorResource', '', readonly=True, null=True)

    class Meta(ResourceMetaCommon):
        queryset = Sensor.objects.all()
        authorization = ReadOnlyAuthorization()
        excludes = ('last_seq',)
        ordering = (
                'id',
                )

    def dehydrate(self, bundle):
        bundle = super(SensorResource, self).dehydrate(bundle)
        now = timezone.now()

        bundle.data['thsensor'] = THSensorResourceInstance.get_resource_uri(bundle.obj)
        bundle._cache = bundle.obj.get_cache()
        if bundle._cache:
            bundle.data['valid'] = bundle._cache.get('valid', None)
            bundle.data['vcc'] = bundle._cache.get('Power', None)
            bundle.data['rssi'] = bundle._cache.get('RSSI', None)
            bundle.data['lqi'] = bundle._cache.get('LQI', None)
            bundle.data['interval'] = bundle._cache.get('intvl', None)
            bundle.data['last_tsf'] = bundle._cache.get('last_tsf', None)
            if bundle.data['valid'] == False:
                bundle.data['sensor_resync'] = SensorResyncResourceInstance.get_resource_uri(bundle.obj)

        bundle.data['server_time'] = now
        if bundle.obj.last_tsf:
            bundle.data['age'] = time.time() - bundle.obj.last_tsf

        return bundle

class THSensorResource(SensorResource):
    temperature = fields.FloatField(null=True)
    humidity = fields.FloatField(null=True)

    class Meta(SensorResource.Meta):
        excludes = ('thsensor',)

    def dehydrate(self, bundle):
        bundle = super(THSensorResource, self).dehydrate(bundle)

        if bundle._cache:
            bundle.data['temperature'] = bundle._cache.get('Temperature', None)
            bundle.data['humidity'] = bundle._cache.get('Humidity', None)

        return bundle

THSensorResourceInstance = THSensorResource()

restapi.RestApi.register(SensorResource())
restapi.RestApi.register(THSensorResourceInstance)

class SensorResyncResource(resources.ModelResource):
    sensor = fields.ToOneField(THSensorResource, 'sensor')
    ts = fields.DateTimeField('ts', readonly=True)

    class Meta(ResourceMetaCommon):
        queryset = SensorResync.objects.all()
        authorization = RWAuthorization()
        detail_allowed_methods = ['get']
        list_allowed_methods = ['get', 'post']

SensorResyncResourceInstance = SensorResyncResource()
restapi.RestApi.register(SensorResyncResourceInstance)

class ConfigureSensorTaskResource(resources.ModelResource):
    sensor_id = fields.IntegerField()
    sensor_name = fields.CharField()
    created = fields.DateTimeField('created', null=True, readonly=True)
    started = fields.DateTimeField('started', null=True, readonly=True)
    first_discovery = fields.DateTimeField('first_discovery', null=True, readonly=True)
    last_discovery = fields.DateTimeField('last_discovery', null=True, readonly=True)
    finished = fields.DateTimeField('finished', null=True, readonly=True)
    error = fields.CharField('error', null=True, readonly=True)

    class Meta(ResourceMetaCommon):
        queryset = ConfigureSensorTask.objects.select_related('sensor')
        authorization = RWAuthorization()
        detail_allowed_methods = ['get']
        list_allowed_methods = ['get', 'post']
        fields = ['id']

    def obj_create(self, bundle, **kwargs):
        if 'sensor_name' in bundle.data:
            sensor_id = (set(range(1, 128)) - set([s.pk for s in Sensor.objects.all()])).pop()
            sensor = Sensor.objects.create(
                    id=sensor_id,
                    name=bundle.data['sensor_name'],
                    )
        else:
            sensor = Sensor.objects.get(pk=bundle.data['sensor_id'])

        return super().obj_create(bundle, sensor=sensor)

    def dehydrate(self, bundle):
        bundle.data['sensor_id'] = bundle.obj.sensor.id
        bundle.data['sensor_name'] = bundle.obj.sensor.name

        return bundle

ConfigureSensorTaskResourceInstance = ConfigureSensorTaskResource()
restapi.RestApi.register(ConfigureSensorTaskResourceInstance)

@receiver(post_save, sender=ConfigureSensorTask)
def _receiver_handletask(sender, instance, created, **kwargs):
    if created:
        channel = grpc.insecure_channel('{}:{}'.format(settings.RECEIVER_HOST, settings.RECEIVER_PORT))
        receiver = api_pb2_grpc.ReceiverStub(channel)
        receiver.HandleTask(cfg_pb2.Task(task_id=instance.id))
        channel.close()
