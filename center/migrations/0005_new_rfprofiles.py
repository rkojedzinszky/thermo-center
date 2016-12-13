# -*- coding: utf-8 -*-
from __future__ import unicode_literals

from django.db import migrations, models
from django.core.management import call_command

def load_new_rfprofiles(apps, schema_editor):
    rfconfig = apps.get_model('center', 'RFConfig').objects.get(pk=1)
    profile_name = rfconfig.rf_profile.name
    RFProfile = apps.get_model('center', 'RFProfile')
    RFProfile.objects.all().delete()
    call_command('loaddata', 'rfprofiles')
    rfconfig.rf_profile = RFProfile.objects.get(name=profile_name)
    rfconfig.save()

class Migration(migrations.Migration):

    dependencies = [
        ('center', '0004_auto_20161108_1219'),
    ]

    operations = [
            migrations.RunPython(load_new_rfprofiles),
    ]
