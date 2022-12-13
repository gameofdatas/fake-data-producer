package kafka

import (
	"fmt"
	"strconv"
	"time"

	"fake-data-producer/pkg/models"
	"fake-data-producer/pkg/queue"
	"fake-data-producer/version"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{
	Use:   "kafka",
	Short: "generate and produce fake data to kafka",
	Long:  `generate and produce fake data to kafka`,
	RunE:  cmd,
}

var bootstrapServer string
var saslMechanism string
var certFolder string
var username string
var password string
var nrMessage string
var maxWaitingTime string
var securityProtocol string
var topic string
var subject string

func init() {
	RunCmd.Flags().StringVarP(&securityProtocol, "security-protocol", "s", "PLAINTEXT", "Security protocol for Kafka (PLAINTEXT, SSL, SASL_SSL)")
	RunCmd.Flags().StringVarP(&saslMechanism, "sasl-mechanism", "m", "", "SASL mechanism for Kafka (PLAIN, GSSAPI, OAUTHBEARER, SCRAM-SHA-256, SCRAM-SHA-512)")
	RunCmd.Flags().StringVarP(&certFolder, "cert-folder", "c", "", "Path to folder containing required Kafka certificates. Required --security-protocol equal SSL or SASL_SSL")
	RunCmd.Flags().StringVarP(&username, "username", "u", "", "Username required if security-protocol is SASL_SSL")
	RunCmd.Flags().StringVarP(&password, "password", "p", "", "Password required if security-protocol is SASL_SSL")
	RunCmd.Flags().StringVarP(&bootstrapServer, "bootstrap-server", "b", "", "Kafka bootstrap server")
	RunCmd.Flags().StringVarP(&topic, "topic", "t", "", "topic name")
	RunCmd.Flags().StringVarP(&nrMessage, "nr-messages", "r", "", "Number of messages to produce (0 for unlimited)")
	RunCmd.Flags().StringVarP(&maxWaitingTime, "max-waiting-time", "w", "", "Max waiting time between messages (0 for none) in seconds")
	RunCmd.Flags().StringVarP(&subject, "model", "l", "", "data models to produce")
	RunCmd.MarkFlagRequired("security-protocol")
	RunCmd.MarkFlagRequired("bootstrap-server")
	RunCmd.MarkFlagRequired("topic")
	RunCmd.MarkFlagRequired("nr-messages")
	RunCmd.MarkFlagRequired("max-waiting-time")
}

func cmd(_ *cobra.Command, _ []string) error {
	log.Info().Msgf("starting the fake data generator to kafka with \n Version: %s %s %s \n on %s => %s",
		version.Version,
		version.BuildDate,
		version.GitCommit,
		version.OsArch,
		version.GoVersion)

	kafkaConfig := &queue.KafkaConfig{
		BootstrapServer:  bootstrapServer,
		SaslMechanism:    saslMechanism,
		CertFolder:       certFolder,
		Username:         username,
		Password:         password,
		NumOfMessage:     nrMessage,
		MaxWaitTime:      maxWaitingTime,
		SecurityProtocol: securityProtocol,
		Topic:            topic,
	}
	producer, err := queue.NewQueue("kafka", kafkaConfig)
	if err != nil {
		return fmt.Errorf("cannot create producer: %w", err)
	}

	defer producer.Close()

	messageCount, err := strconv.Atoi(nrMessage)
	if err != nil {
		return err
	}
	maxWaitTime, err := strconv.Atoi(maxWaitingTime)
	if err != nil {
		return err
	}

	goInfinite := false
	wait := false
	if messageCount <= 0 {
		goInfinite = true
	}
	if maxWaitTime > 0 {
		wait = true
	}

	count := 0
	for {
		key, value, err := models.CreateData(subject)
		if err != nil {
			return fmt.Errorf("error creating data for subject %s: %w", subject, err)
		}
		if wait {
			time.Sleep(time.Duration(maxWaitTime) * time.Second)
		}
		err = producer.Produce(key, value)
		if err != nil {
			return fmt.Errorf("producing error %w", err)
		}
		count++
		if !goInfinite && count == messageCount {
			break
		}
	}
	return nil
}
