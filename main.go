package main

import event_listener "event_listener/listener"

func main() {
	app := event_listener.NewApp()
	app.Run()
}
