## Fake Data Generator to CSV - User Guide

The Fake Data Generator to CSV is a command-line tool that generates fake data based on a provided configuration and writes it to a CSV file.

### Installation

To use the Fake Data Generator to CSV, you need to have the executable binary file. You can either compile it from the source code or obtain a pre-built binary from a distribution. Here's how to get started:

1. Download or clone the source code from the repository.

2. Navigate to the project root directory using a terminal.

3. Build the binary:
   ```bash
   go build -o fake-data-generator cmd/main.go
   ```

### Usage

Once you have the `fake-data-generator` binary, you can use it to generate and produce fake data in CSV format. Here's how to use the tool:

```bash
./fake-data-generator csv [flags]
```

#### Flags

- `--config-dir`, `-c`: Directory containing the configuration file. (Default: current directory)
- `--file`, `-f`: Name of the configuration file. (Default: config.json)
- `--output-dir`, `-o`: Directory for the output CSV file. (Required)
- `--output-filename`, `-n`: Name of the output CSV file. (Required)
- `--nr-messages`, `-m`: Number of messages to generate. Use -1 for an infinite number. (Required)
- `--help`, `-h`: Display usage information.

#### Examples

Generate 100 fake data records and save them to `output.csv` in the `data` directory:

```bash
./fake-data-generator csv --config-dir=configs --file=config.json --output-dir=data --output-filename=output.csv --nr-messages=100
```

Generate an infinite number of fake data records and save them to `output.csv` in the current directory:

```bash
./fake-data-generator csv --config-dir=configs --file=config.json --output-dir=. --output-filename=output.csv --nr-messages=-1
```

### Configuration File

The configuration file defines the structure of the generated fake data. It specifies the data types, ranges, and formatting for each field. An example configuration file might look like this:

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

The Fake Data Generator to CSV is a powerful tool to quickly create and store large amounts of fake data in CSV format. With flexible configuration and command-line options, it adapts to your needs for testing, development, and data analysis.

For more information and updates, please refer to the official documentation and repository. Happy data generation! 🚀

---