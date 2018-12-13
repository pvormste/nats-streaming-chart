package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/nats-io/go-nats"
)

func main() {
	subscribe()
	select {}
}

func subscribe() {
	serverURL := flag.String("server", "http://localhost:4222", "url to the nats streaming server")
	channel := flag.String("channel", "demo", "channel on which the consumer will subscribe")
	flag.Parse()

	if serverURL == nil || *serverURL == "" {
		log.Fatalln("no server url provided")
	}

	if channel == nil || *channel == "" {
		log.Fatalln("no channel provided")
	}

	nc, err := nats.Connect(*serverURL)
	if err != nil {
		log.Fatalf("could not connect to nats on %s\n", *serverURL)
	}
	fmt.Printf("Connected to nats server %s using channel '%s'\n", *serverURL, *channel)

	nc.Subscribe(*channel, func(m *nats.Msg) {
		fmt.Printf("Received a message: %s", string(m.Data))
	})
}
