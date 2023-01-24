package kafka

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/RHEcosystemAppEng/abb-dcs-poc/controller-mock/pkg/workflow"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	KAFKA_CONFIG_FILEPATH = "./pkg/kafka/controller.properties"
)

type KafkaMessage struct {
	WorkflowName string    `json:"workflow_name"`
	Timestamp    time.Time `json:"timestamp"`
	Metric       Metric    `json:"metric"`
}

type Metric struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

func KafkaProducer(wf *workflow.Workflow) {

	// get kafka config
	conf := ReadConfig(KAFKA_CONFIG_FILEPATH)

	// init kafka producer in compliance with the kafka config
	producer, err := kafka.NewProducer(&conf)
	if err != nil {
		log.Fatalf("Failed to create producer: %s", err)
	}

	// Go-routine to handle message delivery reports and
	// possibly other event types (errors, stats, etc)
	go func() {
		for event := range producer.Events() {
			switch ev := event.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Produced event to topic %s: key = %-10s value = %s\n", *ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
				}
			}
		}
	}()

	for _, metric := range wf.Metrics {

		// init kafka message
		kafkaMsg := NewKafkaMessage(wf, metric)

		// marchal kafka message struct to json
		kafkaMsgJson, err := json.Marshal(kafkaMsg)
		if err != nil {
			log.Fatalf("Marshaling kafka message data to JSON failed: %s", err)
		}

		// define kafka message topic
		topic := metric.Name

		producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Key:            []byte("msg"),
			Value:          []byte(kafkaMsgJson),
		}, nil)
	}

	producer.Flush(15 * 1000)
	producer.Close()
}

func NewKafkaMessage(wf *workflow.Workflow, metric *workflow.Metric) *KafkaMessage {

	km := KafkaMessage{
		WorkflowName: wf.WorkflowName,
		Timestamp:    wf.Timestamp,
		Metric: Metric{
			Name:  metric.Name,
			Value: metric.Value,
		},
	}

	return &km
}
