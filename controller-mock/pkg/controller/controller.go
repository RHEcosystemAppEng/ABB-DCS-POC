package controller

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/goombaio/namegenerator"
)

const (
	METRICS_CONFIG_FILEPATH = "./pkg/controller/initial_metrics_config.json"
	CONTROLLER_ID_TEMPLATE  = "controller-%s"
)

type Controller struct {
	ControllerId   string    `json:"controller_id"`
	ControllerName string    `json:"controller_name"`
	Timestamp      time.Time `json:"timestamp"`
	Metrics        []*Metric `json:"metrics"`
}

func InitController() *Controller {

	// init controller struct
	controller := Controller{
		ControllerId:   buildControllerId(),
		ControllerName: generateControllerName(),
		Timestamp:      time.Now(),
	}

	// read metrics config file
	metricsConfig, err := ioutil.ReadFile(METRICS_CONFIG_FILEPATH)
	if err != nil {
		log.Fatalf("Read data from metrics config file failed: %s", err)
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

func buildControllerId() string {

	// get hostname
	hostName, err := os.Hostname()
	if err != nil {
		log.Fatalf("Failed to get hostname: %s", err)
	}

	// hash hostname
	hashedHostName := hashString(hostName)

	// build controller id from controller id template and hashed hostname 4 letter suffix
	return fmt.Sprintf(CONTROLLER_ID_TEMPLATE, hashedHostName[len(hashedHostName)-4:])
}

func hashString(str string) string {
	hashStr := sha1.New()
	hashStr.Write([]byte(str))
	return hex.EncodeToString(hashStr.Sum(nil))
}

func generateControllerName() string {

	// generate name for controller using rand name generator
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)
	return nameGenerator.Generate()
}

func (wf *Controller) ReturnControllerData(w http.ResponseWriter, r *http.Request) {

	// decode controller data to json
	json.NewEncoder(w).Encode(wf)
}
