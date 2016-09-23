# -*- coding: utf-8 -*-
from __future__ import unicode_literals

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('center', '0003_default_rfconfig'),
        ('heatcontrol', '0003_auto_20160914_1536'),
    ]

    operations = [
        migrations.CreateModel(
            name='HeatSensorOverride',
            fields=[
                ('id', models.AutoField(verbose_name='ID', serialize=False, auto_created=True, primary_key=True)),
                ('start', models.DateTimeField()),
                ('end', models.DateTimeField()),
                ('target_temp', models.FloatField()),
                ('sensor', models.ForeignKey(to='center.Sensor')),
            ],
        ),
        migrations.AlterIndexTogether(
            name='heatsensoroverride',
            index_together=set([('sensor', 'end')]),
        ),
    ]
