package event_listener

import (
	context "context"
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//RegistryServer is a container for registry server that provides the GRPC API that supports regsitration requests
type RegistryServer struct {
	server   *grpc.Server
	registry *Registry
}

//RegisterEventHandler is GRPC callback that handles requests from clients to be sent messages on a given subject
func (srv *RegistryServer) RegisterEventHandler(ctx context.Context, handler *Handler) (*RegisterResponse, error) {
	var err error
	out := new(RegisterResponse)
	fmt.Printf("Received a  registration request \n")
	//Check there is no registration already from this server for this subject
	proxy := srv.registry.AlreadyRegisteredClientAndMessage(handler.ClientId, handler.MessageId)
	if proxy != nil {
		fmt.Printf("Received a registration request for the second time for client %s message %s\n", handler.ClientId, handler.MessageId)
		//Delete old registration and create new because chances are sequence number has changed so need
		//to pick up new sequnce number
		fmt.Printf("Deleting duplicate proxy..")
		srv.registry.DeleteProxy(handler.ClientId, handler.MessageId)
		fmt.Printf("Deleted duplicate proxy")
	}
	err = srv.registry.AddNewProxy(handler)

	return out, err
}

//Start starts the GRPC server that listens for registration requests. This function is called at  process start
func (srv *RegistryServer) Start(port string) {
	fmt.Printf("Starting registry server on port %s\n", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		os.Exit(-1)
	}
	srv.server = grpc.NewServer()

	fmt.Printf("Listening for registrations on port %s\n", port)
	RegisterRegistrationListenerServer(srv.server, srv)
	// Register reflection service on gRPC server.
	reflection.Register(srv.server)
	fmt.Printf("Waiting for registrations.....\n")
	// This call blocks but signal handler is called before exit
	if err := srv.server.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v", err)
		os.Exit(-1)
	}
}

//Stop stops the GRPC server that listens for registration requests. This function is called at  process stop
func (srv *RegistryServer) Stop() {
	if srv.server != nil {
		fmt.Printf("Stopping registration server..")
		srv.server.Stop()
	}
}

// NewRegistryServer creates and initialises a RegistryServer
func NewRegistryServer(registry *Registry) *RegistryServer {
	server := RegistryServer{}
	server.registry = registry
	return &server
}
