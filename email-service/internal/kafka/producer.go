package kafka

import (
    "log"
    "github.com/segmentio/kafka-go"
    "context"
)

var writer *kafka.Writer

func InitKafkaProducer() {
    writer = &kafka.Writer{
        Addr:     kafka.TCP("localhost:9092"),
        Topic:    "email-topic",
        Balancer: &kafka.LeastBytes{},
    }
}

func SendEmailMessage(message string) error {
    err := writer.WriteMessages(context.Background(),
        kafka.Message{
            Value: []byte(message),
        },
    )
    if err != nil {
        log.Printf("Failed to send message to Kafka: %v", err)
        return err
    }

    log.Println("Message sent to Kafka")
    return nil
}
