package writer

type KafkaConfig struct {
	BootstrapServer  string
	SaslMechanism    string
	CertFolder       string
	Username         string
	Password         string
	SecurityProtocol string
	Topic            string
}

type CSVConfig struct {
	Path     string
	FileName string
}
