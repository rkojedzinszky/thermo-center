""" Carbon pickle feed client """

import socket
import asyncio
import pickle
import struct
from center.dropqueue import DropQueue

class Client:
    """ Base for carbon clients """
    def __init__(self, loop, endpoint, maxsize):
        self.loop = loop
        self._endpoint = endpoint
        self._socket = None
        self.queue = DropQueue(maxsize=maxsize)
        self._task = None

    async def _feed(self):
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

    def start(self):
        """ Start task """
        self._task = self.loop.create_task(self._feed())

    def stop(self):
        """ Stop task """
        self._task.cancel()


class LineClient(Client):
    """ Carbon feeder using the simple Line protocol """
    def send(self, data):
        payload = ''.join([
                '{} {} {}\n'.format(r[0], r[1][1], r[1][0])
                for r in data]
                )
        payload = payload.encode()
        self.queue.dput(payload)


class PickleClient(Client):
    """ Carbon feeder using the Pickle protocol """
    def send(self, data):
        payload = pickle.dumps(data, protocol=2)
        header = struct.pack('!L', len(payload))
        self.queue.dput(header + payload)

