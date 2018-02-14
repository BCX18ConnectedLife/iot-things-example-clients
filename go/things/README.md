# Bosch IoT Things API for Go.

Created by [Zubair Hamed](https://github.com/zubairhamed).

This is a lightweight, simple to use, low memory footprint but fast client for IoT Things.
This includes both REST and WebSockets-based APIs.

## Installing/Getting
Install Go 1.7.x on your platform of choice. Go supports cross-compiling its binaries to other platforms natively.
No cross-compilation tools are required. 

After installing Go, you can get this SDK package by the following command in terminal/command line:

go get github.com/BCX18ConnectedLife/iot-things-example-clients/go

This would download the package into your GOPATH

## Compiling the examples

The Things Go SDK has been tested on the following platforms.

### Raspberry Pi/Intel Edison/Beaglebone
env GOOS=linux GOARCH=arm go build

### Mindstorms EV3 with EV3Dev
env GOARCH=arm GOOS=linux GOARM=5 go build

### FisherTechnik MicroController

### C.H.I.P (Next Thing Co.)

## Usage
### REST API
``` go
	conn, err := rest.Dial(
		examples.ENDPOINT_URL_REST,
		examples.USERNAME,
		examples.PASSWORD,
		examples.APITOKEN,
		nil,
	)

	// Create a new Thing Instance
	t := things.NewThing()
	t.Attributes["name"] = "NameAttribute"
	t, err = conn.Add(t)
	if err != nil {
		panic(err.Error())
	}
	log.Println("Thing created. id:", t.ThingId)
```


### Websockets API
``` go
	conn, err := ws.Dial(
		examples.ENDPOINT_URL_WS,
		examples.USERNAME,
		examples.PASSWORD,
		examples.APITOKEN,
		nil,
	)

	if err != nil {
		panic(err.Error())
	}

	t := things.NewThing()
	t.ThingId = "BCX18:" + uuid.Must(uuid.NewV4()).String()
	t.Attributes["name"] = "NameAttribute"

	t, err = conn.Add(t)
	if err != nil || t == nil {
		panic(err.Error())
	}
	log.Println("Thing created. id:", t.ThingId)
```

### Using WebSockets API for listening to Events
``` go
	conn, err := ws.Dial(
		examples.ENDPOINT_URL_WS,
		examples.USERNAME,
		examples.PASSWORD,
		examples.APITOKEN,
		nil,
	)

	if err != nil {
		panic(err.Error())
	}

	// Create Asynchronous channel for receiving events
	obsEvents := make(chan *things.WSMessage)
	for {
		select {
		// If we get an event from Things
		case obsMsg, _ := <-obsEvents:
			log.Println(">> Event", obsMsg.Topic)
		}
	}
```

### Configuration
An optional configuration object can be passed in order to configure connection parameters (e.g. Proxies)
``` go
	cfg := &client.Configuration{
		SkipSslVerify: true,
		Proxy: "http://my.proxy:3128",
	}
```

### Examples
#### rest-crud
Shows examples of how to use REST to Create, Read, Update and Delete on IoT Things

#### ws-crud
 Same as the rest-crud example, except it uses the Websockets API

#### ws-events
 An example how to subscribe to events coming from the IoT Things service, namely Additions, Creations and Deletions of Thing instances.

#### things-cli
A command-line interface which can be used with the IoT Things service. Useful in combination with Shell scripting.

#### forwarders
Bunch of examples on how to write persisters (e.g. Subscribing and persisting historical data from Things)

#### proxies
Bunch of examples on how to write sync services which bridges between IoT Things and other services/protocols etc
