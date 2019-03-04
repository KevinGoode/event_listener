package event_listener

import (
	fmt "fmt"

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
	connection        MessagerConnectionAPI
	subscription      stan.Subscription
}

//ListenerAPI exposes MessageListener services
type ListenerAPI interface {
	Subscribe() error
	Unsubscribe() error
}

//Subscribe starts listening for NATs messages
func (listener *MessageListener) Subscribe() error {
	fmt.Printf("Subscribing : ClientID :%s subject :%s...\n", listener.ClientID, listener.MessageID)
	var err error
	var durable = ""
	var qgroup = ""
	if err == nil {
		messageCallack := func(msg *stan.Msg) {
			listener.messageHandler(msg)
		}
		listener.subscription, err = listener.connection.GetConnection().QueueSubscribe(listener.MessageID, qgroup, messageCallack, listener.subOption, stan.DurableName(durable))
		if err == nil {
			fmt.Printf("Successfully subscribed\n")
		} else {
			fmt.Printf("Failed to subscribe, reason: %v\n", err)
		}
	}
	fmt.Printf("Finished subscribing...\n")
	return err
}

//Unsubscribe stops listening for NATs messages
func (listener *MessageListener) Unsubscribe() error {
	fmt.Printf("Stopping listener...\n")
	if listener.subscription != nil {
		fmt.Printf("Unsubscribing...\n")
		listener.subscription.Unsubscribe()
	}
	fmt.Printf("Finished stopping listener...\n")
	return nil
}

//messageHandler is callback that is called when message received
func (listener *MessageListener) messageHandler(msg *stan.Msg) {
	fmt.Printf("Received a message %s\n", msg)
	listener.messageBuffer.Append(msg)
}

//NewMessageListener constructs a MessageListener
func NewMessageListener(clientID, messageID string, lastMessage uint64, messageBuffer *MessageBuffer) ListenerAPI {

	listener := MessageListener{}
	listener.subOption = stan.DeliverAllAvailable()
	if lastMessage > 0 {
		listener.subOption = stan.StartAtSequence(lastMessage)
	}
	listener.messageBuffer = messageBuffer
	listener.ClientID = clientID
	listener.MessageID = messageID
	listener.lastMessageNumber = lastMessage
	listener.connection = CreateMessagerConnection(clientID, &listener)
	return &listener
}
