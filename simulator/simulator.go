package simulator

import (
	"log"
	"time"

	"github.com/context-agent/internal/event"
)

func SimulateRealWorld(publisher event.EventPublisher) {

	log.Println("TASK : New incoming event")
	publisher.PublishEvent("pharma", []string{}, time.Now().Unix())
	time.Sleep(1 * time.Second)

	// Add more events using the publisher
	log.Println("TASK : New incoming event")
	publisher.PublishEvent("opm", []string{"defect", "sample"}, time.Now().Unix())
	time.Sleep(1 * time.Second)

	log.Println("TASK : New incoming event")
	publisher.PublishEvent("users", []string{"someuser"}, time.Now().Unix())
	time.Sleep(1 * time.Second)

	log.Println("TASK : New incoming event")
	publisher.PublishEvent("infra", []string{"disk", "network"}, time.Now().Unix())
	time.Sleep(1 * time.Second)

	// Add nodes with dependencies
	log.Println("TASK : New incoming task")
	publisher.PublishEvent("network", []string{"dns", "dnspoison"}, time.Now().Unix())
	time.Sleep(1 * time.Second)
}
