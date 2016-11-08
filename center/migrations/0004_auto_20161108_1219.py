# -*- coding: utf-8 -*-
from __future__ import unicode_literals

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('center', '0003_default_rfconfig'),
    ]

    operations = [
        migrations.RemoveField(
            model_name='sensor',
            name='last_ts',
        ),
        migrations.AddField(
            model_name='sensor',
            name='last_tsf',
            field=models.FloatField(null=True),
        ),
    ]
