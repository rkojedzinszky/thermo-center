
from tastypie.http import *
from tastypie.exceptions import ImmediateHttpResponse
from django.http import HttpResponse

class HttpPreconditionFailed(HttpResponse):
    status_code = 412

