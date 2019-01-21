import time

class PID:
    """ A tuned PID controller, which can limit the
    integral part in a range """

    class Error:
        """ Represents an error in a time point """
        __slots__ = ('_ts', '_error')

        def __init__(self, error, ts=time.time):
            if callable(ts):
                ts = ts()
            self._ts = ts
            self._error = error

        @property
        def ts(self):
            return self._ts

        @property
        def error(self):
            return self._error

        def __str__(self):
            return 'V: %f @ %f' % (self._error, self._ts)

    def __init__(self):
        self._int = 0
        self._last = None
        self._cur = None

    def feed(self, error, intabsmax=None, ts=time.time):
        self._last = self._cur
        self._cur = PID.Error(error, ts)

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
