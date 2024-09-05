package notifier

import (
    "notification-service/internal/kafka"
    "log"
)

func Notify(userId, message string) {
    err := kafka.SendMessage(userId, message)
    if err != nil {
        log.Printf("Failed to send notification: %v", err)
    }
}
