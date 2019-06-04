
import os
import select
import asyncio
from . import tbf


class InGPIO:
    """ Class for handling gpio ports
    """
    def __init__(self, path):
        self._path = path
        self._fd = os.open(self._path + '/value', os.O_RDWR)
        self._input()

    def __del__(self):
        if hasattr(self, '_fd'):
            os.close(self._fd)

    def _input(self):
        with open(self._path + '/direction', 'w') as fh:
            fh.write('in')

    def interrupt(self, edge):
        with open(self._path + '/edge', 'w') as fh:
            fh.write('none')
        with open(self._path + '/edge', 'w') as fh:
            fh.write(edge)

    def fileno(self):
        return self._fd

    def value(self):
        os.lseek(self._fd, 0, os.SEEK_SET)
        return os.read(self._fd, 1) == b'1'


class InterruptHandler:
    """ Class for handling interrupts on gpio ports
    """

    class Storm(RuntimeError):
        """ Thrown when interrupt storm detected """
        pass

    def __init__(self, loop, gpiopath):
        self._loop = loop
        self._poller = select.epoll()
        self._gpio = InGPIO(gpiopath)
        self._gpio.interrupt('rising')
        self._poller.register(self._gpio.fileno(), select.POLLPRI)
        self._interrupt = asyncio.Event(loop=loop)
        self._tbf = tbf.TokenBucketFilter(5, 60, loop.time)

    def _onread(self):
        self._poller.poll()
        self._interrupt.set()

    def _enable(self):
        self._loop.add_reader(self._poller.fileno(), self._onread)

    def _disable(self):
        self._loop.remove_reader(self._poller.fileno())

    async def _waitforinterrupt(self):
        while True:
            await asyncio.sleep(0)
            if self._gpio.value():
                return
            self._interrupt.clear()
            self._enable()
            try:
                await self._interrupt.wait()
            finally:
                self._disable()

    async def waitforinterrupt(self):
        await self._waitforinterrupt()
        if not self._tbf.feed(1):
            raise InterruptHandler.Storm()

    def resettbf(self):
        self._tbf.reset()
