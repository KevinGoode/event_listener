package main

import (
	"context"
	event_listener "event_listener/listener"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//Server listens for event messages amnd prints then
type Server struct {
	grpcServer *grpc.Server
	Port       string
}

//Start starts listening for messages
func (serv *Server) Start() {
	fmt.Printf("Starting to listen for events...\n")
	//Listen in the background
	go serv.listen()
}

//Stop stops listening for messages
func (serv *Server) Stop() {
	fmt.Printf("Stopping listening\n")
	serv.grpcServer.Stop()
	fmt.Printf("Stopped...\n")
}
func (serv *Server) listen() {
	fmt.Printf("Listening for events on port %s\n", serv.Port)
	lis, err := net.Listen("tcp", serv.Port)
	if err != nil {
		fmt.Printf("Failed to listen: %v\n", err)
	}
	serv.grpcServer = grpc.NewServer()
	event_listener.RegisterGenericEventListenerServer(serv.grpcServer, serv)
	// Register reflection service on gRPC server.
	reflection.Register(serv.grpcServer)
	//Blocking call
	if err := serv.grpcServer.Serve(lis); err != nil {
		fmt.Printf("Failed or finished serving: %v\n", err)
	}
}

//GenericEvent recives event and prints it
func (serv *Server) GenericEvent(ctx context.Context, message *event_listener.GenericEventMessage) (*event_listener.GenericEventResponse, error) {
	response := event_listener.GenericEventResponse{}
	fmt.Printf("Received message : '%s'\n", message.Message)

	return &response, nil
}
