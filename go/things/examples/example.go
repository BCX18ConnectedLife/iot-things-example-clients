package examples

import (
	"github.com/BCX18ConnectedLife/iot-things-example-clients/go/things/client"
)

// For IoT Things
var ENDPOINT_URL_REST = "https://things.s-apps.de1.bosch-iot-cloud.com"
var ENDPOINT_URL_WS = "wss://things.s-apps.de1.bosch-iot-cloud.com"
var USERNAME = "TODO-insert-user"
var PASSWORD = "TODO-insert-password"
var APITOKEN = "TODO-insert-apitoken"
var DEFAULT_CLIENT_CONFIG = &client.Configuration{
	//SkipSslVerify: true,
}
