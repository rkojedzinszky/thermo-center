""" resource-related classes, overriding tastypie defaults """

from tastypie.resources import Resource as tResource, ModelResource as tModelResource
from tastypie.authentication import BasicAuthentication, ApiKeyAuthentication, SessionAuthentication, MultiAuthentication

class Resource(tResource):
    class Meta:
        authentication = MultiAuthentication(SessionAuthentication(), ApiKeyAuthentication(), BasicAuthentication())
        always_return_data = True
        limit = 0
        max_limit = 0

class ModelResource(tModelResource):
    class Meta(Resource.Meta):
        object_class = None

from tastypie.authorization import Authorization

class NoAuthorization(Authorization):
    def read_list(self, object_list, bundle):
        return []

    def read_detail(self, object_list, bundle):
        return False

    def create_list(self, object_list, bundle):
        return []

    def create_detail(self, object_list, bundle):
        return False

    def update_list(self, object_list, bundle):
        return []

    def update_detail(self, object_list, bundle):
        return False

    def delete_list(self, object_list, bundle):
        return []

    def delete_detail(self, object_list, bundle):
        return False

