
class TokenBucketFilter:
    """ A Token Bucket Filter """
    def __init__(self, rate, burst, timefunc):
        self._rate = rate
        self._burst = burst
        self._cap = burst
        self._timefunc = timefunc
        self._now = self._timefunc()

    def _replenish(self):
        now = self._timefunc()
        self._cap += (now - self._now) * self._rate
        self._cap = min(self._cap, self._burst)
        self._now = now

    def feed(self, amount):
        self._replenish()
        if amount <= self._cap:
            self._cap -= amount
            return True
        return False

    def reset(self):
        self._cap = self._burst
        self._now = self._timefunc()

    @property
    def capacity(self):
        self._replenish()
        return self._cap
