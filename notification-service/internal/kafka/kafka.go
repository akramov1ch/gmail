package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

var writer *kafka.Writer

func InitKafka() {
	writer = kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "notifications",
		Balancer: &kafka.LeastBytes{},
	})
}

func SendMessage(userId, message string) error {
	err := writer.WriteMessages(
		context.Background(),
		kafka.Message{
			Key:   []byte(userId),
			Value: []byte(message),
		},
	)
	return err
}

func Close() {
	if err := writer.Close(); err != nil {
		log.Fatalf("Failed to close Kafka writer: %v", err)
	}
}
