import re
from django.db import models
import center.fields

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
        regs = [int(c, base=16) for c in re.findall(r'[0-9a-f]{2}', self.rf_profile.confregs)]
        regs.extend([radio.Radio.ConfReg.CHANNR, self.rf_channel])
        return regs

class Sensor(models.Model):
    """ A sensor device """
    device_id = center.fields.SensorIdField(unique=True)
    name = models.CharField(max_length=100, blank=True)
    last_seq = models.PositiveIntegerField(null=True)
    last_ts = models.DateTimeField(null=True)

    def __str__(self):
        return 'Sensor %02x' % self.device_id
