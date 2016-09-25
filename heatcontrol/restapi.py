""" API for heatcontrol """

from tastypie import fields
from tastypie.authorization import ReadOnlyAuthorization
from center.restapi import THSensorResource, THSensorResourceInstance
from application import restapi, resource
from heatcontrol import models
from tastypie.bundle import Bundle

import logging
logger = logging.getLogger(__name__)

class HeatSensorResource(THSensorResource):
    target_temp = fields.FloatField(null=True)
    pidcontrol = fields.FloatField(null=True)

    class Meta(THSensorResource.Meta):
        fields = ('id', 'name', 'temperature', 'target_temp', 'age', 'server_time')

    def dehydrate(self, bundle):
        bundle = super(HeatSensorResource, self).dehydrate(bundle)

        bundle.data['target_temp'] = bundle.obj.get_target_temp()

        if bundle._cache:
            bundle.data['pidcontrol'] = bundle._cache.get('pidcontrol', None)

        return bundle

HeatSensorResourceInstance = HeatSensorResource()
restapi.RestApi.register(HeatSensorResourceInstance)

class DayTypeResource(resource.ModelResource):
    class Meta(resource.ModelResource.Meta):
        queryset = models.DayType.objects.all()
        authorization = ReadOnlyAuthorization()

DayTypeResourceInstance = DayTypeResource()
restapi.RestApi.register(DayTypeResourceInstance)

class DayTimeResource(resource.ModelResource):
    class Meta(resource.ModelResource.Meta):
        queryset = models.DayTime.objects.all()
        authorization = ReadOnlyAuthorization()

DayTimeResourceInstance = DayTimeResource()
restapi.RestApi.register(DayTimeResourceInstance)

class HeatSensorTimeAuthorization(resource.NoAuthorization):
    def update_detail(self, object_list, bundle):
        return True

class HeatSensorTimeResource(resource.ModelResource):
    id = fields.CharField('id', unique=True, readonly=True)
    target_temp = fields.FloatField('target_temp', null=True)

    class Meta(resource.ModelResource.Meta):
        queryset = models.DayTime.objects.order_by('start')
        authorization = HeatSensorTimeAuthorization()

    def obj_get_list(self, bundle, **kwargs):
        dtbundle = DayTypeResourceInstance.build_bundle(request=bundle.request)
        daytype = DayTypeResourceInstance.obj_get(dtbundle, pk=bundle.request.GET['daytype_id'])
        thsbundle = THSensorResourceInstance.build_bundle(request=bundle.request)
        sensor = THSensorResourceInstance.obj_get(thsbundle, pk=bundle.request.GET['sensor_id'])

        shsdict = {shs.daytime_id: shs.target_temp for shs in sensor.heatsensor_set.filter(daytime__daytype=daytype)}

        res = self._meta.queryset.filter(daytype=daytype)

        for t in res:
            t.target_temp = shsdict.get(t.id, None)
            t.sensor = sensor
            t.id = '%d-%d' % (sensor.id, t.id)

        return res

    def obj_get(self, bundle, **kwargs):
        sensor_id, daytime_id = kwargs[self._meta.detail_uri_name].split('-')
        thsbundle = THSensorResourceInstance.build_bundle(request=bundle.request)
        sensor = THSensorResourceInstance.obj_get(thsbundle, pk=sensor_id)
        daytime = models.DayTime.objects.get(pk=daytime_id)
        daytime.sensor = sensor
        try:
            hs = models.HeatSensor.objects.get(sensor=sensor, daytime=daytime)
            daytime.target_temp = hs.target_temp
        except models.HeatSensor.DoesNotExist:
            daytime.target_temp = None

        daytime._id = daytime.id
        daytime.id = '%d-%d' % (sensor.id, daytime.id)

        return daytime

    def obj_update(self, bundle, **kwargs):
        self.authorized_update_detail(self.get_object_list(bundle.request), bundle)

        bundle.obj.target_temp = bundle.data['target_temp']

        if bundle.obj.target_temp is None:
            models.HeatSensor.objects.filter(daytime=bundle.obj._id, sensor=bundle.obj.sensor).delete()
        else:
            hs = models.HeatSensor.objects.filter(daytime=bundle.obj._id, sensor=bundle.obj.sensor).first()
            if hs is None:
                hs = models.HeatSensor(daytime_id=bundle.obj._id, sensor=bundle.obj.sensor)

            hs.target_temp = bundle.obj.target_temp

            hs.save()

        return bundle

HeatSensorTimeResourceInstance = HeatSensorTimeResource()
restapi.RestApi.register(HeatSensorTimeResourceInstance)
