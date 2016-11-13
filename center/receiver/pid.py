import time

class PID(object):
    """ A tuned PID controller to accumulate errors on a fix
    time interval """

    DERIV_SPAN = 2

    class Value(object):
        """ Represents a value in a time point """
        __slots__ = ('_ts', '_value')

        def __init__(self, value, ts=time.time):
            if callable(ts):
                ts = ts()
            self._ts = ts
            self._value = value

        @property
        def ts(self):
            return self._ts

        @property
        def value(self):
            return self._value

        def __str__(self):
            return 'V: %f @ %f' % (self._value, self._ts)

    def __init__(self, intvl):
        self._intvl = intvl
        self._values = []

    def feed(self, value, ts=time.time):
        v = PID.Value(value, ts)
        self._values.append(v)
        left = v.ts - self._intvl
        popped = None

        while self._values[0].ts < left:
            popped = self._values.pop(0)

        if popped is not None:
            nv = PID.Value(popped.value + (self._values[0].value - popped.value) / (self._values[0].ts - popped.ts) * (left - popped.ts), left)
            self._values.insert(0, nv)

    def value(self, sp, kp=1.0, ki=1.0, kd=1.0):
        fv = self._values[0]       # very first value
        li = len(self._values) - 1 # last index
        lv = self._values[li]      # last value

        error = sp - lv.value

        accum = sp * (lv.ts - fv.ts)
        for i in range(1, len(self._values)):
            accum -= (self._values[i - 1].value + self._values[i].value) / 2.0 * (self._values[i].ts - self._values[i-1].ts)

        deriv = 0
        if li >= self.DERIV_SPAN:
            a = self._values[li - self.DERIV_SPAN]

            deriv = (a.value - lv.value) / (lv.ts - a.ts)

        return kp * error + ki * accum + kd * deriv
