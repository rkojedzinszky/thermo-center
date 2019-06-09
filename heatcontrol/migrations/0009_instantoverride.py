# -*- coding: utf-8 -*-
from __future__ import unicode_literals

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('heatcontrol', '0008_auto_20170424_2052'),
    ]

    operations = [
        migrations.CreateModel(
            name='InstantProfile',
            fields=[
                ('id', models.AutoField(verbose_name='ID', primary_key=True, serialize=False, auto_created=True)),
                ('name', models.CharField(max_length=50)),
                ('active', models.BooleanField(default=False)),
            ],
        ),
        migrations.CreateModel(
            name='InstantProfileEntry',
            fields=[
                ('id', models.AutoField(verbose_name='ID', primary_key=True, serialize=False, auto_created=True)),
                ('target_temp', models.FloatField()),
                ('control', models.ForeignKey(to='heatcontrol.Control', on_delete=models.CASCADE)),
                ('profile', models.ForeignKey(to='heatcontrol.InstantProfile', on_delete=models.CASCADE)),
            ],
        ),
        migrations.AlterUniqueTogether(
            name='instantprofileentry',
            unique_together=set([('profile', 'control')]),
        ),
        migrations.CreateModel(
            name='InstantOverride',
            fields=[
                ('id', models.AutoField(verbose_name='ID', primary_key=True, serialize=False, auto_created=True)),
                ('target_temp', models.FloatField()),
                ('control', models.OneToOneField(to='heatcontrol.Control', on_delete=models.CASCADE)),
                ('profile', models.ForeignKey(to='heatcontrol.InstantProfile', on_delete=models.CASCADE)),
            ],
        ),
        migrations.RunSQL(
            sql='alter table heatcontrol_instantoverride add constraint "profile_control_fk" foreign key (profile_id, control_id) references heatcontrol_instantprofileentry (profile_id, control_id) on delete cascade deferrable initially deferred',
            reverse_sql='alter table heatcontrol_instantoverride drop constraint "profile_control_fk"',
        ),
    ]
