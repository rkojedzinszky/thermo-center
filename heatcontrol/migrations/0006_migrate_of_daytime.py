# -*- coding: utf-8 -*-
from __future__ import unicode_literals

from django.db import migrations, models

def migrate_of_daytime(apps, schema_editor):
    DayTime = apps.get_model('heatcontrol', 'DayTime')
    HeatSensor = apps.get_model('heatcontrol', 'HeatSensor')

    for hs in HeatSensor.objects.filter(daytime__isnull=False).select_related('daytime'):
        hs.daytype = hs.daytime.daytype
        hs.start = hs.daytime.start
        hs.end = hs.daytime.end
        hs.daytime = None
        hs.save()

class Migration(migrations.Migration):

    dependencies = [
        ('heatcontrol', '0005_auto_20160928_2101'),
    ]

    operations = [
            migrations.RunPython(migrate_of_daytime),
    ]
