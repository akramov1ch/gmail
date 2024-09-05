package filter

import (
    "search-service/internal/db"
    "search-service/pkg/models"
)

func CreateFilter(userId, sender, subjectContains, folder string) (string, error) {
    var filterId string
    err := db.DB.QueryRow(`
        INSERT INTO filters (user_id, sender, subject_contains, move_to_folder) 
        VALUES ($1, $2, $3, $4) 
        RETURNING id`, userId, sender, subjectContains, folder).Scan(&filterId)
    if err != nil {
        return "", err
    }
    return filterId, nil
}

func ApplyFilters(userId string) ([]models.FilteredEmail, error) {
    var filteredEmails []models.FilteredEmail
    rows, err := db.DB.Query(`
        SELECT emails.id, filters.move_to_folder
        FROM emails
        JOIN filters ON emails.user_id = filters.user_id
        WHERE filters.user_id = $1 AND 
              (emails.sender ILIKE '%' || filters.sender || '%' 
               OR emails.subject ILIKE '%' || filters.subject_contains || '%')`,
        userId)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var filteredEmail models.FilteredEmail
        if err := rows.Scan(&filteredEmail.EmailId, &filteredEmail.Folder); err != nil {
            return nil, err
        }
        filteredEmails = append(filteredEmails, filteredEmail)
    }

    return filteredEmails, nil
}
