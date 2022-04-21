""" Rest API """

import time
from django.contrib.auth import authenticate, login, logout
from django.conf import settings
from tastypie.resources import Resource
from application.resource import ResourceMetaCommon
from tastypie.authentication import Authentication
from tastypie import fields
from tastypie.bundle import Bundle
from application import http

from tastypie.api import Api
RestApi = Api(api_name='v1')


class SessionResource(Resource):
    id = fields.IntegerField(readonly=True)
    is_admin = fields.BooleanField(readonly=True)

    class Meta(ResourceMetaCommon):
        authentication = Authentication()
        list_allowed_methods = ['get', 'post']
        detail_allowed_methods = ['get', 'delete']

    def detail_uri_kwargs(self, bundle_or_obj):
        if isinstance(bundle_or_obj, Bundle):
            return {'pk': 1}

        return None

    def obj_get_list(self, bundle, **kwargs):
        if bundle.request.user.is_authenticated:
            return [bundle.request.user]

        return []

    def obj_create(self, bundle, **kwargs):
        user = authenticate(request=bundle.request, username=bundle.data['username'], password=bundle.data.pop('password', None))

        if user is not None and user.is_active:
            login(bundle.request, user)
            bundle.obj = user
        else:
            time.sleep(1)
            raise http.ImmediateHttpResponse(http.HttpUnauthorized())

        return bundle

    def obj_get(self, bundle, **kwargs):
        if kwargs.get('pk', None) == '1' and bundle.request.user.is_authenticated:
            return bundle.request.user

        raise http.ImmediateHttpResponse(http.HttpNotFound())

    def obj_delete(self, bundle, **kwargs):
        self.obj_get(bundle, **kwargs)

        logout(bundle.request)

    def dehydrate(self, bundle):
        bundle.data['id'] = 1
        bundle.data['is_admin'] = bundle.obj.is_superuser

        return bundle


SessionInstance = SessionResource()
RestApi.register(SessionInstance)
