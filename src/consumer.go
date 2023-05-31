package main

import (
	"context"
	"log"

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

	// Subscribe to a topic with a given subscription name
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            "my-topic",        // Topic to consume from
		SubscriptionName: "my-subscription", // Name of the subscription
		Type:             pulsar.Shared,     // Subscription type (e.g., Shared)
	})
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}
	defer consumer.Close() // Ensure the consumer closes after receiving

	// Receive a message from the topic
	msg, err := consumer.Receive(context.Background())
	if err != nil {
		log.Fatalf("Failed to receive message: %v", err)
	}

	log.Printf("Received message msgId: %v -- content: '%s'", msg.ID(), string(msg.Payload())) // Log received message

	// Acknowledge the message to tell Pulsar that it has been successfully processed
	consumer.Ack(msg)
}
