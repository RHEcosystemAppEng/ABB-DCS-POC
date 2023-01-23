package metrics

import (
	"encoding/json"
	"log"
	"net"
)

const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

func (m *Metrics) SendMetricsOverTCP() {

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

	// convert metrics struct to json packet
	mJson, err := json.Marshal(m)
	if err != nil {
		log.Fatalf("Marshaling metrics to JSON failed: %s", err)
	}

	// write message through network connection
	_, err = conn.Write([]byte(mJson))
	if err != nil {
		log.Fatalf("Write data failed: %s", err)
	}

	// get response from tcp server
	response := make([]byte, 1024)
	_, err = conn.Read(response)
	if err != nil {
		log.Fatalf("Read data failed: %s", err)
	}
	println("Response message:", string(response))
}
