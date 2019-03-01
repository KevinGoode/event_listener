package event_listener

import (
	"fmt"
	"os"
	"syscall"
)

const (
	port = ":50051"
)

//App is main application
type App struct {
	registry       *Registry
	registryServer *RegistryServer
}

//NewApp Creates new application
func NewApp() *App {
	app := App{}
	return &app
}

//HandleSignal Handles keyboard signal
func (app *App) HandleSignal(os.Signal) {
	fmt.Printf("***************************\n")
	fmt.Printf("**Received signal to stop**\n")
	//Stop listening for registration attempts
	app.registryServer.Stop()
	//Stopping the registry does cleanup of message listener
	app.registry.Stop()
	fmt.Printf("***************************\n")
	os.Exit(0)
}

//Run runs application
func (app *App) Run() {
	//1.) Setup signal handlers
	fmt.Printf("***************************\n")
	fmt.Printf("**Starting event_listener**\n")
	signaller := NewSignaller()
	signaller.HandleSignal(syscall.SIGTERM, app)
	signaller.HandleSignal(syscall.SIGINT, app)

	app.registry = NewRegistry()
	app.registryServer = NewRegistryServer(app.registry)
	app.registryServer.Start(port)
}
