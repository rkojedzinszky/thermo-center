""" Base for grpcio service classes """

class BaseServicer:
    """ A base class with empty start() and shutdown() methods """

    def start(self, server):
        """ Here the servicer should register itself """
        raise NotImplementedError()

    def shutdown(self):
        """ Empty shutdown """
        pass
