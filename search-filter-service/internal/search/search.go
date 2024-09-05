package search

import (
    "search-service/internal/db"
    "search-service/pkg/models"
    "log"
)

func SearchEmails(userId, query string, labels []string) ([]models.Email, error) {
    var emails []models.Email
    rows, err := db.DB.Query(`
        SELECT email_id, sender, subject, content, date_sent
        FROM emails
        WHERE user_id = $1 AND (subject ILIKE '%' || $2 || '%' OR sender ILIKE '%' || $2 || '%' OR content ILIKE '%' || $2 || '%')`,
        userId, query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var email models.Email
        if err := rows.Scan(&email.EmailId, &email.Sender, &email.Subject, &email.Preview, &email.DateSent); err != nil {
            log.Printf("Error scanning row: %v", err)
            continue
        }
        emails = append(emails, email)
    }

    return emails, nil
}
