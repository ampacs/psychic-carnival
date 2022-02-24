package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"

	"github.com/ampacs/psychic-carnival/go_consumer/api"
	"github.com/ampacs/psychic-carnival/go_consumer/message"
)

func main() {
	// Get the broker address, defaults to kafka:9092
	// If you want to
	broker := os.Getenv("KAFKA_BROKER")
	if broker == "" {
		broker = "kafka:9092"
	}

	// Connect consumer
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id":          "go_consumer",
		"auto.offset.reset": "latest",
	})
	if err != nil {
		panic(fmt.Errorf("consumer creation error: %v", err))
	}
	defer func() {
		if err = consumer.Close(); err != nil {
			panic(fmt.Errorf("consumer closing error: %v", err))
		}
	}()

	// Subscribe to the picture topic
	if err = consumer.SubscribeTopics([]string{"picture"}, nil); err != nil {
		panic(fmt.Errorf("topic subscription error: %v", err))
	}

	// create a buffered channel to help prevent the message consumer from blocking
	messageBuffer := make(chan message.Message, 10)
	messageRepository := message.Repository{}

	// create a message register that will update the repository
	messageRegister := message.NewRegister(&messageRepository)
	// messages are obtained from the channel
	messageRegister.RegisterMessages(messageBuffer)

	// create a message consumer
	messageConsumer := message.NewConsumer(consumer)
	// messages are sent through the buffer
	messageConsumer.ConsumeMessages(messageBuffer)

	// create the API with access to the repository
	apiHandler := api.NewApi(&messageRepository)
	router := mux.NewRouter().StrictSlash(true)
	// register the routes
	router.HandleFunc("/count", apiHandler.Count)
	router.HandleFunc("/average-weight", apiHandler.AverageWeight)

	// handle CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})

	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(":80", handler))
}
