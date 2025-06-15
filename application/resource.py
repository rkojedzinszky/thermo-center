"""resource-related classes, overriding tastypie defaults"""

from tastypie import resources
from tastypie.authentication import (
    BasicAuthentication,
    ApiKeyAuthentication,
    SessionAuthentication,
    MultiAuthentication,
)
from tastypie.authorization import Authorization


class ReadonlyPkeyResource(resources.ModelResource):
    @classmethod
    def get_fields(cls, fields=None, excludes=None):
        final_fields = super().get_fields(fields, excludes)

        # Set readonly on primary_key field
        for name, field in final_fields.items():
            djangofield = cls._meta.object_class._meta.get_field(name)
            if djangofield.primary_key:
                field.readonly = True

        return final_fields


class ResourceMetaCommon:
    authentication = MultiAuthentication(
        SessionAuthentication(), ApiKeyAuthentication(), BasicAuthentication()
    )
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
