package main

import (
	event_listener "event_listener/listener"
	"flag"
	"fmt"
	"os"
	"strings"
	"syscall"
	"time"
)

var usageStr = `
Usage: example_client [options]
Options:
	-s subjects (comma separated subjects. Eg example_subject) 
	-c ClientId (example_client)
	-p Port (50052)
`

// NOTE: Use tls scheme for TLS, e.g. stan-pub -s tls://demo.nats.io:4443 foo hello
func usage() {
	fmt.Printf("%s\n", usageStr)
	os.Exit(0)
}

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
	var subjects = ""
	var client = ""
	var port uint = 50052
	flag.StringVar(&subjects, "s", "example_subject", "Subjects")
	flag.StringVar(&client, "c", "example_client", "Client")
	flag.UintVar(&port, "p", 50052, "Port")

	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) > 3 {
		usage()
	}
	messageSubjects := strings.Split(subjects, ",")
	fmt.Printf("Client Id is %s. Subjects are: %s \n", client, subjects)
	//Start GRPC server
	app.server.Port = ":" + fmt.Sprint(port)
	app.server.Start()
	//Register interest in events
	registerer := NewRegisterer(uint32(port))
	for _, sub := range messageSubjects {
		fmt.Printf("Registering interest in subject %s...\n", sub)
		registerer.Register(client, sub, ":50051")
	}
	//Block
	for {
		time.Sleep(time.Second)
	}
}
