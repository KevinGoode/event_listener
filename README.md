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
1.) event_listener detects that a client has gone down while trying to send message. Message is added to buffer
and every 5 seconds a connection is re-attempted. If a connection is re-established then all proxies associated with this client 
are repaired. Messages are added to the buffer until the buffer reaches TBD MB when proxy is deleted.
2.) event_listener detects nats server has gone down. In this scenario, the event listener will attempt reconnection every 10 seconds
for a time of 1 hour. After this time the evetn listener will terminate

3.) Registration request for same message type from same client over same port.  Code checks integrity of return connection with client.
If connection still ok then use existing proxy and return error 'already listening for message on this port'. 
If connection has been terminated, kill all proxies for this client and port and re-create all proxies for this client.
4.) Registration request for same message from same client over different port.  Code checks integrity of return connection with client.
If connection still ok then return error 'already listening for message on port XX'....
If connection is down , delete all proxies relating to this client and port and create new proxy for this new message.



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
5.) Run the publisher and send a message as  follows: ./temp example_subject "MSG1"
(Message should be printed  to example_client console)


