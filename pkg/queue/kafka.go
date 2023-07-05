package queue

import (
	"strings"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rs/zerolog/log"
)

// KafkaQueue represents a Kafka message queue.
type KafkaQueue struct {
	producer *kafka.Producer
	topic    *string
	cfg      *KafkaConfig
}

// Flush flushes the Kafka producer's message queue and waits for messages to be delivered.
// It takes an integer argument 'i' as the timeout value in milliseconds.
// It returns the number of messages flushed from the queue.
func (k *KafkaQueue) Flush(i int) int {
	return k.producer.Flush(i)
}

// Produce publishes a Kafka message with the provided key and value to the configured topic.
// It returns an error if the message fails to be produced.
// If the Kafka local queue is full, it flushes the queue and retries producing the message.
func (k *KafkaQueue) Produce(key, value []byte) error {
	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     k.topic,
			Partition: kafka.PartitionAny,
		},
		Key:   key,
		Value: value,
	}
	err := k.producer.Produce(msg, nil)
	if kafkaError, ok := err.(kafka.Error); ok && kafkaError.Code() == kafka.ErrQueueFull {
		log.Warn().Msg("kafka local queue full error - Going to Flush then retry...")
		flushedMessages := k.producer.Flush(30 * 1000)
		log.Info().Msgf("flushed kafka messages. Outstanding events still un-flushed: %d", flushedMessages)
		return k.Produce(key, value)
	}
	return err
}

// Close closes the Kafka producer and releases any associated resources.
func (k *KafkaQueue) Close() {
	k.producer.Close()
}

// newKafkaWriter creates a new Kafka message queue writer based on the provided KafkaConfig.
// It returns a Queue interface and an error if any.
func newKafkaWriter(cfg *KafkaConfig) (Queue, error) {
	p, err := kafka.NewProducer(createKafkaConfig(cfg))

	return &KafkaQueue{
		producer: p,
		topic:    &cfg.Topic,
		cfg:      cfg,
	}, err
}

// createKafkaConfig creates a Kafka configuration based on the provided KafkaConfig.
// It returns a ConfigMap containing the Kafka configuration properties.
func createKafkaConfig(cfg *KafkaConfig) *kafka.ConfigMap {
	switch strings.ToUpper(cfg.SecurityProtocol) {
	case "SASL_SSL":
		return &kafka.ConfigMap{
			"bootstrap.servers": cfg.BootstrapServer,
			"security.protocol": cfg.SecurityProtocol,
			"sasl.username":     cfg.Username,
			"sasl.password":     cfg.Password,
		}
	}
	return &kafka.ConfigMap{
		"bootstrap.servers":   cfg.BootstrapServer,
		"api.version.request": true,
		"enable.idempotence":  true,
		"security.protocol":   cfg.SecurityProtocol,
	}
}
