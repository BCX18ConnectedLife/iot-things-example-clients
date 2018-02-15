# To 'mock' the calls to grovepi script, change this line to 'import grovepi_mock as grovepi'
# import grovepi
from datetime import datetime

import grovepi_mock as grovepi


class LightSensor:
    """A simple abstraction for using the grovepi light sensor"""
    port = None
    lastUpdate = None

    def __init__(self, analogPort, samplingRate=1):
        self.port = analogPort
        self.samplingRate = samplingRate
        grovepi.pinMode(self.port, "INPUT")

    def get_illumination(self):
        self.lastUpdate = datetime.now().__str__()
        return grovepi.analogRead(self.port)

    def get_sampling_rate(self):
        return self.samplingRate

    def get_last_update(self):
        return self.lastUpdate

    def set_sampling_rate(self, samplingRate):
        self.samplingRate = samplingRate

    def get_properties_json(self, illumination):
        return {
            "sensorValue": illumination,
            "lastUpdate": self.get_last_update(),
            "samplingRate": self.get_sampling_rate()
        }
