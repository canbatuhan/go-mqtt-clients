package client

import (
	"fmt"
	"time"
)

func UIRoutine(client *Client) {
	for {
		conn, _ := client.tcpListener.Accept()
		fmt.Printf("[UserInterface]\t- Connection is made with a user\n")

		for {
			buffer := make([]byte, 4096)
			size, err := conn.Read(buffer)

			if err != nil {
				fmt.Printf("[UserInterface]\t- Connection is lost with the user: %v\n", err)
				break
			}

			receviedUserReq := ByteArrayToRequest(buffer[:size])
			fmt.Printf("[UserInterface]\t- USER-REQ recevied: %v\n", receviedUserReq)

			newNode := InitNode(InitMessage(client.id, "broadcast", receviedUserReq))
			Push(client.requestQueue, &newNode)
			fmt.Printf("[UserInterface]\t- USER-REQ pushed to RequestQueue\n")
		}
	}
}

func SubscriberRoutine(client *Client) {
	token := client.mqttClient.Subscribe(client.topic, 1, nil)
	token.Wait()
	for {
		if !IsEmpty(client.messageQueue) {
			mqttMsgToProcess := Pop(client.messageQueue).Data
			fmt.Printf("[Subscriber]\t- MQTT-MSG received from %v: Message:%v\n",
				mqttMsgToProcess.Source, mqttMsgToProcess.Body)
		}
	}
}

func PublisherRoutine(client *Client) {
	for {
		if IsEmpty(client.messageQueue) && !IsEmpty(client.requestQueue) {
			UserMessageToProcess := Pop(client.requestQueue).Data

			byteArray := MessageToByteArray(UserMessageToProcess)
			token := client.mqttClient.Publish(client.topic, 0, false, byteArray)
			fmt.Printf("[Publisher]\t- USER-REQ published\n")

			token.Wait()
			time.Sleep(time.Second)
		}
	}
}
