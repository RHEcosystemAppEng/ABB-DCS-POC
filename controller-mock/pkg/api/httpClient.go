package api

import (
	"log"
	"net/http"

	"github.com/RHEcosystemAppEng/abb-dcs-poc/controller-mock/pkg/controller"
)

func HandleHttpRequests(wf *controller.Controller) {

	// route controller data to http://localhost:8080/controller
	http.HandleFunc("/controller", wf.ReturnControllerData)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
