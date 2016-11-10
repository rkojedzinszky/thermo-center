# -*- coding: utf-8 -*-
from __future__ import unicode_literals

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('heatcontrol', '0001_initial_20161004'),
    ]

    operations = [
            migrations.RenameModel('PidControlParams', 'HeatControl'),
    ]
