""" PID controller module """

import time

class PID:
    """ A tuned PID controller, which can limit the
    integral part in a range """

    class Error:
        """ Represents an error in a time point """
        __slots__ = ('_ts', '_error')

        def __init__(self, error: float, t=None):
            self._ts = t or time.time()
            self._error = error

        @property
        def ts(self):
            return self._ts

        @property
        def error(self):
            return self._error

        def __str__(self):
            return 'V: %f @ %f' % (self._error, self._ts)

        def to_dict(self) -> dict:
            """ Serialize to dict """
            return {
                't': self._ts,
                'e': self._error,
            }

        @classmethod
        def from_dict(cls, d) -> 'PID.Error':
            """ Deserialize from dict """

            return PID.Error(d['e'], d['t'])

    def __init__(self):
        self._int = 0
        self._last: 'PID.Error' = None
        self._cur: 'PID.Error' = None

    def feed(self, error, intabsmax: int):
        self._last = self._cur
        self._cur = PID.Error(error)

        if self._last:
            newint = self._int + (self._last.error + error) * (self._cur.ts - self._last.ts) / 2

            # handle maximum
            if intabsmax is not None:
                if newint < -intabsmax:
                    newint = -intabsmax
                elif newint > intabsmax:
                    newint = intabsmax

            self._int = newint

    def value(self, kp=1.0, ki=1.0, kd=1.0):
        pv = self._cur.error
        iv = self._int
        dv = 0
        if self._last:
            dv = (self._cur.error - self._last.error) / (self._cur.ts - self._last.ts)

        return kp * pv + kd * dv + ki * iv

    def to_dict(self) -> dict:
        """ Serialize to dict """
        d = {
            'i': self._int,
        }

        if self._last:
            d['l'] = self._last.to_dict()

        if self._cur:
            d['c'] = self._cur.to_dict()

        return d

    @classmethod
    def from_dict(cls, d) -> 'PID':
        """ Deserializa from dict """
        p = PID()

        p._int = d['i']
        if 'l' in d:
            p._last = PID.Error.from_dict(d['l'])

        if 'c' in d:
            p._cur = PID.Error.from_dict(d['c'])

        return p
