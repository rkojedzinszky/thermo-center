# -*- coding: utf-8 -*-
from __future__ import unicode_literals

from django.db import migrations, models

def migrate_HeatSensor(apps, schema_editor):
    HeatSensor = apps.get_model('heatcontrol', 'HeatSensor')
    HeatControl = apps.get_model('heatcontrol', 'HeatControl')
    HeatControlProfile = apps.get_model('heatcontrol', 'HeatControlProfile')

    for hs in HeatSensor.objects.select_related('sensor', 'daytype').all():
        hc, created = HeatControl.objects.get_or_create(sensor=hs.sensor, defaults={'kp': 1, 'ki': 1, 'kd': 1})
        HeatControlProfile.objects.create(heatcontrol=hc, daytype=hs.daytype, start=hs.start, end=hs.end, target_temp=hs.target_temp)

def migrate_HeatSensorOverride(apps, schema_editor):
    HeatSensorOverride = apps.get_model('heatcontrol', 'HeatSensorOverride')
    HeatControl = apps.get_model('heatcontrol', 'HeatControl')
    HeatControlOverride = apps.get_model('heatcontrol', 'HeatControlOverride')

    for hso in HeatSensorOverride.objects.select_related('sensor').all():
        hc, created = HeatControl.objects.get_or_create(sensor=hso.sensor, defaults={'kp': 1, 'ki': 1, 'kd': 1})
        HeatControlOverride.objects.create(heatcontrol=hc, start=hso.start, end=hso.end, target_temp=hso.target_temp)

class Migration(migrations.Migration):

    dependencies = [
        ('heatcontrol', '0003_auto_20161204_0620'),
    ]

    operations = [
            migrations.RunPython(migrate_HeatSensor),
            migrations.RunPython(migrate_HeatSensorOverride),
    ]
