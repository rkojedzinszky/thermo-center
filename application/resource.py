""" resource-related classes, overriding tastypie defaults """

from tastypie.authentication import BasicAuthentication, ApiKeyAuthentication, SessionAuthentication, MultiAuthentication
from tastypie.authorization import Authorization

class ResourceMetaCommon:
    authentication = MultiAuthentication(SessionAuthentication(), ApiKeyAuthentication(), BasicAuthentication())
    always_return_data = True
    limit = 0
    max_limit = 0

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

