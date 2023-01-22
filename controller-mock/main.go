package main

import (
	"time"

	"github.com/RHEcosystemAppEng/abb-dcs-poc/controller-mock/pkg/metrics"
)

func main() {

	metrics := metrics.InitMetrics()

	go runDataLoop(metrics)

	metrics.HandleRequests()
}

func runDataLoop(m *metrics.Metrics) {

	// promote/update metrics data every 1 second
	for {
		m.PromoteMetrics()
		time.Sleep(1 * time.Second)
	}
}
