# -*- coding: utf-8 -*-
from __future__ import unicode_literals

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('heatcontrol', '0006_migrate_of_daytime'),
    ]

    operations = [
        migrations.RemoveField(
            model_name='daytime',
            name='daytype',
        ),
        migrations.AlterField(
            model_name='heatsensor',
            name='daytype',
            field=models.ForeignKey(to='heatcontrol.DayType'),
        ),
        migrations.AlterField(
            model_name='heatsensor',
            name='end',
            field=models.TimeField(),
        ),
        migrations.AlterField(
            model_name='heatsensor',
            name='start',
            field=models.TimeField(),
        ),
        migrations.AlterUniqueTogether(
            name='heatsensor',
            unique_together=set([]),
        ),
        migrations.AlterIndexTogether(
            name='heatsensor',
            index_together=set([('sensor', 'daytype')]),
        ),
        migrations.RemoveField(
            model_name='heatsensor',
            name='daytime',
        ),
        migrations.DeleteModel(
            name='DayTime',
        ),
    ]
