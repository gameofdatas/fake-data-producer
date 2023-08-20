## Console Fake Data Generator - User Guide

The Console Fake Data Generator is a command-line tool that generates fake data based on a provided configuration and displays it on the console.

### Installation

To use the Console Fake Data Generator, you'll need the executable binary file. Here's how to get started:

1. Download or clone the source code from the repository.

2. Navigate to the project root directory using a terminal.

3. Build the binary:
   ```bash
   go build -o console-fake-data-generator cmd/main.go
   ```

### Usage

Once you have the `console-fake-data-generator` binary, you can use it to generate and display fake data on the console. Here's how to use the tool:

```bash
./console-fake-data-generator console [flags]
```

#### Flags

- `--config-dir`, `-c`: Directory containing the configuration file. (Default: current directory)
- `--file`, `-f`: Name of the configuration file. (Default: config.json)
- `--nr-messages`, `-m`: Number of messages to generate (-1 for infinite). (Required)
- `--help`, `-h`: Display usage information.

#### Examples

Generate and display 100 messages of fake data on the console:

```bash
./console-fake-data-generator console --config-dir=configs --file=config.json --nr-messages=100
```

Generate and display an unlimited number of messages of fake data on the console:

```bash
./console-fake-data-generator console --config-dir=configs --file=config.json --nr-messages=-1
```

### Configuration File

The configuration file defines the structure of the generated fake data. It specifies data types, ranges, and formatting for each field. Here's an example of a configuration file:

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

The Console Fake Data Generator is a useful tool for quickly generating and displaying fake data on the console. With its straightforward configuration and command-line options, it's suitable for testing, development, and educational purposes.

For more information and updates, please refer to the official documentation and repository. Enjoy exploring and generating data! 🚀

---