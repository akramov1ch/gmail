package smtp

import (
    "net/smtp"
    "log"
)

func SendEmail(to, subject, body string, cc, bcc []string, attachment []byte) error {
    from := "shaxbozakramovic@gmail.com"
    password := "vakhaboff2502"
    
    smtpServer := "smtp.gmail.com"
    port := "587"

    auth := smtp.PlainAuth("", from, password, smtpServer)

    msg := []byte("To: " + to + "\r\n" +
        "Subject: " + subject + "\r\n" +
        "\r\n" +
        body + "\r\n")

    err := smtp.SendMail(smtpServer+":"+port, auth, from, append([]string{to}, append(cc, bcc...)...), msg)
    if err != nil {
        log.Printf("Failed to send email: %v", err)
        return err
    }

    return nil
}
