""" Carbon pickle feed client """

import socket
import asyncio
from lib import aiothread
from aggregator import dropqueue

class LineClient(aiothread.AIOThread):
    """ Base for carbon clients """
    def __init__(self, endpoint, maxsize):
        super().__init__()

        self.endpoint = endpoint
        self.maxsize = maxsize
        self.queue = None

    async def arun(self):
        self.queue = dropqueue.DropQueue(maxsize=self.maxsize, loop=self.loop)
        while True:
            sock = None
            try:
                sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
                sock.setblocking(False)
                await self.loop.sock_connect(sock, self.endpoint)
                while True:
                    await self.loop.sock_sendall(sock, await self.queue.get())
            except asyncio.CancelledError:
                break
            except Exception as exc:  # pylint: disable=broad-except
                print(exc)
                await asyncio.sleep(1)
            finally:
                sock.close()

    async def _send(self, payload):
        self.queue.dput(payload)

    def send(self, data):
        """ Enqueue data to be sent """

        if not self.queue:
            return

        payload = ''.join(
            ['{} {} {}\n'.format(r[0], r[1][0], r[1][1])
             for r in data]
        )
        payload = payload.encode()
        asyncio.run_coroutine_threadsafe(self._send(payload), self.loop)
