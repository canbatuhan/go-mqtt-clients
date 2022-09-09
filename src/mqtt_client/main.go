package main

import (
	client "mqtt_client/client"
	"time"
)

var clientID string = "Prototype-0"
var tcpPort string = "8080"
var mqttBroker string = "broker.hivemq.com"
var mqttPort uint16 = 1883
var topic string = "racoons-talking"

func main() {
	c := client.InitClient(clientID, tcpPort, mqttBroker, mqttPort, topic)
	go client.UIRoutine(&c)
	go client.SubscriberRoutine(&c)
	go client.PublisherRoutine(&c)
	time.Sleep(time.Hour) // Active for 1 hour
}
