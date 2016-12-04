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

class HeatControlResource(resource.ModelResource):
    name = fields.CharField(readonly=True)
    temperature = fields.FloatField(readonly=True, null=True)
    target_temp = fields.FloatField(readonly=True, null=True)
    pidcontrol = fields.FloatField(readonly=True, null=True)
    age = fields.FloatField(null=True, readonly=True)

    class Meta(resource.ModelResource.Meta):
        queryset = models.HeatControl.objects.select_related('sensor').order_by('sensor__id')
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

HeatControlResourceInstance = HeatControlResource()
restapi.RestApi.register(HeatControlResourceInstance)

class HeatControlProfileResource(resource.ModelResource):
    heatcontrol = fields.ForeignKey(HeatControlResource, 'heatcontrol')
    daytype = fields.ForeignKey(DayTypeResource, 'daytype')

    class Meta(resource.ModelResource.Meta):
        queryset = models.HeatControlProfile.objects.order_by('start')
        authorization = Authorization()
        filtering = {
                'heatcontrol': 'exact',
                'daytype': 'exact',
                }

HeatControlProfileResourceInstance = HeatControlProfileResource()
restapi.RestApi.register(HeatControlProfileResourceInstance)

class HeatControlOverrideResource(resource.ModelResource):
    heatcontrol = fields.ForeignKey(HeatControlResource, 'heatcontrol')

    class Meta(resource.ModelResource.Meta):
        queryset = models.HeatControlOverride.objects.order_by('start')
        authorization = Authorization()
        filtering = {
                'heatcontrol': 'exact',
                }

    def get_object_list(self, request):
        return super(HeatControlOverrideResource, self).get_object_list(request).filter(end__gt=timezone.now())

HeatControlOverrideResourceInstance = HeatControlOverrideResource()
restapi.RestApi.register(HeatControlOverrideResourceInstance)
