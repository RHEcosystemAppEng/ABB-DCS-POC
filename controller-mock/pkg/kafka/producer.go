package kafka

import (
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

var backoffSchedule = []time.Duration{
	1 * time.Second,
	3 * time.Second,
	5 * time.Second,
}

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
		// fmt.Println(bufferKafkaMsgJson)

		// send kafka message with retries according to backoff schedule
		sendKafkaMsgWithRetries(metric.Name, bufferKafkaMsgJson)
	}
}

func sendKafkaMsgWithRetries(metricName string, bufferKafkaMsgJson *bytes.Buffer) {

	var (
		err  error
		resp *http.Response
	)

	// for each retry time interval in backoff schedule, try to send kafka message
	for i := 0; i <= len(backoffSchedule); i++ {

		// send kafka message
		resp, err = sendKafkaMsg(metricName, bufferKafkaMsgJson)

		// if post successful, break loop
		if err == nil && resp.StatusCode < 500 {
			break
		}

		// if error, log error message
		log.Print("Posting kafka message data over HTTP failed")
		if err != nil {
			log.Printf("Error: %v", err)
		}
		if resp != nil {
			log.Printf("StatusCode: %v", resp.Status)
		}

		// if retry time interval scheduled, sleep
		if i < len(backoffSchedule) {

			log.Printf("Retrying in %v\n", backoffSchedule[i])
			time.Sleep(backoffSchedule[i])

		}
	}

	// if all retries failed, panic
	if err != nil || resp.StatusCode >= 500 {
		log.Fatal("All retries failed, posting kafka message data over HTTP fatal")
	}

	// print response
	log.Println(resp.Status)
}

func sendKafkaMsg(metricName string, bufferKafkaMsgJson *bytes.Buffer) (*http.Response, error) {

	// send kafka message over kafka bridge using http protocol
	resp, err := http.Post(fmt.Sprintf(HTTP_KAFKA_URL_TO_TOPIC, os.Getenv(HTTP_KAFKA_URL_ENV_VAR), metricName), HTTP_KAFKA_CONTENT_TYPE, bufferKafkaMsgJson)

	// if error, return error
	if err != nil {
		return nil, err
	}

	// close body at end of transaction
	defer resp.Body.Close()

	// return response
	return resp, nil
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
		log.Fatalf("Marshaling kafka message data to JSON failed: %v", err)
	}

	// add json wrapper to kafka message
	return fmt.Sprintf(HTTP_KAFKA_MSG_WRAPPER, uuid.New(), kafkaMsgJson)
}
