# Generated by Django 2.2.6.dev20191007163800 on 2019-10-07 20:04

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('heatcontrol', '0020_control_intabsmax'),
    ]

    operations = [
        migrations.AlterField(
            model_name='control',
            name='intabsmax',
            field=models.FloatField(blank=True, default=100.0, null=True),
        ),
    ]