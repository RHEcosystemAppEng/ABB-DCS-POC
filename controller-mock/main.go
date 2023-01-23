package main

import (
	"time"

	"github.com/RHEcosystemAppEng/abb-dcs-poc/controller-mock/pkg/api"
	"github.com/RHEcosystemAppEng/abb-dcs-poc/controller-mock/pkg/metrics"
)

const (
	TIME_INTERVAL = 1
)

func main() {

	// init metrics with initial data
	mtrx := metrics.InitMetrics()

	// route initial metric data over tcp
	api.SendMetricsOverTCP(mtrx)

	// in an endless loop, every time interval, promote metrics data and send to server over tcp
	for {

		// wait time interval
		time.Sleep(TIME_INTERVAL * time.Second)

		// promote metrics data
		mtrx.PromoteMetrics()

		// route promoted metric data over tcp
		api.SendMetricsOverTCP(mtrx)
	}

	// // route metric data over http
	// api.HandleHttpRequests(mtrx)
}
