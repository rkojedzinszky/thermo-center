# -*- coding: utf-8 -*-
# Generated by Django 1.11.12 on 2018-12-27 09:40
from __future__ import unicode_literals

import datetime
from django.db import migrations, models


def empty_profile_end(apps, schema_editor):
    Profile = apps.get_model('heatcontrol', 'Profile')

    Profile.objects.all().update(end=None)


def fill_control_profile_end_w_daytype(control, daytype):
    last = None

    for e in control.profile_set.filter(daytype=daytype).order_by('start'):
        if last:
            last.end = e.start
            last.save()

        last = e

    if last:
        last.end = datetime.time()
        last.save()


def fill_profile_end(apps, schema_editor):
    Control = apps.get_model('heatcontrol', 'Control')
    daytypes = apps.get_model('heatcontrol', 'DayType').objects.all()

    for control in Control.objects.all():
        for daytype in daytypes:
            fill_control_profile_end_w_daytype(control, daytype)


class Migration(migrations.Migration):

    dependencies = [
        ('heatcontrol', '0016_heatcontrol_profile_end_null'),
    ]

    operations = [
        migrations.RunPython(empty_profile_end, fill_profile_end),
    ]