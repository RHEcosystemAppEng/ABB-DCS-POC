package kafka

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func ReadConfig(configFile string) kafka.ConfigMap {

	// init kafka configmap
	cm := make(map[string]kafka.ConfigValue)

	// get config file content from config filepath
	file, err := os.Open(configFile)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	// marshal config file content into configmap
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if !strings.HasPrefix(line, "#") && len(line) != 0 {
			kv := strings.Split(line, "=")
			parameter := strings.TrimSpace(kv[0])
			value := strings.TrimSpace(kv[1])
			cm[parameter] = value
		}
	}

	// check for scanner errors
	if err := scanner.Err(); err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	return cm

}
