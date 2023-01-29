package main

import (
	"time"

	"github.com/RHEcosystemAppEng/abb-dcs-poc/controller-mock/pkg/controller"
	"github.com/RHEcosystemAppEng/abb-dcs-poc/controller-mock/pkg/kafka"
)

const (
	TIME_INTERVAL = 1
)

func main() {

	// init controller with initial metrics data
	c := controller.InitController()

	// route initial controller data over kafka
	kafka.HTTPKafkaProducer(c)

	// in an endless loop, every time interval, promote controller metrics data and send to server over kafka
	for {

		// wait time interval
		time.Sleep(TIME_INTERVAL * time.Second)

		// promote controller metrics data
		c.PromoteControllerMetrics()

		// route controller data with promoted metrics over kafka
		kafka.HTTPKafkaProducer(c)
	}

	// // route initial controller data over tcp
	// api.SendControllerDataOverTCP(c)

	// // route controller data over http
	// api.HandleHttpRequests(c)
}
