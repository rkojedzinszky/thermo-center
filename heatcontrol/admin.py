from django.contrib import admin
from heatcontrol import models
from django.utils import timezone

class CalendarAdmin(admin.ModelAdmin):
    list_display = ('day', 'daytype')

    def get_queryset(self, request):
        qs = super().get_queryset(request)
        return qs.filter(day__gte=timezone.now()).order_by('day')

admin.site.register(models.DayType)
admin.site.register(models.Calendar, CalendarAdmin)
admin.site.register(models.Control)
admin.site.register(models.Profile)
admin.site.register(models.ScheduledOverride)
admin.site.register(models.InstantProfile)
admin.site.register(models.InstantProfileEntry)
