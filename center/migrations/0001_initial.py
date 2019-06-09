# -*- coding: utf-8 -*-
from __future__ import unicode_literals

from django.db import migrations, models
import center.fields


class Migration(migrations.Migration):

    dependencies = [
    ]

    operations = [
        migrations.CreateModel(
            name='RFConfig',
            fields=[
                ('id', models.AutoField(verbose_name='ID', serialize=False, auto_created=True, primary_key=True)),
                ('rf_channel', center.fields.RangedIntegerField(max_value=255, min_value=0)),
                ('network_id', center.fields.RangedIntegerField(max_value=65535, min_value=0)),
                ('aes_key', models.CharField(max_length=32)),
            ],
        ),
        migrations.CreateModel(
            name='RFProfile',
            fields=[
                ('id', models.AutoField(verbose_name='ID', serialize=False, auto_created=True, primary_key=True)),
                ('name', models.CharField(max_length=50)),
                ('confregs', models.CharField(max_length=128)),
            ],
            options={
                'ordering': ['pk'],
            },
        ),
        migrations.CreateModel(
            name='Sensor',
            fields=[
                ('id', center.fields.SensorIdField(serialize=False, primary_key=True)),
                ('name', models.CharField(max_length=100, blank=True)),
                ('last_seq', models.PositiveIntegerField(null=True)),
                ('last_ts', models.DateTimeField(null=True)),
            ],
        ),
        migrations.AddField(
            model_name='rfconfig',
            name='rf_profile',
            field=models.ForeignKey(to='center.RFProfile', on_delete=models.PROTECT),
        ),
    ]
