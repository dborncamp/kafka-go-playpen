package main

import (
	"fmt"
    "bufio"
    "os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/Shopify/sarama"
)

var (
	brokerList = kingpin.Flag("brokerList", "List of brokers to connect").Default("localhost:9092").Strings()
	topic      = kingpin.Flag("topic", "Topic name").Default("important").String()
)


func main() {
    // parse the arguements to send kafka things to the right place
	kingpin.Parse()

    // Create a scanner to get user input
    scanner := bufio.NewScanner(os.Stdin)

    // Let the user know what is going on
    fmt.Println("Send a message to another terminal using kafka and zookeeper")
    fmt.Println("Type q to stop.")
    for scanner.Scan() {
        // scan for the user input
        userResponse := scanner.Text()

        // make sure there was not a problem...
        if scanner.Err() != nil {
            fmt.Println("There was a problem scanning")
            panic(scanner.Err())
        }

        // make it escapable
        if (userResponse == "q"){
            break
        }

        // let the user know what they wrote
        fmt.Print("User input: ")
        fmt.Println(userResponse)

        // give it to the producer to stick in kafka
        write(userResponse)
    }
}

// write writes the string as the kafka producer.
func write(s string) {
    // start up the go API for kafka and get it going
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(*brokerList, config)

    // check if there is an error
	if err != nil {
		panic(err)
	}

    // closethe producer when it is done so it isn't hanging around to mess up kafka
	defer func() {
		if err := producer.Close(); err != nil {
			panic(err)
		}
	}()

    // build up the producer message to put in kafka
	msg := &sarama.ProducerMessage{
		Topic: *topic,
		Value: sarama.StringEncoder(s),
	}

    // send the message and check for errors
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		panic(err)
	}

    // let the user know it was sent
	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", *topic, partition, offset)
}
