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

@admin.register(models.ScheduledOverride)
class ScheduledOverrideAdmin(admin.ModelAdmin):
    list_display = ('control', 'start', 'end', 'target_temp')
    list_filter = ('control', )

    def get_queryset(self, request):
        qs = super().get_queryset(request)
        return qs.filter(end__gte=timezone.now()).order_by('start')

@admin.register(models.InstantProfile)
class InstantProfileAdmin(admin.ModelAdmin):
    list_display = ('name', 'active')
    readonly_fields = ('active', )

@admin.register(models.InstantProfileEntry)
class InstantProfileEntryAdmin(admin.ModelAdmin):
    list_display = ('profile', 'control', 'target_temp', 'active')
    list_display_links = ('profile', 'control')
    list_filter = ('profile', 'control')
    readonly_fields = ('active', )
