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

var EVENT_LISTENER_PORT = ":50051"
var usageStr = `
Usage: example_client [options]
Options:
	-s subjects (comma separated subjects. Eg example_subject) 
	-c ClientId (example_client)
	-p Port (50052)
	-w wait (2mins) before listening for evetns. (This tests evetn_listener message buffer)
`

// NOTE: Use tls scheme for TLS, e.g. stan-pub -s tls://demo.nats.io:4443 foo hello
func usage() {
	fmt.Printf("%s\n", usageStr)
	os.Exit(0)
}

//App is entry point for app
type App struct {
	registerer      *Registerer
	server          Server
	messageSubjects []string
	client          string
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
	app.unRegisterAll(EVENT_LISTENER_PORT)
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
	var port uint = 50052
	var wait = false
	flag.StringVar(&subjects, "s", "example_subject", "Subjects")
	flag.StringVar(&app.client, "c", "example_client", "Client")
	flag.UintVar(&port, "p", 50052, "Port")
	flag.BoolVar(&wait, "w", false, "Slow starter") //Set this flag to simulate slow startup

	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) > 3 {
		usage()
	}
	app.messageSubjects = strings.Split(subjects, ",")
	fmt.Printf("Client Id is %s. Subjects are: %s \n", app.client, subjects)
	//Start GRPC server
	app.server.Port = ":" + fmt.Sprint(port)
	//Start server in background
	go app.StartServer(wait)
	//Register interest in events
	app.registerer = NewRegisterer(uint32(port))
	err := app.registerAll(EVENT_LISTENER_PORT)
	if err == nil {
		fmt.Printf("Registration completed ok\n")
	} else {
		fmt.Printf("Event listener down\n")
		os.Exit(-1)
	}
	//Block
	for {
		time.Sleep(time.Second * 5)
		err = app.checkAll()
		if err != nil {
			fmt.Printf("Event listener gone down\n")
			//Blocking call
			app.reRegister(EVENT_LISTENER_PORT)
		}
	}
}

//StartServer Starts server
func (app *App) StartServer(wait bool) {
	if wait {
		//Simulate slow start up
		//Wait two minutes before listening for GRPC. Event listener should queue up messages
		time.Sleep(time.Minute * 2)
	}
	app.server.Start()
}
func (app *App) checkAll() error {
	var err error
	for _, sub := range app.messageSubjects {
		err = app.registerer.CheckStillRegistered(app.client, sub)
		if err != nil {
			break
		}
	}
	return err
}
func (app *App) reRegister(targetPort string) error {
	var err error
	for {
		time.Sleep(time.Second * 5)
		err = app.registerAll(targetPort)
		if err == nil {
			break
		}
	}
	return err
}
func (app *App) registerAll(targetPort string) error {
	err := app.registerer.Connect(targetPort)
	if err == nil {
		for _, sub := range app.messageSubjects {
			fmt.Printf("Registering interest in subject %s...\n", sub)
			err = app.registerer.Register(app.client, sub)
			if err != nil {
				break
			}
		}
	}
	return err
}
func (app *App) unRegisterAll(targetPort string) error {
	err := app.registerer.Connect(targetPort)
	if err == nil {
		for _, sub := range app.messageSubjects {
			fmt.Printf("Un registering interest in subject %s...\n", sub)
			err = app.registerer.Unregister(app.client, sub)
			if err != nil {
				break
			}
		}
	}
	return err
}
