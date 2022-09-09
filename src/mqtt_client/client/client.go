package client

import (
	"fmt"
	"net"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var REQUEST_QUEUE Queue = InitQueue()
var MESSAGE_QUEUE Queue = InitQueue()

// global MQTT pub message processing
var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	messageAsNode := InitNode(ByteArrayToMessage(msg.Payload()))
	Push(&MESSAGE_QUEUE, &messageAsNode)
}

// callback for the connection
var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected to MQTT Broker")
}

// callback for connection loss
var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connection is lost: %v", err)
}

type Client struct {
	id           string
	tcpListener  net.Listener
	mqttClient   mqtt.Client
	messageQueue *Queue
	requestQueue *Queue
	topic        string
}

func InitMqttClient(id string, broker string, port uint16) mqtt.Client {
	// Initialize
	options := mqtt.NewClientOptions()
	options.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	options.SetClientID(id)

	// Implemented in handler.go
	options.SetDefaultPublishHandler(messagePubHandler)
	options.OnConnect = connectHandler
	options.OnConnectionLost = connectLostHandler

	// Error check
	client := mqtt.NewClient(options)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// Return client
	return client
}

func InitClient(id string, tcp_port string, mqtt_broker string, mqtt_port uint16, topic string) Client {
	var newClient Client
	newClient.id = id
	newClient.tcpListener, _ = net.Listen("tcp", "127.0.0.1:"+tcp_port)
	newClient.mqttClient = InitMqttClient(id, mqtt_broker, mqtt_port)
	newClient.messageQueue = &MESSAGE_QUEUE // global
	newClient.requestQueue = &REQUEST_QUEUE // global
	newClient.topic = topic
	return newClient
}
