package writer

import (
	"fmt"
)

// WriterInterface represents an output interface.
type WriterInterface interface {
	Produce([]byte, []byte) error
	Close()
	Flush(i int) int
}

// NewWriter creates a new instance of the message queue based on the provided queue type and configuration.
// It returns a WriterInterface and an error if any.
func NewWriter(output string, cfg interface{}) (WriterInterface, error) {
	switch output {
	case "kafka":
		kafkaConfig := cfg.(*KafkaConfig)
		return newKafkaWriter(kafkaConfig)
	case "csv":
		csvConfig := cfg.(*CSVConfig)
		return NewCSVWriter(csvConfig)
	case "console":
		return NewConsoleWriter()
	default:
		return nil, fmt.Errorf("output type not supported: %s", output)
	}
}
