package queue

import (
	"strings"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rs/zerolog/log"
)

type KafkaQueue struct {
	producer *kafka.Producer
	topic    *string
	cfg      *KafkaConfig
}

func (k *KafkaQueue) Flush(i int) int {
	return k.producer.Flush(i)
}

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

func (k *KafkaQueue) Close() {
	k.producer.Close()
}

func newKafkaWriter(cfg *KafkaConfig) (Queue, error) {
	p, err := kafka.NewProducer(createKafkaConfig(cfg))

	return &KafkaQueue{
		producer: p,
		topic:    &cfg.Topic,
		cfg:      cfg,
	}, err
}

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
