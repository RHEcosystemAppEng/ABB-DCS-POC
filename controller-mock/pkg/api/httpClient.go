package api

import (
	"log"
	"net/http"

	"github.com/RHEcosystemAppEng/abb-dcs-poc/controller-mock/pkg/workflow"
)

func HandleHttpRequests(wf *workflow.Workflow) {

	// route workflow data to http://localhost:8080/workflow
	http.HandleFunc("/workflow", wf.ReturnWorkflowData)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
