package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

func main() {
	// Create a Pulsar client instance
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: "pulsar://localhost:6650", // Pulsar broker URL
	})
	if err != nil {
		log.Fatalf("Failed to create Pulsar client: %v", err)
	}
	defer client.Close() // Ensure the client closes after the program finishes

	// Create a producer to send messages to a specific topic
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: "my-topic", // The topic where messages will be published
	})
	if err != nil {
		log.Fatalf("Failed to create producer: %v", err)
	}
	defer producer.Close() // Ensure the producer closes after sending

	// Message count variable
	messageCount := 10 // Number of messages to send

	// Loop to send multiple numbered messages with timestamps
	for i := 1; i <= messageCount; i++ {
		// Get the current time
		currentTime := time.Now().Format(time.RFC3339) // Time in RFC3339 format

		// Message payload with numbering and timestamp
		messagePayload := fmt.Sprintf("Message %d sent at %s", i, currentTime)

		// Send the message
		msgID, err := producer.Send(context.Background(), &pulsar.ProducerMessage{
			Payload: []byte(messagePayload), // Message content
		})
		if err != nil {
			log.Fatalf("Failed to send message: %v", err)
		}

		log.Printf("Message %d sent successfully at %s with ID: %v", i, currentTime, msgID) // Log the message ID, time, and number
	}
}
