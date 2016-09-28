# -*- coding: utf-8 -*-
from __future__ import unicode_literals

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('heatcontrol', '0004_heatcontrol_heatsensoroverride'),
    ]

    operations = [
        migrations.AddField(
            model_name='heatsensor',
            name='daytype',
            field=models.ForeignKey(to='heatcontrol.DayType', null=True),
        ),
        migrations.AddField(
            model_name='heatsensor',
            name='end',
            field=models.TimeField(null=True),
        ),
        migrations.AddField(
            model_name='heatsensor',
            name='start',
            field=models.TimeField(null=True),
        ),
        migrations.AlterField(
            model_name='heatsensor',
            name='daytime',
            field=models.ForeignKey(to='heatcontrol.DayTime', null=True),
        ),
    ]
