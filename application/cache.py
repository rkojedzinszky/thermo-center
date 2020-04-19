""" Cache module, which directly interacts with Memcached """

import pylibmc
import retry
from django.conf import settings

class _Cache:
    """ Cache class that uses pylibmc pooling to access Memcached """
    def __init__(self):
        self.pool = pylibmc.ThreadMappedPool(
            pylibmc.Client(
                servers=['{}:{}'.format(settings.MEMCACHED_HOST, settings.MEMCACHED_PORT)],
                behaviors={"tcp_keepalive": True, "tcp_nodelay": True}
            ),
        )

    @retry.retry(exceptions=pylibmc.ConnectionError, tries=3, jitter=1, logger=None)
    def set(self, key: str, value: str, time: int=0) -> bool:
        """ Set a key in the cache """
        with self.pool.reserve() as mc:
            return mc.set(key, value, time=time)

    @retry.retry(exceptions=pylibmc.ConnectionError, tries=3, jitter=1, logger=None)
    def get(self, key: str, default: str=None) -> str:
        """ Get a key from the cache """
        with self.pool.reserve() as mc:
            return mc.get(key, default)

cache = _Cache()
