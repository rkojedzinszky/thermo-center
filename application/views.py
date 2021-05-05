""" Healthcheck view """

from django.http import HttpResponse

async def healthz(request):
    return HttpResponse('OK', content_type='text/plain')
