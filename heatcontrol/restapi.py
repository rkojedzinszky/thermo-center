""" API for heatcontrol """

from tastypie import fields
from center.restapi import THSensorResource
from application import restapi, resource

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
