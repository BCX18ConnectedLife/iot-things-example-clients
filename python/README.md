# Bosch IoT Things :: Python example

The python scripts are to be used on a Raspberry Pi with a GrovePi+ board.

They aim to frequently provide and consume property changes of sensor values
and also allow to activate or deactivate the buzzer through a message.

## Prerequisites

* Python 3.x installed on the Raspberry
* Raspberry with installed GrovePi libraries
* Raspberry can connect to the Internet

## Hardware Setup

* GrovePi+
* Raspberry Pi Model B+ v1.2
* Grove Buzzer v1.1b
* Grove Light Sensor v1.0

## Quick start

Install the following python modules:
* websocket-client

You only need to adjust the credentials with which to connect to Bosch IoT THings. These can be found
in `raspberry_thing.py`:

```python
# User and password needed for providing new sensor values
THINGS_USER = "TODO-insert-user"
THINGS_PASSWORD = "TODO-insert-password"
THINGS_API_TOKEN = "TODO-insert-api-token"
# The id of our raspberry Thing
THING_ID = "BCX18/raspberry"
```

To start the script, call it using python3:

```bash
$ python3 ./things_grove_demo.py
```

It will connect to the WebSocket provided by Bosch IoT Things and start sending
sensor values to it.

## Possible adaptions

Read here how to adapt the code if your setup differs from the demo setup.

## Using other digital and analog ports for your sensors
You can change those properties in `raspberry_thing.py`:
```python
# Digital Port D8 on the GrovePi+ is connected to the buzzer
BUZZER_PORT = 8
# Analog Port A0 on the GrovePi+ is connected to the light sensor
LIGHT_SENSOR_PORT = 0
# Digital Port D4 on the GrovePi+ is connected to the temp sensor
TEMP_SENSOR_PORT = 4
```

## Testing the script without a GrovePi+

You can run the scripts without having the GrovePi+ connected to your
Raspberry Pi.

Just replace the import from `import grovepi` inside `raspberry_thing.py`
to the mock import:
```python
# To 'mock' the calls to grovepi script, change this line to 'import grovepi_mock as grovepi'
import grovepi_mock as grovepi
```
