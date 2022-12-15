## fake-producer

### Introduction

An application which allows you to produce datasets and flush it to different queueing service like kafka, rabbitmq etc. 

#### dependencies
- [confluentic-kafka-go](https://github.com/confluentinc/confluent-kafka-go): kafka client to flush data to Kafka system
- [gofakeit](https://github.com/brianvoe/gofakeit): Random data generator written in go

### Prerequisite
- docker
- go
- kafka

### Usage

The application can be run with following ways:

with `PLAINTEXT` as security-protocol:

```shell
go run main.go kafka \
    --security-protocol PLAINTEXT \
    --bootstrap-server localhost:9092 \
    --topic sampleTopic \
    --nr-messages 0 \
    --max-waiting-time 0 \
    --model address
```

with `SSL` as security-protocol:

```shell
go run main.go kafka \
    --security-protocol SSL \
    --bootstrap-server localhost:9092 \
    --topic sampleTopic \
    --nr-messages 0 \
    --max-waiting-time 0 \
    --model address
```

where

- `security-protocol` : Security protocol for Kafka (PLAINTEXT, SSL, SASL_SSL)
- `bootstrap-server` : Kafka bootstrap server
- `topic` : Topic name
- `nr-messages` : Number of messages to produce (0 for unlimited)
- `max-waiting-time` : Max waiting time between messages (0 for none) in seconds
- `model` : data models to produce

#### Run application via docker

```shell
docker build -t fake-data-producer .

docker run fake-data-producer kafka \ 
    --security-protocol PLAINTEXT \
    --bootstrap-server localhost:9092 \
    --topic sampleTopic \
    --nr-messages 0 \
    --max-waiting-time 0 \
    --model address
```

via **[docker-compose](docker-compose.yml)**

It runs three services:
- zookeeper
- kafka
- fake-data-producer

```shell
docker-compose -f docker-compose.yml up -d
```

![docker](images/docker-run-successfully.png)

Let's check the data:

```shell
docker exec -it kafka /bin/sh
cd opt/<kafka version>/bin
kafka-topics.sh --list --zookeeper zookeeper:218
kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic sampleTopic --from-beginning
```

![run](images/data.png)

### Data Models

#### address

```text
{"id":"1098","street":"9806 Station furt","city":"Chandler","state":"California","zip":"68689","country":"Aruba","latitude":85.341363,"longitude":-13.611091}
{"id":"1010","street":"131 New Lake ton","city":"Oklahoma ","state":"Connecticut","zip":"83629","country":"Germany","latitude":71.046268,"longitude":179.745551}
{"id":"1073","street":"4883 North Vista town","city":"Omaha","state":"Arkansas","zip":"39504","country":"Turkmenistan","latitude":37.978425,"longitude":81.937671}
{"id":"1046","street":"78552 East Keys side","city":"Henderson","state":"Georgia","zip":"44646","country":"Jordan","latitude":35.166294,"longitude":83.33641}
{"id":"1099","street":"634 Square town","city":"Newark","state":"South Carolina","zip":"19963","country":"Germany","latitude":-78.26594,"longitude":-106.274358}
{"id":"1039","street":"329 Lake Corners shire","city":"El Paso","state":"Ohio","zip":"89528","country":"Réunion","latitude":-28.468999,"longitude":142.784921}
{"id":"1089","street":"90517 South Fork borough","city":"Laredo","state":"Illinois","zip":"95052","country":"Suriname","latitude":-43.701825,"longitude":108.197004}
{"id":"1080","street":"2121 New Pass side","city":"Atlanta","state":"South Carolina","zip":"56516","country":"Cocos (Keeling) Islands","latitude":-33.157278,"longitude":49.749062}
{"id":"1042","street":"45475 Drive haven","city":"Oklahoma ","state":"New Jersey","zip":"27126","country":"Pakistan","latitude":87.893807,"longitude":28.88553}
{"id":"1057","street":"22438 View chester","city":"St. Paul","state":"Kansas","zip":"68935","country":"Dominican Republic","latitude":-55.825918,"longitude":35.243267}
```