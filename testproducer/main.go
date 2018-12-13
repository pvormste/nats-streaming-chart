package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/nats-io/go-nats"
)

func main() {
	publish()
}

func publish() {
	serverURL := flag.String("server", "http://localhost:4222", "url to the nats streaming server")
	channel := flag.String("channel", "demo", "channel on which the producer will publish")
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

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter a message: ")
		msg, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("error occured: %s\n", err)
			continue
		}

		nc.Publish(*channel, []byte(msg))
	}
}
