from django.urls import include, path
from django.contrib import admin
from django.conf import settings
from django.utils.module_loading import import_string, autodiscover_modules
from django.conf.urls.static import static
from tastypie_openapi import SchemaView
from . import views

autodiscover_modules("restapi")

RestApi = import_string(settings.RESTAPI_CLASS)

urlpatterns = [
    path("api/", include(RestApi.urls)),
    path(
        "openapi/",
        SchemaView.as_view(api=RestApi, title="Thermo Center API", version="0.1.0"),
    ),
    path("admin/", admin.site.urls),
    path("healthz", views.healthz),
] + static("/", document_root=settings.WWW_FILES)
