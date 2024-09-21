# Kafka Notification
A notification exchange system using event streaming with Kafka

## How to Run
1. Ensure Docker is installed and running at local machine;
2. Download Kafka CLI binaries from [Apache Kafka official site](https://kafka.apache.org/downloads);
3. Create a kafka topic using snippet in section [Create a kafka topic](#create-a-kafka-topic)
4. Run `go run producer_client.go` and `go run consumer_client.go`

## Create a kafka topic
```
kafka-topics.bat --bootstrap-server localhost:9092 --topic <topic-name> --create --partitions 3 --replication-factor 1
```

## List all topics:
```
kafka-topics.bat --bootstrap-server localhost:9092 --list
```

## Evoke kafka consumer CLI client
```
kafka-console-consumer.bat --bootstrap-server localhost:9092 --topic <topic-name> --from-beginning
```