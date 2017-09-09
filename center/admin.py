from django.contrib import admin
from center import models

class RFProfileAdmin(admin.ModelAdmin):
    list_display = ('name',)
    ordering = ('id',)

admin.site.register(models.RFProfile, RFProfileAdmin)

class RFConfigAdmin(admin.ModelAdmin):
    list_display = ('rf_channel', 'rf_profile', 'network_id')
    actions = None

admin.site.register(models.RFConfig, RFConfigAdmin)

class SensorAdmin(admin.ModelAdmin):
    list_display = ('id', 'name', 'last_seq', 'last_tsf')
    readonly_fields = ('last_seq', 'last_tsf')
    ordering = ('id',)

admin.site.register(models.Sensor, SensorAdmin)
