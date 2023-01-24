package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/goombaio/namegenerator"
)

const (
	METRICS_CONFIG_FILEPATH = "./pkg/controller/initial_metrics_config.json"
)

type Controller struct {
	ControllerName string    `json:"controller_name"`
	Timestamp      time.Time `json:"timestamp"`
	Metrics        []*Metric `json:"metrics"`
}

func InitController() *Controller {

	// generate name for controller
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)
	controllerName := nameGenerator.Generate()

	// init controller struct
	controller := Controller{
		ControllerName: controllerName,
		Timestamp:      time.Now(),
	}

	// read metrics config file
	metricsConfig, err := ioutil.ReadFile(METRICS_CONFIG_FILEPATH)
	if err != nil {
		log.Fatal(err)
	}

	// unmarshal metrics config onto controller struct
	json.Unmarshal([]byte(metricsConfig), &controller)

	return &controller
}

func (wf *Controller) PromoteControllerMetrics() {

	for _, metric := range wf.Metrics {
		// change metric Strategy if necessary
		metric.determineMetricStrategy()
		// advance metric value in relation to metric strategy
		metric.advanceMetricValue()
	}

	// reset timestamp
	wf.Timestamp = time.Now()
}

func (wf *Controller) ReturnControllerData(w http.ResponseWriter, r *http.Request) {

	// decode controller data to json
	json.NewEncoder(w).Encode(wf)
}
