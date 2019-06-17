""" dropqueue """

import asyncio

class DropQueue(asyncio.Queue):
    """ Queue which drops latest elements when full """
    def dput(self, item):
        if self.full():
            self.get_nowait()
            self.task_done()
        self.put_nowait(item)
