
from twisted.internet import reactor
from center.receiver import RadioBase

class Configurator(RadioBase):
    def __del__(self):
        print 'Configurator destroyed'

    def run(self):
        reactor.callLater(15, self._main.startreceiver)
