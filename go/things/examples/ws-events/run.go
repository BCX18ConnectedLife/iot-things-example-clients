package main

import (
	"github.com/BCX18ConnectedLife/iot-things-example-clients/go/things/client"
	"github.com/BCX18ConnectedLife/iot-things-example-clients/go/things/ws"
	"github.com/BCX18ConnectedLife/iot-things-example-clients/go/things/examples"
	"github.com/BCX18ConnectedLife/iot-things-example-clients/go/things"
	"time"
	"fmt"
)

func main() {
	cfg := &client.Configuration{
		//SkipSslVerify: true,
		//Proxy: "http://localhost:3128",
	}

	fmt.Println("### Connecting to Things via WebSockets..")
	conn, err := ws.Dial(
		examples.ENDPOINT_URL_WS,
		examples.USERNAME,
		examples.PASSWORD,
		examples.APITOKEN,
		cfg,
	)

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("### Connected.")

	obsEvents := make(chan *things.WSMessage)

	conn.ObserveEvents(obsEvents)
	fmt.Println("### Start Observing all events")

	tickChan := time.NewTicker(time.Second * 60).C

	for {
		select {
		case obsMsg, _ := <-obsEvents:
			fmt.Println(">> Incoming Event       ", obsMsg.Topic)

		case <- tickChan:
			common.PrintMemoryStats()
		}
	}
}

