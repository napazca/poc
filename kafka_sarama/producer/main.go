package main

import (
	"log"
	"strings"
	"time"

	"github.com/Shopify/sarama"
)

func main() {
	brokers := strings.Split("localhost:9092", ",")
	saramaConfig := sarama.NewConfig()
	saramaConfig.ClientID = "biller_orchestrator_producer"
	saramaConfig.Producer.RequiredAcks = sarama.WaitForAll
	saramaConfig.Producer.Retry.Max = 3
	saramaConfig.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, saramaConfig)
	if err != nil {
		log.Fatalln(err)
	}

	// sync producer is used because it handles transaction that worth to be waited

	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	for {
		msg := &sarama.ProducerMessage{Topic: "my_topic", Value: sarama.StringEncoder("testing 123")}
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Printf("FAILED to send message: %s\n", err)
		} else {
			log.Printf("> message sent to partition %d at offset %d\n", partition, offset)
		}

		time.Sleep(1 * time.Second)
	}

}
