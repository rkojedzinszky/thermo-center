
from center.models import Sensor
from django.db import models

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

