# -*- coding: utf-8 -*-
from __future__ import unicode_literals

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('center', '0003_default_rfconfig'),
        ('heatcontrol', '0007_auto_20160928_2109'),
    ]

    operations = [
        migrations.CreateModel(
            name='PidControlParams',
            fields=[
                ('id', models.AutoField(verbose_name='ID', serialize=False, auto_created=True, primary_key=True)),
                ('kp', models.FloatField()),
                ('ki', models.FloatField()),
                ('kd', models.FloatField()),
                ('sensor', models.OneToOneField(to='center.Sensor')),
            ],
        ),
    ]
