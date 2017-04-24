# -*- coding: utf-8 -*-
from __future__ import unicode_literals

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('heatcontrol', '0007_auto_20170424_2044'),
    ]

    operations = [
            migrations.RenameModel('HeatControlOverride', 'ScheduledOverride'),
    ]
