""" API for heatcontrol """

import time
from tastypie import fields
from django.core.cache import cache
from tastypie.utils import timezone
from tastypie.authorization import ReadOnlyAuthorization, Authorization
from center.restapi import THSensorResource, THSensorResourceInstance
from application import restapi, resource
from heatcontrol import models
from tastypie.bundle import Bundle
from tastypie import fields

import logging
logger = logging.getLogger(__name__)

class DayTypeResource(resource.ModelResource):
    class Meta(resource.ModelResource.Meta):
        queryset = models.DayType.objects.all()
        authorization = ReadOnlyAuthorization()

DayTypeResourceInstance = DayTypeResource()
restapi.RestApi.register(DayTypeResourceInstance)

class ControlResource(resource.ModelResource):
    sensor_id = fields.IntegerField('sensor_id', readonly=True)
    name = fields.CharField(readonly=True)
    temperature = fields.FloatField(readonly=True, null=True)
    target_temp = fields.FloatField(readonly=True, null=True)
    pidcontrol = fields.FloatField(readonly=True, null=True)
    age = fields.FloatField(null=True, readonly=True)

    class Meta(resource.ModelResource.Meta):
        queryset = models.Control.objects.select_related('sensor').order_by('sensor__id')
        authorization = Authorization()

    def dehydrate_name(self, bundle):
        return bundle.obj.sensor.name

    def dehydrate_target_temp(self, bundle):
        return bundle.obj.get_target_temp()

    def dehydrate_age(self, bundle):
        if bundle.obj.sensor.last_tsf:
            return time.time() - bundle.obj.sensor.last_tsf

        return None

    def dehydrate(self, bundle):
        c = cache.get(bundle.obj.sensor._carbon_path())
        if c:
            bundle.data['temperature'] = c.get('Temperature')
            bundle.data['pidcontrol'] = c.get('pidcontrol')

        return bundle

ControlResourceInstance = ControlResource()
restapi.RestApi.register(ControlResourceInstance)

class ProfileResource(resource.ModelResource):
    control = fields.ForeignKey(ControlResource, 'control')
    daytype = fields.ForeignKey(DayTypeResource, 'daytype')

    class Meta(resource.ModelResource.Meta):
        queryset = models.Profile.objects.order_by('start')
        authorization = Authorization()
        filtering = {
                'control': 'exact',
                'daytype': 'exact',
                }

ProfileResourceInstance = ProfileResource()
restapi.RestApi.register(ProfileResourceInstance)

class ScheduledOverrideResource(resource.ModelResource):
    control = fields.ForeignKey(ControlResource, 'control')

    class Meta(resource.ModelResource.Meta):
        queryset = models.ScheduledOverride.objects.order_by('start')
        authorization = Authorization()
        filtering = {
                'control': 'exact',
                }

    def get_object_list(self, request):
        return super(ScheduledOverrideResource, self).get_object_list(request).filter(end__gt=timezone.now())

ScheduledOverrideResourceInstance = ScheduledOverrideResource()
restapi.RestApi.register(ScheduledOverrideResourceInstance)
