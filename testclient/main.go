package main

import (
	"fmt"
	"github.com/nats-io/go-nats"
)

func main() {

	publishAndSubscribe()

    select {}
}

func publishAndSubscribe() {

	nc, _ := nats.Connect("nats://demo.nats.io:4222")

	nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	nc.Publish("foo", []byte("Hello World1"))
	nc.Publish("foo", []byte("Hello World2"))
}