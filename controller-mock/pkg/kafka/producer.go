package kafka

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/RHEcosystemAppEng/abb-dcs-poc/controller-mock/pkg/controller"
	"github.com/google/uuid"
)

const (
	HTTP_KAFKA_URL_TO_TOPIC = "%s/topics/%s"
	HTTP_KAFKA_URL_ENV_VAR  = "HTTP_KAFKA_URL"
	HTTP_KAFKA_CONTENT_TYPE = "application/vnd.kafka.json.v2+json"
	HTTP_KAFKA_MSG_WRAPPER  = "{\"records\":[{\"key\": \"%s\",\"value\": %s}]}"
)

type KafkaMessage struct {
	ControllerId   string    `json:"controller_id"`
	ControllerName string    `json:"controller_name"`
	Timestamp      time.Time `json:"timestamp"`
	Metric         Metric    `json:"metric"`
}

type Metric struct {
	Name      string    `json:"name"`
	Value     float64   `json:"value"`
	MinValue  float64   `json:"min_value"`
	MaxValue  float64   `json:"max_value"`
	Unit      string    `json:"unit"`
	Timestamp time.Time `json:"timestamp"`
}

func HTTPKafkaProducer(c *controller.Controller) {

	for _, metric := range c.Metrics {

		// init kafka message struct
		kafkaMsg := newKafkaMessage(c, metric)

		// marshal kafka message struct to json and add json wrapper
		kafkaMsgJson := kafkaMsg.buildBody()

		// buffer kafka message
		bufferKafkaMsgJson := bytes.NewBuffer([]byte(kafkaMsgJson))
		//fmt.Println(bufferKafkaMsgJson)
		// send kafka message over kafka using http protocol, wait for response
		resp, err := http.Post(fmt.Sprintf(HTTP_KAFKA_URL_TO_TOPIC, os.Getenv(HTTP_KAFKA_URL_ENV_VAR), metric.Name), HTTP_KAFKA_CONTENT_TYPE, bufferKafkaMsgJson)
		if err != nil {
			log.Fatalf("Posting kafka message data over HTTP failed: %s", err)
		}
		defer resp.Body.Close()

		// print response
		//log.Println(resp.Status)
		bodyAnswer := bufio.NewScanner(resp.Body)
		for bodyAnswer.Scan() {
			log.Println(bodyAnswer.Text())
		}
	}
}

func newKafkaMessage(c *controller.Controller, metric *controller.Metric) *KafkaMessage {

	km := KafkaMessage{
		ControllerId:   c.ControllerId,
		ControllerName: c.ControllerName,
		Timestamp:      c.Timestamp,
		Metric: Metric{
			Name:      metric.Name,
			Value:     metric.Value,
			MinValue:  metric.RangeMin,
			MaxValue:  metric.RangeMax,
			Unit:      metric.Unit,
			Timestamp: metric.Timestamp,
		},
	}

	return &km
}

func (msg *KafkaMessage) buildBody() string {

	// marshal kafka message struct to json
	kafkaMsgJson, err := json.Marshal(msg)
	if err != nil {
		log.Fatalf("Marshaling kafka message data to JSON failed: %s", err)
	}

	// add json wrapper to kafka message
	return fmt.Sprintf(HTTP_KAFKA_MSG_WRAPPER, uuid.New(), kafkaMsgJson)
}
