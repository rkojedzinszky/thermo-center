""" API for center """

from django.core.cache import cache
from application import restapi, resource
from center.models import Sensor
from tastypie.authentication import Authentication
from tastypie.authorization import ReadOnlyAuthorization
from tastypie import fields

class SensorResource(resource.ModelResource):
    vcc = fields.FloatField(null=True, help_text='Power in battery in sensor')
    rssi = fields.FloatField(null=True, help_text='RSSI of sensor')
    lqi = fields.FloatField(null=True, help_text='Line Quality Indicator of sensor')
    interval = fields.FloatField(null=True, help_text='Elapsed time since last report')

    thsensor = fields.ForeignKey('center.restapi.THSensorResource', '', readonly=True, null=True)

    class Meta(resource.ModelResource.Meta):
        queryset = Sensor.objects.all()
        authentication = Authentication()
        authorization = ReadOnlyAuthorization()
        excludes = ('last_seq',)

    def dehydrate(self, bundle):
        bundle = super(SensorResource, self).dehydrate(bundle)

        bundle.data['thsensor'] = THSensorResourceInstance.get_resource_uri(bundle.obj)
        bundle._cache = cache.get(bundle.obj._carbon_path())
        if bundle._cache:
            bundle.data['vcc'] = bundle._cache.get('Power', None)
            bundle.data['rssi'] = bundle._cache.get('RSSI', None)
            bundle.data['lqi'] = bundle._cache.get('LQI', None)
            bundle.data['interval'] = bundle._cache.get('intvl', None)

        return bundle

class THSensorResource(SensorResource):
    temperature = fields.FloatField(null=True)
    humidity = fields.FloatField(null=True)

    def dehydrate(self, bundle):
        bundle = super(THSensorResource, self).dehydrate(bundle)

        if bundle._cache:
            bundle.data['temperature'] = bundle._cache.get('Temperature', None)
            bundle.data['humidity'] = bundle._cache.get('Humidity', None)

        return bundle

THSensorResourceInstance = THSensorResource()

restapi.RestApi.register(SensorResource())
restapi.RestApi.register(THSensorResourceInstance)
