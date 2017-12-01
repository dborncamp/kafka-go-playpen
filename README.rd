# Go Kafka test
This is to test producing and consuming messages in kafka using golang.
Right now this only works in standalone mode because I wanted to start with
a simple example to start with. This is heavily cribbed from
[vsouza/go-kafka-example](https://github.com/vsouza/go-kafka-example) but this
will read user input from the command line and put it through kafka. It 
basically just works with string passing for now. A lot of this README is mostly
notes for myself.


A good quick start guide to get things up and running: [https://kafka.apache.org/quickstart](https://kafka.apache.org/quickstart)


## To start

### In one terminal

#### zookeeper
Start zookeeper. All of the config defaults are fine, but if you change the port
you will need to change the code. Use `zkServer start zoo.cfg` to start zookeeper.


#### kafka
Again, here the defaults are fine for Kafka. Using the included `server.properties`
file should make things easy. Use `kafka-server-start server.properties`


### In another terminal

#### Start the consumer

We should make sure that we have all of the dependicies so go into the consumer
directory and get everything. Then start the consumer. Both zookeeper and Kafka
need to be going for the consumer to start properly

    cd consumer
    go get
    go run main.go

This should not print out anything yet but be an empty terminal until the
producer is started.


### In yet another terminal

#### Start the producer

Again, make sure the dependiceies are there then run everything.

    cd producer
    go get
    go run main.go

This terminal will prompt you for user input which should appear in the
consumer terminal window.

Now you should see things happening in all three terminals, the producer you
type in, the consumer reporting what you typed, and the kafka server routing
everything.
