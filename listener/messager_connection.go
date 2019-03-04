package event_listener

import (
	fmt "fmt"
	"time"

	stan "github.com/nats-io/go-nats-streaming"
)

//MessagerConnection is container for NATS connection parameters
type MessagerConnection struct {
	connection stan.Conn
	ClientID   string
	listeners  []ListenerAPI
}

//MessagerConnectionAPI exposes MessageConnection services
type MessagerConnectionAPI interface {
	Connect() error
	Disconnect()
	GetConnection() stan.Conn
	AddListener(listener ListenerAPI)
}

//Connect connects toNATs messager
func (conn *MessagerConnection) Connect() error {
	fmt.Printf("Connecting to message server...\n")
	var err error
	var clusterID = "test-cluster"
	conn.connection, err = stan.Connect(clusterID, conn.ClientID, stan.NatsURL(stan.DefaultNatsURL),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			fmt.Printf("Connection lost, reason: %v\n", reason)
			conn.connectionLossHandler(reason)
		}))
	if err == nil {
		fmt.Printf("Successfully connected to message server\n")
	} else {
		fmt.Printf("Failed to connect to message server\n")
	}
	return err
}

//Disconnect disconnects from  NATs server
func (conn *MessagerConnection) Disconnect() {
	if conn.connection != nil {
		fmt.Printf("Closing connection...\n")
		conn.connection.Close()
	}
}

//GetConnection is an accessor used to return connection
func (conn *MessagerConnection) GetConnection() stan.Conn {
	return conn.connection
}

//AddListener adds a listener reference to the connection so in the event the
//connection is lost, asscaoited listeners can be re-initialised (IE subscribe called again)
func (conn *MessagerConnection) AddListener(listener ListenerAPI) {
	conn.listeners = append(conn.listeners, listener)
}
func (conn *MessagerConnection) connectionLossHandler(reason error) {
	fmt.Printf("Disconnected from message server, reason: %v\n", reason)
	fmt.Printf("Trying to reconnect......\n")
	err := conn.reconnect(reason)
	if err == nil {
		//Successfully reconnected so now re-subscribe on all listeners
		for _, listener := range conn.listeners {
			listener.Subscribe()
		}
	}
}
func (conn *MessagerConnection) reconnect(reason error) error {
	var maxCount uint = 100
	var count uint
	var err error
	for err = reason; err != nil; err = conn.Connect() {
		count++
		fmt.Printf("Attempting to re-connect to message server...\n")
		if count > maxCount {
			fmt.Printf("Cannot reconnect after 10 mins, reason: %v\n", err)
			break
		}
		time.Sleep(6 * time.Second)
	}
	return err
}

// CreateMessagerConnection creates and initialises a MessagerConnection or returns existing on if it exists already
func CreateMessagerConnection(clientID string, listener ListenerAPI) MessagerConnectionAPI {
	fmt.Printf("Creating new connection to messenger\n")
	//Do we have connection already?
	cache := GetMessagerConnectionCacheInstance()
	connectionAPI := cache.GetConnection(clientID)
	if connectionAPI == nil {
		//No connction so create  one.
		connection := MessagerConnection{}
		connection.ClientID = clientID
		//Now connect.
		connectionAPI = &connection
		connectionAPI.Connect()
		//Cache connection
		cache.AddNewConnection(clientID, connectionAPI)
	}
	//Register listener with connection so when connection is lost , on reconnection can re-subscribe
	connectionAPI.AddListener(listener)
	return connectionAPI
}
