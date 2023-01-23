package workflow

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Workflow struct {
	Uuid      uuid.UUID `json:"workflow_Id"`
	TimeStamp time.Time `json:"timestamp"`
	Metrics   *Metrics  `json:"metrics"`
}

func InitWorkflow() *Workflow {

	// init workflow with unique id and metrics
	workflow := Workflow{
		Uuid:      uuid.New(),
		TimeStamp: time.Now(),
		Metrics:   InitMetrics(),
	}

	return &workflow
}

func (wf *Workflow) PromoteWorkflowMetrics() {

	// change metric Strategy if necessary
	wf.Metrics.MotorTemp.DetermineMetricStrategy()
	wf.Metrics.MotorRPM.DetermineMetricStrategy()
	wf.Metrics.MotorNoise.DetermineMetricStrategy()

	// advance metric value in relation to metric strategy
	wf.Metrics.MotorTemp.AdvanceMetricValue()
	wf.Metrics.MotorRPM.AdvanceMetricValue()
	wf.Metrics.MotorNoise.AdvanceMetricValue()

	// set new timeStamp
	wf.TimeStamp = time.Now()
}

func (wf *Workflow) ReturnWorkflowData(w http.ResponseWriter, r *http.Request) {

	// decode workflow data to json
	json.NewEncoder(w).Encode(wf)
}
