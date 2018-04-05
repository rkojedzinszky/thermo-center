
import os

class GPIO:
    def __init__(self, path):
        self._path = path
        self._fd = os.open(self._path + '/value', os.O_RDWR)
        if self._fd == -1:
            raise RuntimeError('cannot open gpio')

    def __del__(self):
        os.close(self._fd)

    def fileno(self):
        return self._fd

    def input(self):
        with open(self._path + '/direction', 'w') as fh:
            fh.write('in')

    def interrupt(self, edge):
        with open(self._path + '/edge', 'w') as fh:
            fh.write('none')
        with open(self._path + '/edge', 'w') as fh:
            fh.write(edge)

    def output(self):
        with open(self._path + '/direction', 'w') as fh:
            fh.write('out')

    def value(self):
        os.lseek(self._fd, 0, os.SEEK_SET)
        return os.read(self._fd, 1) == b'1'

