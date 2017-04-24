# -*- coding: utf-8 -*-
from __future__ import unicode_literals

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('heatcontrol', '0005_auto_20161204_1429'),
    ]

    operations = [
            migrations.RenameModel('HeatControl', 'Control'),
            migrations.RenameField('HeatControlProfile', 'heatcontrol', 'control'),
            migrations.RenameField('HeatControlOverride', 'heatcontrol', 'control'),
    ]
