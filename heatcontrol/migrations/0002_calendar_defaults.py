# -*- coding: utf-8 -*-
from __future__ import unicode_literals

import datetime
from django.db import migrations, models

def default_daytypes(apps, schema_editor):
    daytypemodel = apps.get_model('heatcontrol', 'DayType')
    daytypemodel.objects.create(name='workday')
    daytypemodel.objects.create(name='rest')

def default_calendar(apps, schema_editor):
    daytypemodel = apps.get_model('heatcontrol', 'DayType')
    workday = daytypemodel.objects.get(name='workday')
    rest = daytypemodel.objects.get(name='rest')
    calendarmodel = apps.get_model('heatcontrol', 'Calendar')

    day = datetime.date.today()
    for i in range(50*365): # 50 years
        daytype = workday

        # weekends
        if day.weekday() >= 5:
            daytype = rest

        # new year
        if day.month == 1 and day.day == 1:
            daytype = rest

        calendarmodel.objects.create(day=day, daytype=daytype)

        day += datetime.timedelta(1)

class Migration(migrations.Migration):

    dependencies = [
        ('heatcontrol', '0001_initial'),
    ]

    operations = [
            migrations.RunPython(default_daytypes),
            migrations.RunPython(default_calendar),
    ]
