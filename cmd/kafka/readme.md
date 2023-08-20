## Kafka Fake Data Generator - User Guide

The Kafka Fake Data Generator is a command-line tool designed to generate fake data based on a provided configuration and publish it to a Kafka topic.

### Installation

To utilize the Kafka Fake Data Generator, you'll need the executable binary file. Here's how to get started:

1. Download or clone the source code from the repository.

2. Navigate to the project root directory using a terminal.

3. Build the binary:
   ```bash
   go build -o kafka-fake-data-generator cmd/main.go
   ```

### Usage

After obtaining the `kafka-fake-data-generator` binary, you can use it to generate and produce fake data to a Kafka topic. Here's how to use the tool:

```bash
./kafka-fake-data-generator kafka [flags]
```

#### Flags

- `--security-protocol`, `-s`: Kafka security protocol (PLAINTEXT, SSL, SASL_SSL). (Required)
- `--sasl-mechanism`, `-m`: SASL mechanism for authentication (PLAIN, GSSAPI, OAUTHBEARER, SCRAM-SHA-256, SCRAM-SHA-512).
- `--cert-folder`, `-c`: Path to the folder containing Kafka certificates. Required if `--security-protocol` is SSL or SASL_SSL.
- `--username`, `-u`: Username for SASL_SSL authentication.
- `--password`, `-p`: Password for SASL_SSL authentication.
- `--bootstrap-server`, `-b`: Kafka bootstrap server. (Required)
- `--topic`, `-t`: Kafka topic to produce messages to. (Required)
- `--nr-messages`, `-n`: Number of messages to produce (0 for unlimited). (Required)
- `--max-waiting-time`, `-w`: Maximum waiting time between messages (0 for no waiting) in seconds. (Required)
- `--config-dir`, `-d`: Directory path of the config file. (Default: /foo/dir1)
- `--file`, `-f`: Name of the config file. (Default: config.json)
- `--help`, `-h`: Display usage information.

#### Examples

Generate and produce 100 messages to the "my-topic" Kafka topic with a maximum waiting time of 5 seconds:

```bash
./kafka-fake-data-generator kafka --security-protocol=PLAINTEXT --bootstrap-server=localhost:9092 --topic=my-topic --nr-messages=100 --max-waiting-time=5
```

Generate an unlimited number of messages to the "my-topic" Kafka topic with no waiting time:

```bash
./kafka-fake-data-generator kafka --security-protocol=PLAINTEXT --bootstrap-server=localhost:9092 --topic=my-topic --nr-messages=0 --max-waiting-time=0
```

### Configuration File

The configuration file specifies the structure of the generated fake data. It defines the data types, ranges, and formatting for each field. Here's an example of a configuration file:

```json
{
  "Guid": { "type": "string", "isId": true },
  "Platform": { "type": "string", "options": ["web", "mobile", "ai"] },
  "PhoneNumber": "string",
  "Weight": { "type": "float", "min": 30, "max": 120, "precision": 2 },
  "Address": "string",
  "Latitude": "float",
  "Longitude": "float",
  "UnitPrice": { "type": "float", "min": 2.23, "max": 5.21, "precision": 4 },
  "Age": { "type": "int", "min": 20, "max": 60 },
  "Salary": { "type": "int", "min": 50000, "max": 100000 },
  "Timestamp": { "type": "timestamp", "format": "yyyy-MM-dd'T'HH:mm:ssZZZZ", "zone": "UTC" },
  "Ts": { "type": "epoch", "unit": "millis", "zone": "UTC" }
}
```

### Summary

The Kafka Fake Data Generator is a powerful tool for generating and publishing fake data to a Kafka topic. With its flexible configuration and command-line options, it's well-suited for testing, development, and data analysis scenarios.

For further information and updates, refer to the official documentation and repository. Happy generating! 🚀

---