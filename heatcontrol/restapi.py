""" API for heatcontrol """

import datetime, time
from tastypie import fields
from tastypie import resources
from tastypie.utils import timezone
from tastypie.authorization import ReadOnlyAuthorization, Authorization
from center.restapi import THSensorResource, THSensorResourceInstance
from application import restapi
from application.resource import ResourceMetaCommon
from heatcontrol import models
from tastypie.bundle import Bundle

import logging
logger = logging.getLogger(__name__)

class DayTypeResource(resources.ModelResource):
    class Meta(ResourceMetaCommon):
        queryset = models.DayType.objects.all()
        authorization = ReadOnlyAuthorization()
        filtering = {
                'name': 'exact',
                }

DayTypeResourceInstance = DayTypeResource()
restapi.RestApi.register(DayTypeResourceInstance)

class ControlResource(resources.ModelResource):
    sensor_id = fields.IntegerField('sensor_id', readonly=True)
    name = fields.CharField(readonly=True)
    temperature = fields.FloatField(readonly=True, null=True)
    target_temp = fields.FloatField(readonly=True, null=True)
    pidcontrol = fields.FloatField(readonly=True, null=True)
    age = fields.FloatField(null=True, readonly=True)

    class Meta(ResourceMetaCommon):
        queryset = models.Control.objects.select_related('sensor').order_by('sensor__id')
        authorization = Authorization()
        filtering = {
                'sensor_id': 'exact',
                }

    def dehydrate_name(self, bundle):
        return bundle.obj.sensor.name

    def dehydrate_target_temp(self, bundle):
        return bundle.obj.get_target_temp()

    def dehydrate_age(self, bundle):
        if bundle.obj.sensor.last_tsf:
            return time.time() - bundle.obj.sensor.last_tsf

        return None

    def dehydrate(self, bundle):
        c = bundle.obj.sensor.get_cache()
        if c:
            bundle.data['temperature'] = c.get('Temperature')
            bundle.data['pidcontrol'] = c.get('pidcontrol')

        return bundle

ControlResourceInstance = ControlResource()
restapi.RestApi.register(ControlResourceInstance)

class ProfileResource(resources.ModelResource):
    control = fields.ForeignKey(ControlResource, 'control')
    daytype = fields.ForeignKey(DayTypeResource, 'daytype')

    class Meta(ResourceMetaCommon):
        queryset = models.Profile.objects.order_by('start')
        authorization = Authorization()
        filtering = {
                'control': 'exact',
                'daytype': 'exact',
                }
        ordering = ('start',)

ProfileResourceInstance = ProfileResource()
restapi.RestApi.register(ProfileResourceInstance)

class ScheduledOverrideResource(resources.ModelResource):
    control = fields.ForeignKey(ControlResource, 'control')

    class Meta(ResourceMetaCommon):
        queryset = models.ScheduledOverride.objects.order_by('start')
        authorization = Authorization()
        filtering = {
                'control': 'exact',
                }

    def get_object_list(self, request):
        return super(ScheduledOverrideResource, self).get_object_list(request).filter(end__gt=timezone.now())

ScheduledOverrideResourceInstance = ScheduledOverrideResource()
restapi.RestApi.register(ScheduledOverrideResourceInstance)

class InstantProfileResourceAuthorization(ReadOnlyAuthorization):
    def update_detail(self, object_list, bundle):
        return True

class InstantProfileResource(resources.ModelResource):
    class Meta(ResourceMetaCommon):
        always_return_data = False
        queryset = models.InstantProfile.objects.order_by('id')
        authorization = InstantProfileResourceAuthorization()

InstantProfileResourceInstance = InstantProfileResource()
restapi.RestApi.register(InstantProfileResourceInstance)

class CurrentDaytypeAuthorization(ReadOnlyAuthorization):
    def update_detail(self, object_list, bundle):
        return True

class CurrentDaytypeResource(resources.ModelResource):
    daytype = fields.CharField()

    class Meta(ResourceMetaCommon):
        queryset = models.Calendar.objects.all()
        authorization = CurrentDaytypeAuthorization()
        list_allowed_methods = []
        detail_allowed_methods = ['get', 'patch']
        fields = ('daytype',)

    def dehydrate_daytype(self, bundle):
        return bundle.obj.daytype.name

    def hydrate_daytype(self, bundle):
        bundle.obj.daytype = DayTypeResourceInstance.obj_get(bundle=Bundle(request=bundle.request), name=bundle.data['daytype'])
        return bundle

    def obj_get(self, bundle, **kwargs):
        return models.Calendar.objects.get(day=datetime.date.today())

CurrentDaytypeResourceInstance = CurrentDaytypeResource()
restapi.RestApi.register(CurrentDaytypeResourceInstance)
