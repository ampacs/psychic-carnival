package message

import (
	"encoding/json"
	"fmt"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

// Consumer represents a kafka message consumer
type Consumer struct {
	consumer *kafka.Consumer
}

// NewConsumer returns a Consumer that will read messages from the given kafka consumer
func NewConsumer(c *kafka.Consumer) Consumer {
	return Consumer{
		consumer: c,
	}
}

// ConsumeMessages handles reading messages from the kafka consumer and sends them through the given channel
func (h *Consumer) ConsumeMessages(mc chan<- Message) {
	go func() {
		for {
			msg, err := h.consumer.ReadMessage(-1)
			if err != nil {
				// The client will automatically try to recover from all errors.
				fmt.Printf("Consumer error: %v (%v)\n", err, msg)
				continue
			}

			if err = h.handleMessage(msg, mc); err != nil {
				fmt.Printf("Message handling error: %v (%v)\n", err, msg)
			}
		}
	}()
}

// handleMessage handles a single consumed message and sends it through the channel
func (h Consumer) handleMessage(message *kafka.Message, mc chan<- Message) error {
	var m Message
	if err := json.Unmarshal(message.Value, &m); err != nil {
		return err
	}

	topic := *message.TopicPartition.Topic
	fmt.Printf("[GoLang] Received a message on topic %s: %s\n", topic, string(message.Value))

	mc <- m

	return nil
}
