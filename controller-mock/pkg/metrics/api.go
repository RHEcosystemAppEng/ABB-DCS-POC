package metrics

import (
	"encoding/json"
	"log"
	"net/http"
)

func (m *Metrics) ReturnAllMetrics(w http.ResponseWriter, r *http.Request) {

	// decode metrics data to json
	json.NewEncoder(w).Encode(m)
}

func (m *Metrics) HandleRequests() {

	// route metrics data to http://localhost:8080/metrics
	http.HandleFunc("/metrics", m.ReturnAllMetrics)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
