""" Aggregator services """

import socket
import os
import threading
import logging
import time
import random
import json
from django.conf import settings
from django.core.cache import cache
from center import models
from aggregator import api_pb2
from aggregator import api_pb2_grpc
from aggregator import sensorvalue
from aggregator import (mqtt, pid)

PIDCONTROL_INT_ABS_MAX = 100.0


logger = logging.getLogger(__name__)


class PIDControlValue(sensorvalue.Value):
    metric = 'pidcontrol'


class Aggregator(api_pb2_grpc.AggregatorServicer):
    SENSOR_CACHE_LOCK_KEY = 'tc-sensor-{}-lock'
    SENSOR_CACHE_LOCK_TIMEOUT = 15

    def __init__(self):
        super().__init__()

        self.mqtt = None

        if settings.MQTT_HOST:
            self.mqtt = mqtt.MqttClient((settings.MQTT_HOST, settings.MQTT_PORT))
            self.mqtt.start()

    def lock_sensor(self, sensor_id):
        thread_id = '{}-{}-{}'.format(socket.gethostname(), os.getpid(), threading.get_ident())
        key = self.SENSOR_CACHE_LOCK_KEY.format(sensor_id)
        if cache.add(key, thread_id, timeout=self.SENSOR_CACHE_LOCK_TIMEOUT):
            logger.info('Locked sensor {}'.format(sensor_id))
            return True

        holder = cache.get(key)
        logger.info('Locking sensor {} failed: {} holds lock'.format(sensor_id, holder))
        return False

    def _parse_metrics(self, packet):
        """ Prepare metrics dictionary """
        metrics = [sensorvalue.RSSI(packet.rssi), sensorvalue.LQI(packet.lqi)]

        si = iter(packet.raw)
        try:
            while True:
                try:
                    metrics.append(sensorvalue.SensorValueParser.parse(next(si), next(si)))
                except sensorvalue.SensorValueParser.InvalidType:
                    pass
        except StopIteration:
            pass

        mh = {m.metric: m for m in metrics}
        try:
            t = mh['Temperature']
            h = mh['Humidity']
            if (isinstance(t, sensorvalue.HTU21DTemperature)
                    and isinstance(h, sensorvalue.HTU21DHumidty)):
                h.temp_compensate(t)
        except KeyError:
            pass

        return mh

    def FeedSensorPacket(self, packet, context):
        timestamp = time.time()
        if not self.lock_sensor(packet.id):
            return api_pb2.FeedResponse(processed=False)

        try:
            s = models.Sensor.objects.select_related('control').get(id=packet.id)
        except models.Sensor.DoesNotExist:
            logger.warn('Unknown device id: %02x' % id_)
            return api_pb2.FeedResponse(processed=False)

        cachevalues = s.get_cache()
        cachevalues['intvl'] = s.validate_seq(timestamp, packet.seq)
        cachevalues['valid'] = cachevalues['intvl'] is not None
        if not cachevalues['valid']:
            s.set_cache(cachevalues)
            return api_pb2.FeedResponse(processed=False)

        # Save to db using probability
        if random.random() < settings.SENSOR_DB_UPDATE_PROBABILITY:
            s.save(update_fields=('last_seq', 'last_tsf'))

        # Prepare metrics dictionary
        mh = self._parse_metrics(packet)

        # Calculate PID
        if hasattr(s, 'control'):
            pcp = s.control

            try:
                pc = cachevalues['pid']
            except KeyError:
                pc = cachevalues['pid'] = pid.PID()

            target_temp = pcp.get_target_temp()
            if target_temp is not None:
                error = target_temp - mh['Temperature'].value()
                pc.feed(error=error, intabsmax=PIDCONTROL_INT_ABS_MAX)
                pcv = pc.value(kp=pcp.kp, ki=pcp.ki, kd=pcp.kd)
                logger.debug('%s: pid control=%f', s, pcv)
                mh['pidcontrol'] = PIDControlValue(pcv)
        else:
            cachevalues.pop('pid', None)

        # Update cache
        cachevalues.update({m.metric: m.value() for m in mh.values()})

        # Save cache
        s.set_cache(cachevalues)

        # XXX: remove pid to avoid serialization
        cachevalues.pop('pid', None)

        # Publish mh to mqtt, carbon
        if self.mqtt:
            self.mqtt.publish('{}{:02x}/report'.format(settings.MQTT_PREFIX, s.pk),
                json.dumps(cachevalues, separators=(',', ':')).encode())

        return api_pb2.FeedResponse(processed=True)


def add_services(server):
    """ Register services """
    api_pb2_grpc.add_AggregatorServicer_to_server(Aggregator(), server)
