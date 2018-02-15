# To 'mock' the calls to grovepi script, change this line to 'import grovepi_mock as grovepi'
import grovepi_mock as grovepi
# import grovepi


class Buzzer:
    """A simple abstraction for using the grovepi buzzer"""
    port = None

    def __init__(self, digitalPort, enabled=False):
        self.port = digitalPort
        self.enabled = enabled
        grovepi.pinMode(self.port, "OUTPUT")

    def is_enabled(self):
        return self.enabled

    def set_enabled(self, enabled):
        self.enabled = enabled
        int_to_write = 1 if self.enabled else 0
        grovepi.digitalWrite(self.port, int_to_write)
