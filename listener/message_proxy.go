package event_listener

import "fmt"

//MessageProxy is a container for proxy
type MessageProxy struct {
	listener  *MessageListener
	forwarder *MessageForwarder
	buffer    *MessageBuffer
}

//Start starts listening for messages and starts forwarder
func (proxy *MessageProxy) Start() error {
	fmt.Printf("Starting proxy...\n")
	err := proxy.forwarder.Start()
	if err == nil {
		err = proxy.listener.Start()
	}
	fmt.Printf("Finished starting proxy...\n")
	return err
}

//Stop stops listening for messages and stops forwarder
func (proxy *MessageProxy) Stop() error {
	fmt.Printf("Stopping proxy...\n")
	err := proxy.listener.Stop()
	if err == nil {
		err = proxy.forwarder.Stop()
	}
	fmt.Printf("Finished stopping proxy...\n")
	return nil
}

// NewMessageProxy creates and initialises a MessageProxy
func NewMessageProxy(handler *Handler) *MessageProxy {
	fmt.Printf("Creating new proxy . CliendID %s, MessageId %s, Port: %s,Last Message %s \n",
		handler.ClientId, handler.MessageId, fmt.Sprint(handler.Port), fmt.Sprint(handler.LastMessageNumber))
	proxy := MessageProxy{}
	proxy.buffer = NewMessageBuffer()
	proxy.forwarder = NewMessageForwarder(handler.Port, proxy.buffer)
	proxy.listener = NewMessageListener(handler.ClientId, handler.MessageId, handler.LastMessageNumber, proxy.buffer)
	return &proxy
}
