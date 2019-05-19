from django.conf.urls import include, url
from django.contrib import admin
from django.conf import settings
from django.utils.module_loading import import_string, autodiscover_modules
from django.conf.urls.static import static
from . import views

autodiscover_modules('restapi')

RestApi = import_string(settings.RESTAPI_CLASS)

urlpatterns = [
    url(r'^' + settings.WWW_ROOT + 'api/', include(RestApi.urls)),
    url(r'^' + settings.WWW_ROOT + 'admin/', admin.site.urls),
    url(r'^healthz$', views.healthz),
] + static(settings.WWW_ROOT, document_root=settings.WWW_FILES)
