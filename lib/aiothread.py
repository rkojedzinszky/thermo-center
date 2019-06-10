""" Asyncio loop in a separate thread """

import threading
import asyncio

class AIOThread(threading.Thread):
    def __init__(self):
        super().__init__()

        self.loop = None
        self._arunTask = None

    async def _cancel(self):
        self._arunTask.cancel()

    def run(self):
        self.loop = asyncio.new_event_loop()

        self.init()

        self._arunTask = self.loop.create_task(self.arun())

        try:
            self.loop.run_until_complete(self._arunTask)
        except asyncio.CancelledError:
            pass  # noqa
        finally:
            self.deinit()
            self.loop.close()

    def init(self):
        """ Can be overridden, defaults to nothing """
        pass

    def deinit(self):
        """ Can be overridden, defaults to nothing """
        pass

    def cancel(self):
        """ Cancel the thread

        Imitates posix cancel, though the async loop might
        catch the exception and simply ignore it
        """
        if self.loop:
            self.loop.call_soon_threadsafe(self._cancel())

    async def arun(self):
        raise NotImplementedError()
