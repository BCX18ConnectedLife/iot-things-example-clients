package main

import (
	"log"
	"github.com/BCX18ConnectedLife/iot-things-example-clients/go/things/examples"
	"github.com/BCX18ConnectedLife/iot-things-example-clients/go/things"
	"github.com/BCX18ConnectedLife/iot-things-example-clients/go/things/client"
	"github.com/BCX18ConnectedLife/iot-things-example-clients/go/things/rest"
)

func main() {
	// Use Proxy
	cfg := &client.Configuration{
		//SkipSslVerify: true,
		//Proxy: "http://localhost:3128",
	}

	// Get an instance of a rest connection to Things
	conn, err := rest.Dial(
		examples.ENDPOINT_URL_REST,
		examples.USERNAME,
		examples.PASSWORD,
		examples.APITOKEN,
		cfg,
	)

	if err != nil {
		panic(err.Error())
	}

	// Create a new Thing Instance
	t := things.NewThing()
	t.Attributes["name"] = "NameAttribute"
	t, err = conn.Add(t)
	if err != nil {
		panic(err.Error())
	}
	log.Println("Thing created. id:", t.ThingId)

	// Verify if Thing instance was really created
	thingId := t.ThingId
	ts, err := conn.Get(thingId)
	if err != nil {
		panic(err.Error())
	}

	if len(ts) != 1 {
		panic(err.Error())
	}

	t = ts[0]
	if t.ThingId != thingId {
		panic("Unequal thing ID returned")
	}
	log.Println("Got back Thing. id:", t.ThingId)

	// Update a Thing Instance
	t.Attributes["prop"] = "val"
	err = conn.Update(t)
	if err != nil {
		panic(err.Error())
	}
	log.Println("Thing updated.")

	// Verify if we got back the value right after updating
	ts, err = conn.Get(thingId)
	if err != nil {
		panic(err.Error())
	}

	if len(ts) != 1 {
		panic(err.Error())
	}
	t = ts[0]

	if t.Attributes["prop"] != "val" {
		log.Println(t)
		panic("Property 'prop' is not of value 'val'")
	}

	// Finally destroy the thing
	err = conn.Delete(thingId)
	if err != nil {
		panic(err.Error())
	}
	log.Println("Deleted Thing")

	// Verify if the thing's really kaput.
	ts, err = conn.Get(thingId)
	if err != nil {
		panic(err.Error())
	}

	if len(ts) != 0 {
		panic("Should be 0 Things returned after deletion")
	}

	log.Println("CRUD test completed")
}
