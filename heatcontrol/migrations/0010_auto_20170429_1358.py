# -*- coding: utf-8 -*-
from __future__ import unicode_literals

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('heatcontrol', '0009_instantoverride'),
    ]

    operations = [
        migrations.AlterField(
            model_name='instantoverride',
            name='target_temp',
            field=models.FloatField(null=True),
        ),
        migrations.AlterField(
            model_name='instantprofileentry',
            name='target_temp',
            field=models.FloatField(null=True, blank=True),
        ),
    ]
