package main

import (
	"log"
	"time"

	"github.com/iamshreeram/context-agent/internal/pubsub"
	"github.com/iamshreeram/context-agent/simulator"
)

func main() {
	natsPubSub, err := pubsub.NewNatsPubSub()
	if err != nil {
		panic(err)
	}

	eventPublisher := &pubsub.NatsEventPublisher{natsPubSub: natsPubSub}

	log.Println("INIT : Initializing context-agent app..")
	time.Sleep(1 * time.Second)

	simulator.SimulateRealWorld(eventPublisher)
}
