# -*- coding: utf-8 -*-
from __future__ import unicode_literals

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('heatcontrol', '0002_auto_20161110_1416'),
    ]

    operations = [
        migrations.CreateModel(
            name='HeatControlOverride',
            fields=[
                ('id', models.AutoField(verbose_name='ID', serialize=False, auto_created=True, primary_key=True)),
                ('start', models.DateTimeField()),
                ('end', models.DateTimeField()),
                ('target_temp', models.FloatField()),
                ('heatcontrol', models.ForeignKey(to='heatcontrol.HeatControl')),
            ],
        ),
        migrations.CreateModel(
            name='HeatControlProfile',
            fields=[
                ('id', models.AutoField(verbose_name='ID', serialize=False, auto_created=True, primary_key=True)),
                ('start', models.TimeField()),
                ('end', models.TimeField()),
                ('target_temp', models.FloatField()),
                ('daytype', models.ForeignKey(to='heatcontrol.DayType')),
                ('heatcontrol', models.ForeignKey(to='heatcontrol.HeatControl')),
            ],
        ),
        migrations.AlterIndexTogether(
            name='heatcontrolprofile',
            index_together=set([('heatcontrol', 'daytype')]),
        ),
        migrations.AlterIndexTogether(
            name='heatcontroloverride',
            index_together=set([('heatcontrol', 'end')]),
        ),
    ]
