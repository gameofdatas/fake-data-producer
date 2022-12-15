package queue

import (
	"fmt"
)

type Queue interface {
	Produce([]byte, []byte) error
	Close()
	Flush(i int) int
}

func NewQueue(queue string, cfg interface{}) (Queue, error) {
	switch queue {
	case "kafka":
		kafkaConfig := cfg.(*KafkaConfig)
		return newKafkaWriter(kafkaConfig)
	default:
		return nil, fmt.Errorf("queue type not supported %s", queue)
	}
}
