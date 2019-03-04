package event_listener

import "fmt"

//MessageProxy is a container for proxy
type MessageProxy struct {
	listener  ListenerAPI
	forwarder *MessageForwarder
	buffer    *MessageBuffer
}

//ProxyAPI exposes MessageProxy services
type ProxyAPI interface {
	Start() error
	Stop() error
	IsClientAndMessage(clientID string, messageID string) bool
}

//IsClientAndMessage returns true if proxy client and subject match arguments
func (proxy *MessageProxy) IsClientAndMessage(clientID string, messageID string) bool {
	if proxy.listener.GetClientID() == clientID && proxy.listener.GetMessageID() == messageID {
		return true
	}
	return false
}

//Start starts listening for messages and starts forwarder
func (proxy *MessageProxy) Start() error {
	fmt.Printf("Starting proxy...\n")
	err := proxy.forwarder.Start()
	if err == nil {
		err = proxy.listener.Subscribe()
	}
	fmt.Printf("Finished starting proxy...\n")
	return err
}

//Stop stops listening for messages and stops forwarder
func (proxy *MessageProxy) Stop() error {
	fmt.Printf("Stopping proxy...\n")
	err := proxy.listener.Unsubscribe()
	if err == nil {
		err = proxy.forwarder.Stop()
	}
	fmt.Printf("Finished stopping proxy...\n")
	return nil
}

// NewMessageProxy creates and initialises a MessageProxy
func NewMessageProxy(handler *Handler) ProxyAPI {
	fmt.Printf("Creating new proxy . CliendID %s, MessageId %s, Port: %s,Last Message %s \n",
		handler.ClientId, handler.MessageId, fmt.Sprint(handler.Port), fmt.Sprint(handler.LastMessageNumber))
	proxy := MessageProxy{}
	proxy.buffer = NewMessageBuffer()
	proxy.forwarder = NewMessageForwarder(handler.Port, proxy.buffer)
	proxy.listener = NewMessageListener(handler.ClientId, handler.MessageId, handler.LastMessageNumber, proxy.buffer)
	return &proxy
}
