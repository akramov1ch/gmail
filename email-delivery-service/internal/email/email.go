package email

import (
    "email-delivery-service/internal/smtp"
    "log"
)

func SendEmail(to, subject, body string, cc, bcc []string, attachment []byte) {
    err := smtp.SendEmail(to, subject, body, cc, bcc, attachment)
    if err != nil {
        log.Printf("Failed to send email: %v", err)
    }
}
