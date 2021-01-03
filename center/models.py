""" thermo-center main models """

import re
import logging
import json
import time
import random
import os
from django.conf import settings
from django.db import models
from application.cache import cache
import center.fields

logger = logging.getLogger(__name__)  # pylint: disable=invalid-name

_METRICS_CACHE_TIMEOUT = 120


class RFProfile(models.Model):
    """ A profile for RF communication """
    name = models.CharField(max_length=50)
    confregs = models.CharField(max_length=128)

    class Meta:  # pylint: disable=too-few-public-methods,missing-docstring
        ordering = ['pk']

    def __str__(self):
        return '%s' % self.name


class RFConfig(models.Model):
    """ The current RF configuration """
    rf_channel = center.fields.RangedIntegerField(min_value=0, max_value=255)
    rf_profile = models.ForeignKey(RFProfile, on_delete=models.CASCADE)
    network_id = center.fields.RangedIntegerField(min_value=0, max_value=65535)
    aes_key = models.CharField(max_length=32)

    def __str__(self):
        return 'RFConfig'

    def _generate_config(self):
        """ Initialise a new RF envronment """
        generator = random.SystemRandom()

        self.rf_channel = generator.randrange(256)
        self.network_id = generator.randrange(65536)
        self.aes_key = ''.join('{:02x}'.format(c) for c in os.urandom(16))


class Sensor(models.Model):
    """ A sensor device """
    id = center.fields.SensorIdField(primary_key=True)
    name = models.CharField(max_length=100, blank=True)
    last_seq = models.PositiveIntegerField(null=True)
    last_tsf = models.FloatField(null=True)

    def __str__(self):
        return '{} ({:02x})'.format(self.name or 'NONAME', self.id)

    def _cache_key(self):
        return 'sensor.{:02x}'.format(self.pk)

    def resync(self):
        """ Resync a sensor, when a battery change or rarely a time
        synchronization error occured.
        """

        self.last_seq = None
        self.last_tsf = time.time()
        self.set_cache({'valid': False})
        self.save()

    def get_cache(self):
        """ Retrieve metrics stored in cache """
        data = cache.get(self._cache_key())

        if data:
            return json.loads(data)

        return {}

    def set_cache(self, values):
        """ Saves values in cache """
        data = json.dumps(values, check_circular=False, separators=(',', ':'))

        cache.set(self._cache_key(), data, time=_METRICS_CACHE_TIMEOUT)


class SensorResync(models.Model):
    """ Represents a resync event
    """
    sensor = models.ForeignKey(Sensor, on_delete=models.CASCADE)
    ts = models.DateTimeField(auto_now_add=True)

    def save(self, **kwargs):
        if self.pk is None:
            self.sensor.resync()

        return super().save(**kwargs)


class ConfigureSensorTask(models.Model):
    """ Represents a sensor configuration task

    Once created, a ConfigureRequest should be called on a receiver
    with pk. The receiver will update status on this.
    """

    sensor = models.ForeignKey(Sensor, on_delete=models.CASCADE)
    created = models.DateTimeField(auto_now_add=True)
    started = models.DateTimeField(null=True)
    first_discovery = models.DateTimeField(null=True)
    last_discovery = models.DateTimeField(null=True)
    finished = models.DateTimeField(null=True)
    error = models.CharField(max_length=100, blank=True, null=True)
