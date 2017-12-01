package main

import (
	"fmt"
	"os"
	"os/signal"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/Shopify/sarama"
)

var (
	brokerList        = kingpin.Flag("brokerList", "List of brokers to connect").Default("localhost:9092").Strings()
	topic             = kingpin.Flag("topic", "Topic name").Default("important").String()
	partition         = kingpin.Flag("partition", "Partition number").Default("0").String()
	offsetType        = kingpin.Flag("offsetType", "Offset Type (OffsetNewest | OffsetOldest)").Default("-1").Int()
	messageCountStart = kingpin.Flag("messageCountStart", "Message counter start from:").Int()
)

func main() {
    // parse the arguements
	kingpin.Parse()

    // start using the sarama API
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
    // set up which brokers are goingto be used
	brokers := *brokerList
    // make a new consumer using the config and check for errors
	master, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		panic(err)
	}

    // close the consumer when it closes so that kafka knows it is gone and let
    // the user know if there is an error
	defer func() {
		if err := master.Close(); err != nil {
			panic(err)
		}
	}()

    // start consuming things!
	consumer, err := master.ConsumePartition(*topic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	doneCh := make(chan struct{})

    // print out what we consume
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				*messageCountStart++
				fmt.Println("Received messages", string(msg.Key), string(msg.Value))
			case <-signals:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()
	<-doneCh

    // let the user know how many messages we consumed
	fmt.Println("Processed", *messageCountStart, "messages")
}
