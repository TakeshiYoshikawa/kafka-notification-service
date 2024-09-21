# Kafka Notification

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