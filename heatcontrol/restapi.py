""" API for heatcontrol """

from tastypie import fields
from tastypie.utils import timezone
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

class HeatSensorTimeAuthorization(resource.NoAuthorization):
    def read_list(self, object_list, bundle):
        return object_list

    def read_detail(self, object_list, bundle):
        return True

    def create_detail(self, object_list, bundle):
        return True

    def update_detail(self, object_list, bundle):
        return True

    def delete_detail(self, object_list, bundle):
        return True

class HeatSensorTimeResource(resource.ModelResource):
    sensor = fields.ForeignKey(HeatSensorResource, 'sensor')
    daytype = fields.ForeignKey(DayTypeResource, 'daytype')

    class Meta(resource.ModelResource.Meta):
        queryset = models.HeatSensor.objects.order_by('start')
        authorization = HeatSensorTimeAuthorization()
        filtering = {
                'sensor': 'exact',
                'daytype': 'exact',
                }

HeatSensorTimeResourceInstance = HeatSensorTimeResource()
restapi.RestApi.register(HeatSensorTimeResourceInstance)

class HeatSensorOverrideAuthorization(resource.NoAuthorization):
    def read_list(self, object_list, bundle):
        return object_list

    def read_detail(self, object_list, bundle):
        return True

    def create_detail(self, object_list, bundle):
        return True

    def update_detail(self, object_list, bundle):
        return True

    def delete_detail(self, object_list, bundle):
        return True

class HeatSensorOverrideResource(resource.ModelResource):
    sensor = fields.ForeignKey(HeatSensorResource, 'sensor')

    class Meta(resource.ModelResource.Meta):
        queryset = models.HeatSensorOverride.objects.order_by('start')
        authorization = HeatSensorOverrideAuthorization()
        filtering = {
                'sensor': 'exact',
                }

    def get_object_list(self, request):
        return super(HeatSensorOverrideResource, self).get_object_list(request).filter(end__gt=timezone.now())

HeatSensorOverrideResourceInstance = HeatSensorOverrideResource()
restapi.RestApi.register(HeatSensorOverrideResourceInstance)
