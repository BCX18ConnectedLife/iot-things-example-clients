package main

import (
	"log"
	"github.com/BCX18ConnectedLife/iot-things-example-clients/go/things/client"
	"github.com/BCX18ConnectedLife/iot-things-example-clients/go/things/ws"
	"github.com/BCX18ConnectedLife/iot-things-example-clients/go/things/examples"
	"github.com/BCX18ConnectedLife/iot-things-example-clients/go/things"
)

func main() {
	cfg := &client.Configuration{
		//SkipSslVerify: true,
		//Proxy: "http://localhost:3128",
	}

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

	t := things.NewThing()
	t.ThingId = "BCX18:" + uuid.Must(uuid.NewV4()).String()
	t.Attributes["name"] = "NameAttribute"

	t, err = conn.Add(t)
	if err != nil || t == nil {
		panic(err.Error())
	}
	log.Println("Thing created. id:", t.ThingId)

	thingId := t.ThingId
	t, err = conn.Get(thingId)
	if err != nil || t == nil {
		panic(err.Error())
	}

	if t.ThingId != thingId {
		panic("Unequal thing ID returned")
	}
	log.Println("Got back thing. id:", t.ThingId)

	t.Attributes["prop"] = "val"
	err = conn.Update(t)
	if err != nil {
		panic(err.Error())
	}

	log.Println("Thing updated.")

	t, err = conn.Get(thingId)
	if err != nil || t == nil {
		panic(err.Error())
	}

	if t.Attributes["prop"] != "val" {
		log.Println(t)
		panic("Property 'prop' is not of value 'val'")
	}

	err = conn.Delete(thingId)
	if err != nil {
		panic(err.Error())
	}

	log.Println("Deleted thing")

	t, err = conn.Get(thingId)
	if t != nil {
		panic("Thing should have been deleted.")
	}

	log.Println("CRUD test completed")
}
