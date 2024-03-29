# Fake Data Producer

The **Fake Data Producer** is a versatile tool designed to generate synthetic data for various purposes, such as
testing, development, and data analysis. This command-line application empowers you to effortlessly create realistic
datasets with customizable data types, ranges, and formats. Whether you're populating databases, validating software, or
exploring data pipelines, the Fake Data Producer streamlines the process of generating diverse data that mimics
real-world scenarios. With support for a wide range of data types and customization options, you can tailor your
generated data to match your specific use case.

## Key Features

- **Configurable Data Generation:** Define your dataset's structure using an intuitive JSON configuration. Specify data
  types, ranges, and additional parameters to precisely control the generated content. **[Configuration Dictionary](conf/readme.md)**

```shell
sample config

{
  "Guid": {
    "type": "string",
    "isId": true
  },
  "Platform": {
    "type": "string",
    "options": [
      "web",
      "mobile",
      "ai"
    ]
  },
  "PhoneNumber": "string",
  "Weight": {
    "type": "float",
    "min": 30,
    "max": 120,
    "precision": 2
  },
  "Address": "string",
  "Latitude": "float",
  "Longitude": "float",
  "UnitPrice": {
    "type": "float",
    "min": 2.23,
    "max": 5.21,
    "precision": 4
  },
  "Age": {
    "type": "int",
    "min": 20,
    "max": 60
  },
  "Salary": {
    "type": "int",
    "min": 50000,
    "max": 100000
  },
  "Timestamp": {
    "type": "timestamp",
    "format": "yyyy-MM-dd'T'HH:mm:ssZZZZ",
    "zone": "UTC"
  },
  "Ts": {
    "type": "epoch",
    "unit": "millis",
    "zone": "UTC"
  }
```
- **Rich Data Types:** The Fake Data Producer supports various data types, including strings, integers, floating-point numbers, booleans, timestamps, and more. Customizable options like "options" for strings and "min/max" for numeric values offer fine-grained control.

- **Realistic Field Generators:** Leveraging the power of the [gofakeit](https://github.com/brianvoe/gofakeit) library, the Fake Data Producer employs sophisticated field generators to create authentic and diverse data for different columns.

- **Dynamic Output Formats:** Generate data in multiple formats such as CSV, Kafka messages, or even directly to the console. Seamlessly integrate the generated data into your preferred workflows and systems.

  - [kafka producer](cmd/kafka/readme.md)
  - [csv producer](cmd/csv/readme.md)
  - [console_producer](cmd/console/readme.md)

- **Easy Integration:** Integrate the Fake Data Producer into your development and testing processes using a simple command-line interface. Effortlessly generate datasets with a few commands, allowing you to focus on more critical aspects of your project.

## Running via Docker

The Fake Data Producer can be conveniently run using Docker, allowing you to quickly generate synthetic data without worrying about installing dependencies or configuring your environment. Follow these steps to run the Fake Data Producer using Docker:

### 1. Build the Docker Image

Before running the Fake Data Producer, you need to build the Docker image. Open a terminal and navigate to the root directory of the project. Then, run the following command:

```bash
docker build -t fake-data-producer .

or 

make build (make should be installed)
```

This will build the Docker image with the tag `fake-data-producer`.

### 2. Prepare Your Configuration File

Create or prepare a configuration JSON file (e.g., `data.json`) that defines the structure and data types for your generated dataset. You can customize this configuration file based on your requirements.

### 3. Run the Fake Data Producer

To run the Fake Data Producer using Docker, use the following command:

```bash
docker run -it -v /path/to/your/config/dir:/config fake-data-producer console --config-dir /config --file data.json --nr-messages 10

or

make run-console CONFIG_DIR=${PWD}/conf CONFIG_FILE=data.json NUM_MESSAGES=10
```

Replace `/path/to/your/config/dir` with the actual path to the directory containing your configuration JSON file. You can also adjust the `data.json` and `--nr-messages` values as needed.

### 4. Generated Data

The Fake Data Producer will generate synthetic data based on the configuration you provided. The generated data will be displayed in the console output.

### Cleaning Up

After you're done generating data, you can clean up the Docker image by running:

```bash
docker rmi fake-data-producer

or 

make clean
```

This will remove the Docker image used for running the Fake Data Producer.

By following these steps, you can easily generate synthetic data using the Fake Data Producer within a Docker container, making it a convenient and isolated environment for data generation.

**Note:** Make sure you have Docker installed on your machine before running the above commands.

---
