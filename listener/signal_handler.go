package event_listener
import (
       "os"
       "os/signal"
)
type SignalHandler interface {
	HandleSignal(os.Signal)
}
type Signaller struct {}
func (signalHandler *Signaller) HandleSignal(sig os.Signal, handler SignalHandler) () {
	var waitSignal = make(chan os.Signal)
	signal.Notify(waitSignal, sig)
	go func() {
		//Block until signal received
		sig := <-waitSignal
		//Now call signal handler
		handler.HandleSignal(sig)
		
    }()
}
func NewSignaller()(*Signaller){
	signaller := Signaller{}
	return &signaller
}