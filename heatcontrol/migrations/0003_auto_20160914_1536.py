# -*- coding: utf-8 -*-
from __future__ import unicode_literals

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('heatcontrol', '0002_calendar_defaults'),
    ]

    operations = [
        migrations.AlterUniqueTogether(
            name='heatsensor',
            unique_together=set([('sensor', 'daytime')]),
        ),
    ]
