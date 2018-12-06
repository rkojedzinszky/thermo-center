""" API for center """

import datetime
import time
from django.utils import timezone
from django.core.cache import cache
from application import restapi
from application.resource import ResourceMetaCommon
from center.models import Sensor, SensorResync
from tastypie import resources
from tastypie.authentication import Authentication
from tastypie.authorization import ReadOnlyAuthorization, Authorization as RWAuthorization
from tastypie import fields

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
        bundle._cache = cache.get(bundle.obj._carbon_path())
        if bundle._cache:
            bundle.data['valid'] = bundle._cache.get('valid', None)
            bundle.data['vcc'] = bundle._cache.get('Power', None)
            bundle.data['rssi'] = bundle._cache.get('RSSI', None)
            bundle.data['lqi'] = bundle._cache.get('LQI', None)
            bundle.data['interval'] = bundle._cache.get('intvl', None)
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

