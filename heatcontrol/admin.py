from django.contrib import admin
from heatcontrol import models

admin.site.register(models.DayType)
admin.site.register(models.Calendar)
admin.site.register(models.HeatControl)
admin.site.register(models.HeatControlProfile)
admin.site.register(models.HeatControlOverride)
