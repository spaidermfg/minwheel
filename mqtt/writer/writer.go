package main

import (
	"log"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// write data to mqtt
func main() {
	mqtt.ERROR = log.New(os.Stdout, "[ERROR]", 0)

}

type MqttWriter struct{}
