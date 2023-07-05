package queue

import (
	"fmt"
)

// Queue represents a message queue.
type Queue interface {
	Produce([]byte, []byte) error
	Close()
	Flush(i int) int
}

// NewQueue creates a new instance of the message queue based on the provided queue type and configuration.
// It returns a Queue interface and an error if any.
func NewQueue(queue string, cfg interface{}) (Queue, error) {
	switch queue {
	case "kafka":
		kafkaConfig := cfg.(*KafkaConfig)
		return newKafkaWriter(kafkaConfig)
	default:
		return nil, fmt.Errorf("queue type not supported: %s", queue)
	}
}
