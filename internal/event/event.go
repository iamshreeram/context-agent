package event

type EventPublisher interface {
	PublishEvent(id string, dependencies []string, timestamp int64)
}
