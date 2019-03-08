//What is this code?
-----------------
The event_listener is a proxy process that fowards on NATS messages
to interested applications over GRPC.

The event_listener listens for registration messages over GRPC.
When it receives a registration request it starts a new proxy component that
subscribes to the message type with the NATS server and sits waiting for messages of this type to be received.
When it receives a message, it forwards it on to interested parties. Because clients keep a record of last message they received
and pass this as argument when registering, no messages will ever be lost!

Supported +ve use cases:
1.) Different message types from same example_client
2.) Same message type from 2 or more clients
3.) Multiple clients requesting multiple message types


Supported -ve use cases:
1.) event_listener detects nats server has gone down. In this scenario, the event listener will attempt reconnection every 10 seconds for a time of 1 hour. After this time the event listener will terminate. When reconnection is successful
all proxies are re-subscribed 
2.) Registration request for same message type from same client.  Old proxy is stopped and deleted and new one created. 
(This use case typically happens after client crash, in which case the client will probably start with a higher sequence
 number so best to delete old.) 
3.) There ia a time delay between client registering for events and listening for them. (Probably should listen first then regsiter but anyway....) In this scenario event_listener stores up failed message forwards in a buffer and finally when the client listens it forwards all messages in buffer on. To test this scenario :
./example_client -c example_clientxx -s example_subject -w
Then send a few messages
4.) Clients determine whether event_listener has crashed by calling GetEventHandlerStatus. If event listener fails then example client re-registers

Install GRPC
https://grpc.io/docs/quickstart/go.html

(NB Protoc zip need this: yum install libatomic)


Generate go grpc apis

protoc   --go_out=plugins=grpc:./ ./event_registry.proto
protoc   --go_out=plugins=grpc:./ ./event_generic.proto


//API Guide is at
https://nats.io/documentation/streaming/nats-streaming-quickstart/

//Example https://github.com/nats-io/go-nats-streaming/tree/master/examples

//How do I run this code?
1.) Download nats_streaming server from here and run it (no command line options)
https://github.com/nats-io/nats-streaming-server/releases (linuxamd64)
2.) Download publisher code into temp dir and build. Code is  here :
https://github.com/nats-io/go-nats-streaming/blob/master/examples/stan-pub/main.go
3.) Start the main event_listener code (./event_listener)
4.) Start the example_client code (./example_client) which registers with event event_listener
and asks to be sent messages on subject 'example_subject'
eg2 start client code registering interest in 2 messages
./example_client -c example_clientxx -s example_subject,example_subject2
5.) Run the publisher and send a message as  follows: ./temp example_subject "MSG1"
(Message should be printed  to example_client console)


