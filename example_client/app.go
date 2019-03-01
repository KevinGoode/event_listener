package main

import (
	event_listener "event_listener/listener"
	"fmt"
	"os"
	"syscall"
	"time"
)

//App is entry point for app
type App struct {
	server Server
}

//NewApp creates app
func NewApp() *App {
	app := App{}
	return &app
}

//HandleSignal handles signals from keyboard. Stops listening for events
func (app *App) HandleSignal(os.Signal) {
	fmt.Printf("***************************\n")
	fmt.Printf("**Received signal to stop**\n")
	//Stop listening for registration attempts
	app.server.Stop()
	fmt.Printf("***************************\n")
	os.Exit(0)
}

//Run runs app. Blocks until signal
func (app *App) Run() {
	//1.) Setup signal handlers
	fmt.Printf("***************************\n")
	fmt.Printf("**Starting example_client**\n")
	signaller := event_listener.NewSignaller()
	signaller.HandleSignal(syscall.SIGTERM, app)
	signaller.HandleSignal(syscall.SIGINT, app)
	//Start GRPC server
	var port uint32 = 50052
	app.server.Port = ":" + fmt.Sprint(port)
	app.server.Start()
	//Register interest in events
	registerer := NewRegisterer(port)
	registerer.Register("example_client", "example_subject", ":50051")
	//Block
	for {
		time.Sleep(time.Second)
	}
}
