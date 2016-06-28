# -*- coding: utf-8 -*-
from __future__ import unicode_literals

from django.db import migrations, models

def create_rfconfig(apps, schema_editor):
    import random, os
    generator = random.SystemRandom()

    apps.get_model('center', 'RFConfig').objects.create(
        pk=1,
        rf_channel=generator.randrange(256),
        rf_profile=apps.get_model('center', 'RFProfile').objects.get(pk=1),
        network_id=generator.randrange(65536),
        aes_key=''.join('%02x' % ord(c) for c in os.urandom(16)),
    )

def remove_rfconfig(apps, schema_editor):
    apps.get_model('center', 'RFConfig').objects.all().delete()

class Migration(migrations.Migration):

    dependencies = [
        ('center', '0002_rfprofiles'),
    ]

    operations = [
            migrations.RunPython(create_rfconfig, reverse_code=remove_rfconfig),
    ]
