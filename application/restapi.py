""" Rest API """

import time
import ipaddress
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
    id = fields.CharField()

    class Meta(ResourceMetaCommon):
        authentication = Authentication()
        list_allowed_methods = ['get', 'post']
        detail_allowed_methods = ['delete']

    def detail_uri_kwargs(self, bundle_or_obj):
        if isinstance(bundle_or_obj, Bundle):
            return {'pk': bundle_or_obj.request.session.session_key}

        return None

    def obj_get_list(self, bundle, **kwargs):
        if bundle.request.user.is_authenticated():
            return [bundle.request.user]

        return []

    def obj_get(self, bundle, **kwargs):
        if bundle.request.user.is_authenticated() and kwargs['pk'] == bundle.request.session.session_key:
            return bundle.request.user

        raise http.ImmediateHttpResponse(http.HttpUnauthorized())

    def obj_create(self, bundle, **kwargs):
        user = authenticate(request=bundle.request, username=bundle.data['username'], password=bundle.data.pop('password', None))

        if user is not None and user.is_active:
            login(bundle.request, user)
            bundle.obj = user
        else:
            time.sleep(1)
            raise http.ImmediateHttpResponse(http.HttpUnauthorized())

        return bundle

    def obj_delete(self, bundle, **kwargs):
        logout(bundle.request)

    def dehydrate(self, bundle):
        bundle = super(SessionResource, self).dehydrate(bundle)
        bundle.data['id'] = bundle.request.session.session_key

        return bundle

SessionInstance = SessionResource()
RestApi.register(SessionInstance)
