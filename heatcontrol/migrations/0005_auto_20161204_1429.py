# -*- coding: utf-8 -*-
from __future__ import unicode_literals

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('heatcontrol', '0004_migrate_to_heatcontrol'),
    ]

    operations = [
        migrations.AlterIndexTogether(
            name='heatsensor',
            index_together=set([]),
        ),
        migrations.RemoveField(
            model_name='heatsensor',
            name='daytype',
        ),
        migrations.RemoveField(
            model_name='heatsensor',
            name='sensor',
        ),
        migrations.AlterIndexTogether(
            name='heatsensoroverride',
            index_together=set([]),
        ),
        migrations.RemoveField(
            model_name='heatsensoroverride',
            name='sensor',
        ),
        migrations.DeleteModel(
            name='HeatSensor',
        ),
        migrations.DeleteModel(
            name='HeatSensorOverride',
        ),
    ]
