# -*- coding: utf-8 -*-
# Generated by Django 1.11.12 on 2018-12-06 20:21
from __future__ import unicode_literals

from django.db import migrations, models
import django.db.models.deletion


class Migration(migrations.Migration):

    dependencies = [
        ('center', '0005_new_rfprofiles'),
    ]

    operations = [
        migrations.CreateModel(
            name='SensorResync',
            fields=[
                ('id', models.AutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('ts', models.DateTimeField(auto_now_add=True)),
                ('sensor', models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, to='center.Sensor')),
            ],
        ),
    ]
