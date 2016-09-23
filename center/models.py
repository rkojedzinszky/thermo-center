import re
import logging
from django.conf import settings
from django.db import models
from django.utils import timezone
from django.core.cache import cache
import center.fields
import datetime
import pytz
import time

logger = logging.getLogger(__name__)

def _parse_hex(s):
    return [int(c, base=16) for c in re.findall(r'[0-9a-f]{2}', s)]

class RFProfile(models.Model):
    """ A profile for RF communication """
    name = models.CharField(max_length=50)
    confregs = models.CharField(max_length=128)

    class Meta:
        ordering = ['pk']

    def __str__(self):
        return 'RFProfile %s' % self.name

class RFConfig(models.Model):
    """ The current RF configuration """
    rf_channel = center.fields.RangedIntegerField(min_value=0, max_value=255)
    rf_profile = models.ForeignKey(RFProfile)
    network_id = center.fields.RangedIntegerField(min_value=0, max_value=65535)
    aes_key = models.CharField(max_length=32)

    def __str__(self):
        return 'RFConfig'

    def config_bytes(self):
        from center.receiver import radio
        regs = _parse_hex(self.rf_profile.confregs)
        regs.extend([radio.Radio.ConfReg.CHANNR, self.rf_channel])
        return regs

    def aes_bytes(self):
        return _parse_hex(self.aes_key)

class Sensor(models.Model):
    """ A sensor device """
    id = center.fields.SensorIdField(primary_key=True)
    name = models.CharField(max_length=100, blank=True)
    last_seq = models.PositiveIntegerField(null=True)
    last_ts = models.DateTimeField(null=True)

    def __str__(self):
        return 'Sensor %02x' % self.id

    def _validate_seq(self, seq):
        now = timezone.now()
        avg = 0

        if self.last_ts is not None:
            interval = (now - self.last_ts).total_seconds()

            if self.last_seq is None:
                valid = interval <= 34
            else:
                diff = (seq - self.last_seq) & 0x7fffffff
                avg = interval / diff
                valid = 26 <= avg <= 34

            if not valid:
                logger.warn('%s: received invalid update' % self)
                return None

        self.last_seq = seq
        self.last_ts = now

        return avg

    def _carbon_path(self):
        return 'sensor.%02x' % self.pk

    def feed(self, seq, metrics, carbon=None):
        avg = self._validate_seq(seq)
        cachevalues = {'valid': avg is not None}

        if cachevalues['valid']:
            logger.info('%s: update: seq=%d' % (self, seq))

            self.save(update_fields=('last_seq', 'last_ts'))

            cachevalues.update({m.metric: m.value() for m in metrics})
            cachevalues['intvl'] = avg

            if carbon:
                ts = int(time.time())
                carbon_data = [('%s.%s' % (self._carbon_path(), k), (ts, v)) for k, v in cachevalues.iteritems()]
                carbon.send(carbon_data)

            cachevalues['last_seq'] = self.last_seq
            cachevalues['last_ts'] = self.last_ts

        cache.set(self._carbon_path(), cachevalues)

    def resync(self):
        """ Resync a sensor, when a battery change or rarely a time
        synchronization error occured.
        """

        self.last_seq = None
        self.last_ts = timezone.now()
        self.save()

    def _get_heatsensor(self, dt):
        day = dt.date()
        tm = dt.time()

        return self.heatsensor_set.filter(models.Q(daytime__end='00:00:00') | models.Q(daytime__end__gt=tm), daytime__start__lte=tm, daytime__daytype=heatcontrol.models.Calendar.objects.get(day=day).daytype).first()

    def _get_heatvalues(self):
        """ Extract last 15 minutes data from carbon """
        from graphite.render.evaluator import evaluateTarget
        now = datetime.datetime.utcnow().replace(second=0, microsecond=0) - datetime.timedelta(minutes=1)
        now = now.replace(tzinfo=pytz.utc)
        start = now - datetime.timedelta(minutes=15)
        end = now
        rc = {
            'startTime': start,
            'endTime': end,
            'now': now,
            'localOnly': False,
            'data': []
        }

        return evaluateTarget(rc, '%s.%s' % (self._carbon_path(), 'Temperature'))[0]

    def _get_heatcontrol_pid(self, target, kp=1.0, ki=1.0, kd=1.0):
        ts = list(self._get_heatvalues())
        i = 0
        while i < len(ts) and ts[i] is None:
            i += 1

        if i == len(ts):
            return None

        for j in range(i):
            ts[j] = ts[i]

        for j in range(i+1, len(ts)):
            if ts[j] is None:
                ts[j] = ts[i]
            else:
                i = j

        et = target - ts[-1:][0]
        it = sum([target - i for i in ts])
        dt = ts[-2:-1][0] - ts[-1:][0]

        return kp * et + ki * it + kd * dt

    def get_heatcontrol_pid(self):
        now = timezone.now()
        target_temp = None

        ho = self.heatsensoroverride_set.filter(end__gt=now, start__lte=now).first()
        if ho is not None:
            target_temp = ho.target_temp

        if target_temp is None:
            hs = self._get_heatsensor(now)
            if hs is not None:
                target_temp = hs.target_temp

        if target_temp is None:
            return None

        return self._get_heatcontrol_pid(target_temp)

import heatcontrol.models
