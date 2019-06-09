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
                ('start', models.TimeField()),
                ('end', models.TimeField()),
                ('target_temp', models.FloatField()),
                ('daytype', models.ForeignKey(to='heatcontrol.DayType', on_delete=models.PROTECT)),
                ('sensor', models.ForeignKey(to='center.Sensor', on_delete=models.CASCADE)),
            ],
        ),
        migrations.CreateModel(
            name='HeatSensorOverride',
            fields=[
                ('id', models.AutoField(verbose_name='ID', serialize=False, auto_created=True, primary_key=True)),
                ('start', models.DateTimeField()),
                ('end', models.DateTimeField()),
                ('target_temp', models.FloatField()),
                ('sensor', models.ForeignKey(to='center.Sensor', on_delete=models.CASCADE)),
            ],
        ),
        migrations.CreateModel(
            name='PidControlParams',
            fields=[
                ('id', models.AutoField(verbose_name='ID', serialize=False, auto_created=True, primary_key=True)),
                ('kp', models.FloatField()),
                ('ki', models.FloatField()),
                ('kd', models.FloatField()),
                ('sensor', models.OneToOneField(to='center.Sensor', on_delete=models.CASCADE)),
            ],
        ),
        migrations.AddField(
            model_name='calendar',
            name='daytype',
            field=models.ForeignKey(to='heatcontrol.DayType', on_delete=models.PROTECT),
        ),
        migrations.AlterIndexTogether(
            name='heatsensoroverride',
            index_together=set([('sensor', 'end')]),
        ),
        migrations.AlterIndexTogether(
            name='heatsensor',
            index_together=set([('sensor', 'daytype')]),
        ),
    ]
