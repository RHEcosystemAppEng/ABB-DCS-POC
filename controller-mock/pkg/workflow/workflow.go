package workflow

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/goombaio/namegenerator"
)

const (
	METRICS_CONFIG_FILEPATH = "./pkg/workflow/initial_metrics_config.json"
)

type Workflow struct {
	WorkflowName string    `json:"workflow_name"`
	Timestamp    time.Time `json:"timestamp"`
	Metrics      []*Metric `json:"metrics"`
}

func InitWorkflow() *Workflow {

	// generate name for workflow
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)
	workflowName := nameGenerator.Generate()

	// init workflow struct
	workflow := Workflow{
		WorkflowName: workflowName,
		Timestamp:    time.Now(),
	}

	// read metrics config file
	metricsConfig, err := ioutil.ReadFile(METRICS_CONFIG_FILEPATH)
	if err != nil {
		log.Fatal(err)
	}

	// unmarshal metrics config onto workflow struct
	json.Unmarshal([]byte(metricsConfig), &workflow)

	return &workflow
}

func (wf *Workflow) PromoteWorkflowMetrics() {

	for _, metric := range wf.Metrics {
		// change metric Strategy if necessary
		metric.DetermineMetricStrategy()
		// advance metric value in relation to metric strategy
		metric.AdvanceMetricValue()
	}

	// reset timestamp
	wf.Timestamp = time.Now()
}

func (wf *Workflow) ReturnWorkflowData(w http.ResponseWriter, r *http.Request) {

	// decode workflow data to json
	json.NewEncoder(w).Encode(wf)
}
