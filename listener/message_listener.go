package event_listener

import (
	fmt "fmt"
	"time"

	stan "github.com/nats-io/go-nats-streaming"
)

//MessageListener is container for NAT configuration and connection parameters
type MessageListener struct {
	ClientID string
	//MessageID is used as subject of message
	MessageID         string
	lastMessageNumber uint64
	messageBuffer     *MessageBuffer
	subOption         stan.SubscriptionOption
	connection        stan.Conn
	subscription      stan.Subscription
}

//Start starts listening for NATs messages
func (listener *MessageListener) Start() error {
	fmt.Printf("Starting listener...\n")
	var err error
	var durable = ""
	var qgroup = ""
	err = listener.connect()
	if err != nil {
		err = listener.reconnect(err)
	}
	if err == nil {
		messageCallack := func(msg *stan.Msg) {
			listener.messageHandler(msg)
		}
		listener.subscription, err = listener.connection.QueueSubscribe(listener.MessageID, qgroup, messageCallack, listener.subOption, stan.DurableName(durable))
		if err == nil {
			fmt.Printf("Successfully subscribed\n")
		} else {
			fmt.Printf("Failed to subscribe, reason: %v\n", err)
		}
	}
	fmt.Printf("Finished starting listener...\n")
	return err
}

//Stop stops listening for NATs messages
func (listener *MessageListener) Stop() error {
	fmt.Printf("Stopping listener...\n")
	if listener.subscription != nil {
		fmt.Printf("Unsubscribing...\n")
		listener.subscription.Unsubscribe()
	}
	if listener.connection != nil {
		fmt.Printf("Closing connection...\n")
		listener.connection.Close()
	}
	fmt.Printf("Finished stopping listener...\n")
	return nil
}

//messageHandler is callback that is called when message received
func (listener *MessageListener) messageHandler(msg *stan.Msg) {
	fmt.Printf("Received a message %s\n", msg)
	listener.messageBuffer.Append(msg)
}
func (listener *MessageListener) connect() error {
	fmt.Printf("Connecting to message server...\n")
	var err error
	var clusterID = "test-cluster"
	listener.connection, err = stan.Connect(clusterID, listener.ClientID, stan.NatsURL(stan.DefaultNatsURL),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			fmt.Printf("Connection lost, reason: %v\n", reason)
			listener.connectionLossHandler(reason)
		}))
	if err == nil {
		fmt.Printf("Successfully connected to message server\n")
	} else {
		fmt.Printf("Failed to connect to message server\n")
	}
	return err
}
func (listener *MessageListener) connectionLossHandler(reason error) {
	fmt.Printf("Disconnected from message server, reason: %v\n", reason)
	fmt.Printf("Trying to reconnect......\n")
	listener.reconnect(reason)
	//TODO DO I need to resubscribe??
}
func (listener *MessageListener) reconnect(reason error) error {
	var maxCount uint = 100
	var count uint
	var err error
	for err = reason; err != nil; err = listener.connect() {
		count++
		fmt.Printf("Attempting to re-connect to message server...\n")
		if count > maxCount {
			fmt.Printf("Cannot reconnect after 10 mins, reason: %v\n", err)
			break
		}
		time.Sleep(6 * time.Second)
	}
	return err
}

//NewMessageListener constructs a MessageListener
func NewMessageListener(clientID, messageID string, lastMessage uint64, messageBuffer *MessageBuffer) *MessageListener {

	listener := MessageListener{}
	listener.subOption = stan.DeliverAllAvailable()
	if lastMessage > 0 {
		listener.subOption = stan.StartAtSequence(lastMessage)
	}
	listener.messageBuffer = messageBuffer
	listener.ClientID = clientID
	listener.MessageID = messageID
	listener.lastMessageNumber = lastMessage
	return &listener
}
