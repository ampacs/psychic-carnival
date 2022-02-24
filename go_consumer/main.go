package main

import (
	"fmt"
	"os"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func main() {
	// Get the broker address, defaults to kafka:9092
	// If you want to
	broker := os.Getenv("KAFKA_BROKER")
	if broker == "" {
		broker = "kafka:9092"
	}

	// Connect consumer
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id":          "go_consumer",
		"auto.offset.reset": "latest",
	})

	if err != nil {
		panic(err)
	}

	// Subscribe to the picture topic
	c.SubscribeTopics([]string{"picture"}, nil)
	ConsumeMessages(c)

	c.Close()
}

// ConsumeMessages indefinitely consumes messages from the picture topic
func ConsumeMessages(c *kafka.Consumer) {
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			HandleMessage(msg)
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}

// HandleMessage handles a single message consumed from the picture topic
func HandleMessage(message *kafka.Message) {
	topic := *message.TopicPartition.Topic
	fmt.Printf("[GoLang] Received a message on topic %s: %s\n", topic, string(message.Value))
}
