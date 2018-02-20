package examples

import (
	"github.com/BCX18ConnectedLife/iot-things-example-clients/go/things/client"
)

// For IoT Things
var ENDPOINT_URL_REST = "https://things.s-apps.de1.bosch-iot-cloud.com"
var ENDPOINT_URL_WS = "wss://things.s-apps.de1.bosch-iot-cloud.com"
var USERNAME = "bcx18"
var PASSWORD = "bcx18!Open2"
var APITOKEN = "db7f4e0cca344d32be72914311f1055f"
var DEFAULT_CLIENT_CONFIG = &client.Configuration{
	//SkipSslVerify: true,
}
