package pubsub

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
)

type Node struct {
	Id         string   `json:"id"`
	Dependency []string `json:"dependency"`
	Timestamp  int64    `json:"timestamp"`
}

type PubSub interface {
	Subscribe(subject string, callback func(data []byte))
	Publish(subject string, data []byte)
	Close()
}

type NatsEventPublisher struct {
	PubSub PubSub
}

func (p *NatsEventPublisher) PublishEvent(id string, dependencies []string, timestamp int64) {
	node := Node{Id: id, Dependency: dependencies, Timestamp: timestamp}
	subject := node.Id

	log.Println(node)
	data, _ := json.Marshal(node)

	p.PubSub.Publish(subject, data)
}

type NatsPubSub struct {
	nc *nats.Conn
}

func NewNatsPubSub() (*NatsPubSub, error) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return nil, err
	}

	return &NatsPubSub{nc: nc}, nil
}

func (np *NatsPubSub) Subscribe(subject string, callback func(data []byte)) {
	np.nc.Subscribe(subject, func(msg *nats.Msg) {
		callback(msg.Data)
	})
}

func (np *NatsPubSub) Publish(subject string, data []byte) {
	np.nc.Publish(subject, data)
}

func (np *NatsPubSub) Close() {
	np.nc.Close()
}
