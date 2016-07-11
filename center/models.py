import re
import logging
from django.conf import settings
from django.db import models
from django.utils import timezone
from django.core.cache import cache
import center.fields

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

        if avg is None:
            return

        logger.info('%s: update: seq=%d' % (self, seq))

        self.save(update_fields=('last_seq', 'last_ts'))

        m = {m.metric: m.value() for m in metrics}
        m['intvl'] = avg

        if carbon:
            ts = int(self.last_ts.strftime('%s'))
            carbon_data = [('%s.%s' % (self._carbon_path(), k), (ts, v)) for k, v in m.iteritems()]
            carbon.send(carbon_data)

        m['last_seq'] = self.last_seq
        m['last_ts'] = self.last_ts
        cache.set(self._carbon_path(), m)

    def resync(self):
        """ Resync a sensor, when a battery change or rarely a time
        synchronization error occured.
        """

        self.last_seq = None
        self.last_ts = timezone.now()
        self.save()
