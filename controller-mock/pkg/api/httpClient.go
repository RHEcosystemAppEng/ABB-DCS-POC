package api

import (
	"log"
	"net/http"

	"github.com/RHEcosystemAppEng/abb-dcs-poc/controller-mock/pkg/metrics"
)

func HandleHttpRequests(m *metrics.Metrics) {

	// route metrics data to http://localhost:8080/metrics
	http.HandleFunc("/metrics", m.ReturnAllMetrics)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
