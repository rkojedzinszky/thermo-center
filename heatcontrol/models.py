
import datetime
from center.models import Sensor
from django.db import models, IntegrityError
from django.core.exceptions import ValidationError

class DayType(models.Model):
    """ Represents day types """
    name = models.CharField(max_length=50)

    def __str__(self):
        return 'DayType %s' % self.name

class Calendar(models.Model):
    """ A calendar, map each day to a daytype """
    day = models.DateField(unique=True)
    daytype = models.ForeignKey(DayType)

    def __str__(self):
        return '%s' % self.day

class DayTime(models.Model):
    """ Divides the day into parts, for which different settings can be given """
    daytype = models.ForeignKey(DayType)
    start = models.TimeField()
    end = models.TimeField()

    def __str__(self):
        return '%s[%s-%s]' % (self.daytype, self.start, self.end)

    def save(self, *args, **kwargs):
        if self.end != datetime.time(0, 0) and self.end < self.start:
            raise ValidationError()

        qs = DayTime.objects.filter(daytype=self.daytype).filter(models.Q(end='00:00:00') | models.Q(end__gt=self.start))
        if self.end != datetime.time(0, 0):
            qs = qs.filter(start__lt=self.end)

        if self.pk is not None:
            qs = qs.exclude(pk=self.pk)

        if qs.exists():
            raise IntegrityError()

        return super(DayTime, self).save(*args, **kwargs)

class HeatSensor(models.Model):
    """ A specific target temperature setting for a sensor for a given daytime """
    sensor = models.ForeignKey(Sensor, on_delete=models.CASCADE)
    daytime = models.ForeignKey(DayTime, on_delete=models.CASCADE)
    target_temp = models.FloatField()

    class Meta:
        unique_together = (
                ('sensor', 'daytime'),
                )

    def __str__(self):
        return '%s at %s: %f' % (self.sensor, self.daytime, self.target_temp)

class HeatSensorOverride(models.Model):
    """ Simply override a setting for a period of time for a sensor """
    sensor = models.ForeignKey(Sensor, on_delete=models.CASCADE)
    start = models.DateTimeField()
    end = models.DateTimeField()
    target_temp = models.FloatField()

    class Meta:
        index_together = (
                ('sensor', 'end'),
                )
