package main

import (
	"time"

	"github.com/RHEcosystemAppEng/abb-dcs-poc/controller-mock/pkg/metrics"
)

func main() {

	// init metrics with initial data
	mtrx := metrics.InitMetrics()

	// concurrently promote metrics with new data
	go runDataLoop(mtrx)

	// route metric data over http
	// metrics.HandleRequests()

	// route metric data over tcp
	mtrx.SendMetricsOverTCP()
}

func runDataLoop(m *metrics.Metrics) {

	// promote/update metrics data every 1 second
	for {
		m.PromoteMetrics()
		time.Sleep(1 * time.Second)
	}
}
