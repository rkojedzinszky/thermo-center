""" Carbon pickle feed client """

import socket
import pickle
import struct

class PickleClient:
    def __init__(self, endpoint):
        self._endpoint = endpoint
        self._socket = None

    def _open(self):
        """ open a new connection """
        self._socket = None
        s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        s.connect(self._endpoint)
        self._socket = s

    def send(self, data):
        payload = pickle.dumps(data, protocol=2)
        header = struct.pack('!L', len(payload))
        loop = 2

        while loop > 0:
            if self._socket is not None:
                try:
                    self._socket.send(header + payload)
                    break
                except IOError:
                    pass

            self._open()
            loop -= 1
