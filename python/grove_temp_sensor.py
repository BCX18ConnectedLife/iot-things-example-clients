# To 'mock' the calls to grovepi script, change this line to 'import grovepi_mock as grovepi'
# import grovepi
from datetime import datetime

import grovepi_mock as grovepi


class TemperatureHumiditySensor:
    """A simple abstraction for using the grovepi temperature and humidity sensor"""
    port = None
    lastUpdate = None
    # 0 = blue sensor, 1 = white sensor
    module_type = 0

    def __init__(self, digitalPort, samplingRate=1):
        self.port = digitalPort
        self.samplingRate = samplingRate

    def get_temperature_and_humidity(self):
        self.lastUpdate = datetime.now().__str__()
        return grovepi.dht(self.port, self.module_type)

    def get_sampling_rate(self):
        return self.samplingRate

    def get_last_update(self):
        return self.lastUpdate

    def set_sampling_rate(self, samplingRate):
        self.samplingRate = samplingRate

    def get_properties_json(self, temperature, humidity):
        return {
            "temperatureValue": temperature,
            "humidityValue": humidity,
            "lastUpdate": self.get_last_update(),
            "samplingRate": self.get_sampling_rate()
        }
