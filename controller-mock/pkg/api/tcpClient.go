package api

import (
	"encoding/json"
	"log"
	"net"

	"github.com/RHEcosystemAppEng/abb-dcs-poc/controller-mock/pkg/controller"
)

const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

func SendControllerDataOverTCP(c *controller.Controller) {

	// define tcp address
	tcpServer, err := net.ResolveTCPAddr(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatalf("ResolveTCPAddr failed: %s", err)
	}

	// connect to tcp network
	conn, err := net.DialTCP(TYPE, nil, tcpServer)
	if err != nil {
		log.Fatalf("Dial failed: %s", err)
	}
	defer conn.Close()

	// convert controller struct to json packet
	cJson, err := json.Marshal(c)
	if err != nil {
		log.Fatalf("Marshaling controller data to JSON failed: %s", err)
	}

	// write message through network connection
	_, err = conn.Write([]byte(cJson))
	if err != nil {
		log.Fatalf("Write data failed: %s", err)
	}

	// get response from tcp server
	response := make([]byte, 1024)
	_, err = conn.Read(response)
	if err != nil {
		log.Fatalf("Read data failed: %s", err)
	}
	println("Response:", string(response))
}
