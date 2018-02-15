#!/usr/bin/python3

import json
import math
import time
import websocket

import raspberry_thing

IOT_THINGS_HOSTNAME = "things.s-apps.de1.bosch-iot-cloud.com"
websocketOpen = False
thing = raspberry_thing.RaspberryDemoThing()


def on_new_illumination_value(illumination):
    if websocketOpen and math.isnan(illumination) == False:
        send_modify_message(thing.create_illumination_change_message(illumination))


def on_new_temperature_value(temperature, humidity):
    if websocketOpen and math.isnan(temperature) == False and math.isnan(humidity) == False:
        send_modify_message(thing.create_temperature_change_message(temperature, humidity))


def on_message(ws, message):
    if message.startswith("{"):
        thing.handle_websocket_message(json.loads(message))
    else:
        print('Received: {}'.format(message))


def send_modify_message(message):
    # convert to JSON
    json_message = json.dumps(message)
    # send via websocket
    ws.send(json_message)


def on_error(ws, error):
    print('An unexpected error happened while using the WebSocket connection: {}'.format(error))


def on_close(ws):
    print('WebSocket closed - trying to reconnect any second.')
    global websocketOpen
    websocketOpen = False
    time.sleep(5)
    start_websocket()


def on_open(ws):
    print("### WebSocket opened ###")
    global websocketOpen
    websocketOpen = True
    # start listening for events and messages
    ws.send("START-SEND-MESSAGES")
    ws.send("START-SEND-EVENTS")


def start_websocket():
    print('Establishing WebSocket connection ...')
    ws_address = "wss://" + IOT_THINGS_HOSTNAME + "/ws/2"
    basic_auth = 'Authorization: Basic {}'.format(raspberry_thing.get_b64_auth())
    api_token = 'x-cr-api-token: {}'.format(raspberry_thing.get_api_token())
    global ws
    # websocket.enableTrace(True)
    ws = websocket.WebSocketApp(ws_address,
                                header=[basic_auth,api_token],
                                on_message=on_message,
                                on_error=on_error,
                                on_close=on_close)
    ws.on_open = on_open
    ws.run_forever()
    # ws.run_forever(http_proxy_host="localhost",
    #                http_proxy_port=3128)


if __name__ == "__main__":
    # init our raspberry thing
    thing.start_polling_illumination(on_new_illumination_value)
    thing.start_polling_temperatures(on_new_temperature_value)

    # start websocket
    start_websocket()
