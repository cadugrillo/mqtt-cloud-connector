package Mqttbuffer

import (
	"errors"
	"fmt"
)

type Mqttbuffer struct {
	Buffer       [5000]Message
	ReadPointer  int
	WritePointer int
}

// Message
type Message struct {
	Duplicate bool
	Qos       byte
	Retained  bool
	Topic     string
	MessageID uint16
	Payload   string
	Ack       bool
}

func NewMqttbuffer() Mqttbuffer {
	b := Mqttbuffer{}
	return b
}

func (b Mqttbuffer) GetReadPointer() int {
	return b.ReadPointer
}

func (b Mqttbuffer) GetWritePointer() int {
	return b.WritePointer
}

func (b Mqttbuffer) AddMessage(message Message) ([5000]Message, int) {
	if b.WritePointer == len(b.Buffer)-1 {
		b.Buffer[b.WritePointer] = message
		b.WritePointer = 0
		return b.Buffer, b.WritePointer
	}
	b.Buffer[b.WritePointer] = message
	b.WritePointer++
	return b.Buffer, b.WritePointer
}

func (b Mqttbuffer) ReadMessage(index int) (Message, error) {
	if index < len(b.Buffer) {
		return b.Buffer[index], nil
	}
	msg := Message{}
	return msg, errors.New(fmt.Sprintf("Index %d greater then buffer size [%d]", index, len(b.Buffer)))
}

func (b Mqttbuffer) NextMessage() int {
	if b.ReadPointer == len(b.Buffer)-1 {
		b.ReadPointer = 0
		return b.ReadPointer
	}
	if b.ReadPointer != b.WritePointer {
		b.ReadPointer++
		return b.ReadPointer
	}
	fmt.Println("No new messages on the buffer")
	return b.ReadPointer
}

func (b Mqttbuffer) NewMessage() bool {
	return b.WritePointer != b.ReadPointer
}
