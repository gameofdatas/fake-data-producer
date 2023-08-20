package kafka

import (
	"fmt"
	"path/filepath"
	"strconv"
	"time"

	"fake-data-producer/config"
	"fake-data-producer/pkg/fakedata"
	"fake-data-producer/pkg/writer"
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
var configDir string
var file string

func init() {
	RunCmd.Flags().StringVarP(&securityProtocol, "security-protocol", "s", "PLAINTEXT", "Security protocol for Kafka (PLAINTEXT, SSL, SASL_SSL)")
	RunCmd.Flags().StringVarP(&saslMechanism, "sasl-mechanism", "m", "", "SASL mechanism for Kafka (PLAIN, GSSAPI, OAUTHBEARER, SCRAM-SHA-256, SCRAM-SHA-512)")
	RunCmd.Flags().StringVarP(&certFolder, "cert-folder", "c", "", "Path to folder containing required Kafka certificates. Required if --security-protocol is SSL or SASL_SSL")
	RunCmd.Flags().StringVarP(&username, "username", "u", "", "Username required if --security-protocol is SASL_SSL")
	RunCmd.Flags().StringVarP(&password, "password", "p", "", "Password required if --security-protocol is SASL_SSL")
	RunCmd.Flags().StringVarP(&bootstrapServer, "bootstrap-server", "b", "", "Kafka bootstrap server")
	RunCmd.Flags().StringVarP(&topic, "topic", "t", "", "Name of the topic to produce messages to")
	RunCmd.Flags().StringVarP(&nrMessage, "nr-messages", "n", "", "Number of messages to produce (0 for unlimited)")
	RunCmd.Flags().StringVarP(&maxWaitingTime, "max-waiting-time", "w", "", "Maximum waiting time between messages (0 for no waiting) in seconds")
	RunCmd.Flags().StringVarP(&configDir, "config-dir", "d", "/foo/dir1", "Directory path of the config file")
	RunCmd.Flags().StringVarP(&file, "file", "f", "config.json", "Name of the config file")
	RunCmd.MarkFlagRequired("security-protocol")
	RunCmd.MarkFlagRequired("bootstrap-server")
	RunCmd.MarkFlagRequired("topic")
	RunCmd.MarkFlagRequired("nr-messages")
	RunCmd.MarkFlagRequired("max-waiting-time")
	RunCmd.MarkFlagRequired("config-dir")
	RunCmd.MarkFlagRequired("file")
}

func cmd(_ *cobra.Command, _ []string) error {
	log.Info().Msgf("starting the fake data generator to kafka with \n Version: %s %s %s \n on %s => %s",
		version.Version,
		version.BuildDate,
		version.GitCommit,
		version.OsArch,
		version.GoVersion)

	configData, err := config.ReadConfigFile(filepath.Join(configDir, file))
	if err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}

	kafkaConfig := &writer.KafkaConfig{
		BootstrapServer:  bootstrapServer,
		SaslMechanism:    saslMechanism,
		CertFolder:       certFolder,
		Username:         username,
		Password:         password,
		SecurityProtocol: securityProtocol,
		Topic:            topic,
	}
	producer, err := writer.NewWriter("kafka", kafkaConfig)
	if err != nil {
		return fmt.Errorf("cannot create producer: %w", err)
	}

	defer producer.Close()

	messageCount, err := strconv.Atoi(nrMessage)
	if err != nil {
		return err
	}
	dataStruct, err := fakedata.CreateStructFromJSON(configData)
	if err != nil {
		return err
	}
	count := 0
	waitTime := time.Duration(0)
	if maxWaitingTime != "" {
		parsedWaitTime, err := strconv.Atoi(maxWaitingTime)
		if err != nil {
			return fmt.Errorf("invalid max wait time: %w", err)
		}
		if parsedWaitTime > 0 {
			waitTime = time.Duration(parsedWaitTime) * time.Second
		}
	}

	goInfinite := messageCount <= 0
	if waitTime > 0 {
		ticker := time.NewTicker(waitTime)
		defer ticker.Stop()

		for range ticker.C {
			key, value, err := fakedata.CreateData(dataStruct, configData)
			if err != nil {
				return fmt.Errorf("error creating data: %w", err)
			}

			err = producer.Produce(key, value)
			if err != nil {
				return fmt.Errorf("producing error: %w", err)
			}
			count++
			if !goInfinite && count == messageCount {
				break
			}
		}
	} else {
		for {
			key, value, err := fakedata.CreateData(dataStruct, configData)
			if err != nil {
				return fmt.Errorf("error creating data: %w", err)
			}
			err = producer.Produce(key, value)
			if err != nil {
				return fmt.Errorf("producing error: %w", err)
			}

			count++
			if !goInfinite && count == messageCount {
				break
			}
		}
	}
	producer.Flush(15 * 1000)
	return nil
}
