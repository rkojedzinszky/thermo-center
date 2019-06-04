""" Configuration package """

import array
from Crypto.Cipher import AES

class Config:
    """ Represents communication configuration
    - radio config bytes (frequency, modulation)
    - network_id
    - aes key
    """

    def __init__(self, config):
        self.network = config.network
        self.radio_config = list(array.array('B', config.radio_config))
        self.aes_key = config.aes_key
        self.cipher = AES.new(self.aes_key)

