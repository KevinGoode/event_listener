package event_listener

import stan "github.com/nats-io/go-nats-streaming"
import (
	"sync"
)

//MessageBuffer maintains a thread safe list of messages
type MessageBuffer struct {
	messages []*stan.Msg
	lock     *sync.Mutex
}

//Append adds a message to end of buffer
func (buffer *MessageBuffer) Append(message *stan.Msg) {
	buffer.lock.Lock()
	buffer.messages = append(buffer.messages, message)
	buffer.lock.Unlock()
}

//GetFirst returns oldest message. If no message left then return nil
func (buffer *MessageBuffer) GetFirst() *stan.Msg {
	var first *stan.Msg
	buffer.lock.Lock()
	if len(buffer.messages) > 0 {
		first = buffer.messages[0]
	}
	buffer.lock.Unlock()
	return first
}

//RemoveFirst remove oldest message. Call after forward success
func (buffer *MessageBuffer) RemoveFirst() {
	buffer.lock.Lock()
	if len(buffer.messages) > 0 {
		buffer.messages = buffer.messages[1:]
	}
	buffer.lock.Unlock()
}

//NewMessageBuffer creates and initialises a MessageBuffer
func NewMessageBuffer() *MessageBuffer {
	buffer := MessageBuffer{}
	buffer.lock = &sync.Mutex{}
	return &buffer
}
