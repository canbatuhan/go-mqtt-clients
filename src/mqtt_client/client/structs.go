package client

import "encoding/json"

var CAPACITY uint16 = 100

/*********************/
/****** REQUEST ******/
/*********************/

type Request struct {
	Command string
	Name    string
	Size    uint16
}

func ByteArrayToRequest(data []byte) Request {
	var newRequest Request
	json.Unmarshal(data, &newRequest)
	return newRequest
}

func RequestToByteArray(request Request) []byte {
	byteArray, err := json.Marshal(request)
	if err != nil {
		return nil
	} else {
		return byteArray
	}
}

/*********************/
/****** MESSAGE ******/
/*********************/

type Message struct {
	Source      string
	Destination string
	Body        Request
}

func InitMessage(source string, destination string, body Request) Message {
	var newMessage Message
	newMessage.Source = source
	newMessage.Destination = destination
	newMessage.Body = body
	return newMessage
}

func ByteArrayToMessage(data []byte) Message {
	var newMessage Message
	json.Unmarshal(data, &newMessage)
	return newMessage
}

func MessageToByteArray(message Message) []byte {
	byteArray, err := json.Marshal(message)
	if err != nil {
		return nil
	} else {
		return byteArray
	}
}

/******************/
/****** NODE ******/
/******************/

type Node struct {
	Data Message
	Next *Node
	Prev *Node
}

func InitNode(data Message) Node {
	var newNode Node
	newNode.Data = data
	newNode.Next = nil
	return newNode
}

/*******************/
/****** QUEUE ******/
/*******************/

type Queue struct {
	Capacity uint16
	Size     uint16
	Front    *Node
	Back     *Node
}

func InitQueue() Queue {
	var newQueue Queue
	newQueue.Capacity = CAPACITY
	newQueue.Size = 0
	newQueue.Front = nil
	newQueue.Back = nil
	return newQueue
}

func IsEmpty(q *Queue) bool {
	return q.Size == 0
}

func IsFull(q *Queue) bool {
	return q.Size == CAPACITY
}

func Push(q *Queue, node *Node) {
	if !IsFull(q) {
		if q.Size == 0 {
			q.Front = node
			q.Back = node
			q.Front.Prev = q.Back
			q.Back.Next = q.Front
		} else {
			q.Back.Prev = node
			node.Next = q.Back
			q.Back = node
		}
		q.Size = q.Size + 1
	}
}

func Pop(q *Queue) *Node {
	if !IsEmpty(q) {
		popNode := q.Front
		q.Front = q.Front.Prev
		q.Front.Next = nil
		popNode.Prev = nil
		q.Size = q.Size - 1
		return popNode
	} else {
		return nil
	}
}
