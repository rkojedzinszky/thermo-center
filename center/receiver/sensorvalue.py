""" Package for handling SensorValues

Mirrors https://github.com/rkojedzinszky/thermo-sensor/blob/master/common/include/sensorvalue.hpp
"""

class Value:
    """ Base object representing a Sensor's value """

    def __init__(self, raw):
        """ Process raw-data into normal value """
        self._value = raw

    def value(self):
        return self._value

class UnsignedValue(Value):
    pass

class SignedValue(Value):
    def __init__(self, raw):
        super(SignedValue, self).__init__(raw)
        if self._value & (1 << (13 - 2*self.T)) != 0:
            self._value |= -1 << (14 - 2*self.T)

class HTU21DTemperature(UnsignedValue):
    T = 0
    t = 0
    metric = 'Temperature'

    def __init__(self, raw):
        super(HTU21DTemperature, self).__init__(raw)
        self._value = -46.85 + 175.72 * self._value / 16384.0

class HTU21DHumidty(UnsignedValue):
    T = 1
    t = 0
    metric = 'Humidity'

    def __init__(self, raw):
        super(HTU21DHumidty, self).__init__(raw)
        self._value = -6 + 125 * self._value / 4096.0

    def temp_compensate(self, temp):
        self._value -= 0.15 * (25 - temp.value())

class AM2302Temperature(SignedValue):
    T = 1
    t = 1
    metric = 'Temperature'

    def __init__(self, raw):
        super(AM2302Temperature, self).__init__(raw)
        self._value /= 10.0

class VCC(UnsignedValue):
    T = 2
    t = 0
    metric = 'Power'

    def __init__(self, raw):
        super(VCC, self).__init__(raw)
        if self._value > 0:
            self._value = 1.1 * 1023 / self._value

class AM2302Humidity(UnsignedValue):
    T = 2
    t = 1
    metric = 'Humidity'

    def __init__(self, raw):
        super(AM2302Humidity, self).__init__(raw)
        self._value /= 10.0

class RSSI(SignedValue):
    T = 3
    t = 0
    metric = 'RSSI'

    def __init__(self, raw):
        super(RSSI, self).__init__(raw)
        self._value = self._value / 2.0 - 74

class LQI(UnsignedValue):
    T = 3
    t = 1
    metric = 'LQI'

class SensorValueParser:
    class InvalidType(Exception):
        pass

    def __init__(self):
        self._T = {}

    def parse(self, d0, d1):
        T = 0
        b = 0x80

        while True:
            if d0 & b == 0:
                break
            b >>= 1
            T += 1
            if T == 4:
                raise SensorValueParser.InvalidType()

        t = (d0 >> (6 - 2*T)) & (0xff >> (7 - T))
        d = (d0 & (0x3f >> (2*T))) << 8 | d1

        try:
            return self._T[T][t](d)
        except KeyError:
            raise SensorValueParser.InvalidType()

    def register(self, cls):
        self._T.setdefault(cls.T, {})[cls.t] = cls

SensorValueParser = SensorValueParser()

SensorValueParser.register(HTU21DTemperature)
SensorValueParser.register(HTU21DHumidty)
SensorValueParser.register(AM2302Temperature)
SensorValueParser.register(VCC)
SensorValueParser.register(AM2302Humidity)
SensorValueParser.register(RSSI)
SensorValueParser.register(LQI)
