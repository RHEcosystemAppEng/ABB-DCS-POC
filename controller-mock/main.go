package main

import (
	"time"

	"github.com/RHEcosystemAppEng/abb-dcs-poc/controller-mock/pkg/api"
	"github.com/RHEcosystemAppEng/abb-dcs-poc/controller-mock/pkg/workflow"
)

const (
	TIME_INTERVAL = 1
)

func main() {

	// init Workflow with initial metrics data
	wf := workflow.InitWorkflow()

	// route initial workflow data over tcp
	api.SendWorkflowDataOverTCP(wf)

	// in an endless loop, every time interval, promote workflow metrics data and send to server over tcp
	for {

		// wait time interval
		time.Sleep(TIME_INTERVAL * time.Second)

		// promote workflow metrics data
		wf.PromoteWorkflowMetrics()

		// route workflow data with promoted metrics over tcp
		api.SendWorkflowDataOverTCP(wf)
	}

	// // route workflow data over http
	// api.HandleHttpRequests(wf)
}
