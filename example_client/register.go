package main

import (
	"context"
	event_listener "event_listener/listener"
	"fmt"
	"time"

	"google.golang.org/grpc"
)

// Registerer regsiters interest in events
type Registerer struct {
	port       uint32
	connection *grpc.ClientConn
	client     event_listener.RegistrationListenerClient
}

//Register registers interest in events with event_listener
func (registerer *Registerer) Register(client string, subject string) error {
	fmt.Printf("Registering interest in generic event...\n")

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	message := event_listener.Handler{}
	message.Port = registerer.port
	message.ClientId = client
	message.MessageId = subject
	fmt.Printf("Sending registration message. ClientID %s MessageID %s Port %s\n", message.ClientId, message.MessageId, fmt.Sprint(message.Port))
	_, err := registerer.client.RegisterEventHandler(ctx, &message)

	return err
}

//CheckStillRegistered checks still registered
func (registerer *Registerer) CheckStillRegistered(client string, subject string) error {
	fmt.Printf("Registering interest in generic event...\n")

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	message := event_listener.Handler{}
	message.ClientId = client
	message.MessageId = subject
	fmt.Printf("Sending check registration message. ClientID %s MessageID %s Port %s\n", message.ClientId, message.MessageId, fmt.Sprint(message.Port))
	_, err := registerer.client.GetEventHandlerStatus(ctx, &message)
	return err
}

//Connect connects to event_listener
func (registerer *Registerer) Connect(targetPort string) error {
	var address = "localhost" + targetPort
	fmt.Printf("Connecting to: %s...\n", targetPort)
	var err error
	registerer.connection, err = grpc.Dial(address, grpc.WithInsecure())
	if err == nil {
		fmt.Printf("Connected to: %s...\n", targetPort)
		registerer.client = event_listener.NewRegistrationListenerClient(registerer.connection)
	}
	return err
}

//Disconnect disconnects from event_listener
func (registerer *Registerer) Disconnect() {
	fmt.Printf("Disconnecting...\n")
	registerer.connection.Close()
}

// NewRegisterer creates and initialises a Registerer
func NewRegisterer(port uint32) *Registerer {
	registerer := Registerer{}
	registerer.port = port
	return &registerer
}
