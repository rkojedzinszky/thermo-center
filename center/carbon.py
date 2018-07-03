""" Carbon pickle feed client """

import socket
import asyncio
import pickle
import struct
from center.dropqueue import DropQueue

class PickleClient:
    def __init__(self, loop, endpoint, maxsize):
        self.loop = loop
        self._endpoint = endpoint
        self._socket = None
        self.queue = DropQueue(maxsize=maxsize)

    async def feed(self):
        while True:
            sock = None
            try:
                sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
                sock.setblocking(False)
                await self.loop.sock_connect(sock, self._endpoint)
                while True:
                    await self.loop.sock_sendall(sock, await self.queue.get())
            except asyncio.CancelledError:
                break
            except:
                await asyncio.sleep(1)
            finally:
                sock.close()

    def send(self, data):
        payload = pickle.dumps(data, protocol=2)
        header = struct.pack('!L', len(payload))
        self.queue.dput(header + payload)
