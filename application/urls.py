from django.conf.urls import include, url
from django.contrib import admin
from django.conf import settings
from django.utils.module_loading import import_string, autodiscover_modules

autodiscover_modules('restapi')

RestApi = import_string(settings.RESTAPI_CLASS)

urlpatterns = [
    url(r'^' + settings.WWW_ROOT + 'api/', include(RestApi.urls)),
    url(r'^' + settings.WWW_ROOT + 'admin/', admin.site.urls),
]
