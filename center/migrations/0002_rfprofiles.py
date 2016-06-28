# -*- coding: utf-8 -*-
from __future__ import unicode_literals

from django.db import migrations
from django.core.management import call_command

def load_rfprofiles(apps, schema_editor):
    call_command('loaddata', 'rfprofiles')

def unload_rfprofiles(apps, schema_editor):
    apps.get_model('center', 'RFProfile').objects.all().delete()

class Migration(migrations.Migration):

    dependencies = [
        ('center', '0001_initial'),
    ]

    operations = [
        migrations.RunPython(load_rfprofiles, reverse_code=unload_rfprofiles),
    ]
