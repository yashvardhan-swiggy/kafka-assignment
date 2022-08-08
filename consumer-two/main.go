package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {

	topic := "example-123"

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		GroupID: "group-2",
		Topic:   topic,
	})

	for {
		message, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}

		fmt.Printf("message at offset %d: %s = %s\n", message.Offset, string(message.Key), string(message.Value))
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
