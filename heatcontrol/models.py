
import datetime
from django.utils import timezone
from django.db import models, IntegrityError
from django.core.exceptions import ValidationError
from center.models import Sensor

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

class Control(models.Model):
    """
    This enables heat-control operation for a sensor

    Describes PID Control loop coefficients
    """
    sensor = models.OneToOneField(Sensor, on_delete=models.CASCADE)
    kp = models.FloatField()
    ki = models.FloatField()
    kd = models.FloatField()

    def get_target_temp(self):
        """
        This calculates the target temperature, taking into account the available overrides
        """
        now = timezone.now()

        hco = self.heatcontroloverride_set.filter(end__gt=now, start__lte=now).order_by('-pk').first()
        if hco is not None:
            return hco.target_temp

        day = now.date()
        tm = now.time()
        hcp = self.heatcontrolprofile_set.filter(daytype__calendar__day=day).filter(models.Q(end='00:00:00') | models.Q(end__gt=tm), start__lte=tm).first()
        if hcp is not None:
            return hcp.target_temp

        return None

    def __str__(self):
        return '%s[Kp=%f,Ki=%f,Kd=%f]' % (self.sensor, self.kp, self.ki, self.kd)

class HeatControlProfile(models.Model):
    """ Profile setting for a Control unit """
    control = models.ForeignKey(Control, on_delete=models.CASCADE)
    daytype = models.ForeignKey(DayType, on_delete=models.CASCADE)
    start = models.TimeField()
    end = models.TimeField()
    target_temp = models.FloatField()

    class Meta:
        index_together = (
                ('control', 'daytype'),
                )

    def __str__(self):
        return '%s at %s[%s-%s]: %f' % (self.control.sensor, self.daytype, self.start, self.end, self.target_temp)

    def save(self, *args, **kwargs):
        if self.end != datetime.time(0, 0) and self.end < self.start:
            raise ValidationError()

        qs = HeatControlProfile.objects.filter(control=self.control, daytype=self.daytype).filter(models.Q(end='00:00:00') | models.Q(end__gt=self.start))
        if self.end != datetime.time(0, 0):
            qs = qs.filter(start__lt=self.end)

        if self.pk is not None:
            qs = qs.exclude(pk=self.pk)

        if qs.exists():
            raise IntegrityError()

        return super(HeatControlProfile, self).save(*args, **kwargs)

class HeatControlOverride(models.Model):
    """ Simply override a setting for a period of time for a Control unit """
    control = models.ForeignKey(Control, on_delete=models.CASCADE)
    start = models.DateTimeField()
    end = models.DateTimeField()
    target_temp = models.FloatField()

    class Meta:
        index_together = (
                ('control', 'end'),
                )
