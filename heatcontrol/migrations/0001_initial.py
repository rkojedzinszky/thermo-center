# -*- coding: utf-8 -*-
from __future__ import unicode_literals

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('center', '0003_default_rfconfig'),
    ]

    operations = [
        migrations.CreateModel(
            name='Calendar',
            fields=[
                ('id', models.AutoField(verbose_name='ID', serialize=False, auto_created=True, primary_key=True)),
                ('day', models.DateField(unique=True)),
            ],
        ),
        migrations.CreateModel(
            name='DayTime',
            fields=[
                ('id', models.AutoField(verbose_name='ID', serialize=False, auto_created=True, primary_key=True)),
                ('start', models.TimeField()),
                ('end', models.TimeField()),
            ],
        ),
        migrations.CreateModel(
            name='DayType',
            fields=[
                ('id', models.AutoField(verbose_name='ID', serialize=False, auto_created=True, primary_key=True)),
                ('name', models.CharField(max_length=50)),
            ],
        ),
        migrations.CreateModel(
            name='HeatSensor',
            fields=[
                ('id', models.AutoField(verbose_name='ID', serialize=False, auto_created=True, primary_key=True)),
                ('target_temp', models.FloatField()),
                ('daytime', models.ForeignKey(to='heatcontrol.DayTime')),
                ('sensor', models.ForeignKey(to='center.Sensor')),
            ],
        ),
        migrations.AddField(
            model_name='daytime',
            name='daytype',
            field=models.ForeignKey(to='heatcontrol.DayType'),
        ),
        migrations.AddField(
            model_name='calendar',
            name='daytype',
            field=models.ForeignKey(to='heatcontrol.DayType'),
        ),
    ]
