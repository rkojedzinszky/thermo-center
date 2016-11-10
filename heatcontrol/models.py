
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

class HeatControl(models.Model):
    """ Describes PID Control loop coefficients """
    sensor = models.OneToOneField(Sensor, on_delete=models.CASCADE)
    kp = models.FloatField()
    ki = models.FloatField()
    kd = models.FloatField()

    def __str__(self):
        return '%s[Kp=%f,Ki=%f,Kd=%f]' % (self.sensor, self.kp, self.ki, self.kd)

class HeatSensor(models.Model):
    """ A specific target temperature setting for a sensor """
    sensor = models.ForeignKey(Sensor, on_delete=models.CASCADE)
    daytype = models.ForeignKey(DayType, on_delete=models.CASCADE)
    start = models.TimeField()
    end = models.TimeField()
    target_temp = models.FloatField()

    class Meta:
        index_together = (
                ('sensor', 'daytype'),
                )

    def __str__(self):
        return '%s at %s[%s-%s]: %f' % (self.sensor, self.daytype, self.start, self.end, self.target_temp)

    def save(self, *args, **kwargs):
        if self.end != datetime.time(0, 0) and self.end < self.start:
            raise ValidationError()

        qs = HeatSensor.objects.filter(sensor=self.sensor, daytype=self.daytype).filter(models.Q(end='00:00:00') | models.Q(end__gt=self.start))
        if self.end != datetime.time(0, 0):
            qs = qs.filter(start__lt=self.end)

        if self.pk is not None:
            qs = qs.exclude(pk=self.pk)

        if qs.exists():
            raise IntegrityError()

        return super(HeatSensor, self).save(*args, **kwargs)

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
