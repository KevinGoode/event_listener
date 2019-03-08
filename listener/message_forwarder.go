package event_listener

import (
	"context"
	fmt "fmt"
	"time"

	stan "github.com/nats-io/go-nats-streaming"
	"google.golang.org/grpc"
)

// MessageForwarder forwards messages to clients using GRPC
type MessageForwarder struct {
	port          uint32
	messageBuffer *MessageBuffer
	connection    *grpc.ClientConn
	client        GenericEventListenerClient
	stop          chan bool
	stopped       chan bool
}

//Start connects to client and the starts a background thread that periodically empties message buffer
func (forwarder *MessageForwarder) Start() error {
	fmt.Printf("Starting forwarder...\n")
	err := forwarder.connect()
	if err == nil {
		fmt.Printf("Forwarder task starting....\n")
		//Process message queue in background
		go forwarder.processMessageQueue()
	}
	fmt.Printf("Finished starting forwarder\n")
	return err
}

//Stop stops background thread then disconnects from client
func (forwarder *MessageForwarder) Stop() error {
	fmt.Printf("Stopping forwarder...\n")
	var err error
	if forwarder.connection != nil {
		//Stop forwarder thread
		forwarder.stop <- true
		//Block waiting for stopped
		<-forwarder.stopped
		fmt.Printf("Stopped forwarder thread\n")
		err = forwarder.connection.Close()
	}
	fmt.Printf("Stopped forwarder\n")
	return err
}

//Stop stops background thread then disconnects from client
func (forwarder *MessageForwarder) connect() error {
	//TODO probably want to try to reconnect here
	var err error
	var address = "localhost:" + fmt.Sprint(forwarder.port)
	forwarder.connection, err = grpc.Dial(address, grpc.WithInsecure())
	if err == nil {
		fmt.Printf("Forwarder successfuly connected\n")
		forwarder.client = NewGenericEventListenerClient(forwarder.connection)
	} else {
		fmt.Printf("Forwarder failed to connect. Reason: %v\n", err)
	}
	return err
}
func (forwarder *MessageForwarder) processMessageQueue() {
	for {
		select {
		case <-forwarder.stop:
			//Set stopped flag
			forwarder.stopped <- true
			return
		default:
			// Process all messages in buffer
			for mess := forwarder.messageBuffer.GetFirst(); mess != nil; mess = forwarder.messageBuffer.GetFirst() {
				err := forwarder.forward(mess)
				if err == nil {
					forwarder.messageBuffer.RemoveFirst()
				} else {
					break
				}
			}
			//Sleep for a second
			time.Sleep(time.Second)
		}
	}
}
func (forwarder *MessageForwarder) forward(mess *stan.Msg) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()
	message := GenericEventMessage{}
	message.Message = mess.String()
	fmt.Printf("Forwarding a message %s\n", mess)
	_, err := forwarder.client.GenericEvent(ctx, &message)
	if err == nil {
		fmt.Printf("Successfully sent message\n")
	} else {
		fmt.Printf("Failed to forward message, reason: %v\n", err)
	}
	return err
}

// NewMessageForwarder creates and initialises a MessageForwarder
func NewMessageForwarder(port uint32, messageBuffer *MessageBuffer) *MessageForwarder {
	forwarder := MessageForwarder{}
	forwarder.stop = make(chan bool)
	forwarder.stopped = make(chan bool)
	forwarder.messageBuffer = messageBuffer
	forwarder.port = port
	return &forwarder
}
