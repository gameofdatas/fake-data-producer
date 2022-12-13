package queue

type KafkaConfig struct {
	BootstrapServer  string
	SaslMechanism    string
	CertFolder       string
	Username         string
	Password         string
	NumOfMessage     string
	MaxWaitTime      string
	SecurityProtocol string
	Topic            string
}
