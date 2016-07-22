""" API for center """

import datetime
from django.utils import timezone
from django.core.cache import cache
from application import restapi, resource
from center.models import Sensor
from tastypie.authentication import Authentication
from tastypie.authorization import ReadOnlyAuthorization
from tastypie import fields
from graphite.render.evaluator import evaluateTarget

class SensorResource(resource.ModelResource):
    vcc = fields.FloatField(null=True, help_text='Power in battery in sensor')
    rssi = fields.FloatField(null=True, help_text='RSSI of sensor')
    lqi = fields.FloatField(null=True, help_text='Line Quality Indicator of sensor')
    interval = fields.FloatField(null=True, help_text='Elapsed time since last report')
    age = fields.FloatField(null=True, readonly=True)
    server_time = fields.DateTimeField(readonly=True)

    thsensor = fields.ForeignKey('center.restapi.THSensorResource', '', readonly=True, null=True)

    class Meta(resource.ModelResource.Meta):
        queryset = Sensor.objects.all()
        authentication = Authentication()
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
            bundle.data['vcc'] = bundle._cache.get('Power', None)
            bundle.data['rssi'] = bundle._cache.get('RSSI', None)
            bundle.data['lqi'] = bundle._cache.get('LQI', None)
            bundle.data['interval'] = bundle._cache.get('intvl', None)

        bundle.data['server_time'] = now
        if bundle.obj.last_ts:
            bundle.data['age'] = (now - bundle.obj.last_ts).total_seconds()

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

class MetricResource(SensorResource):
    start = fields.DateTimeField(readonly=True)
    end = fields.DateTimeField(readonly=True)
    step = fields.IntegerField(readonly=True)
    values = fields.ListField(readonly=True)

    class Meta(SensorResource.Meta):
        fields = ('id', 'start', 'end', 'step', 'values')
        list_allowed_methods = []
        detail_allowed_methods = ['get']

    def dehydrate(self, bundle):
        rc = {
            'startTime': self.start.convert(bundle.request.GET.get('start')),
            'endTime': self.end.convert(bundle.request.GET.get('end')),
            'now': timezone.now(),
            'localOnly': False,
            'data': []
        }
        s = bundle.obj
        ts = evaluateTarget(rc, '%s.%s' % (s._carbon_path(), self.metric))[0]
        bundle.data['start'] = datetime.datetime.fromtimestamp(ts.start)
        bundle.data['end'] = datetime.datetime.fromtimestamp(ts.end)
        bundle.data['step'] = ts.step
        bundle.data['values'] = ts

        return bundle
        #return super(MetricResource, self).dehydrate(bundle)

class TemperatureResource(MetricResource):
    metric = 'Temperature'

restapi.RestApi.register(TemperatureResource())
