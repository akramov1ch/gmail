package contacts

import (
	"contact-management-service/internal/db"
	"contact-management-service/pkg/models"
)

func AddContact(userId, name, email, phoneNumber string) (string, error) {
	var contactId string
	err := db.DB.QueryRow(`
        INSERT INTO contacts (user_id, name, email, phone_number) 
        VALUES ($1, $2, $3, $4) 
        RETURNING id`, userId, name, email, phoneNumber).Scan(&contactId)
	if err != nil {
		return "", err
	}
	return contactId, nil
}

func UpdateContact(contactId, name, email, phoneNumber string) error {
	_, err := db.DB.Exec(`
        UPDATE contacts 
        SET name = $1, email = $2, phone_number = $3
        WHERE id = $4`, name, email, phoneNumber, contactId)
	return err
}

func DeleteContact(contactId string) error {
	_, err := db.DB.Exec(`DELETE FROM contacts WHERE id = $1`, contactId)
	return err
}

func GetContacts(userId string) ([]models.Contact, error) {
	var contacts []models.Contact
	rows, err := db.DB.Query(`SELECT id, name, email, phone_number FROM contacts WHERE user_id = $1`, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var contact models.Contact
		if err := rows.Scan(&contact.ContactId, &contact.Name, &contact.Email, &contact.PhoneNumber); err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
	}
	return contacts, nil
}

func SearchContacts(userId, query string) ([]models.Contact, error) {
	var contacts []models.Contact
	rows, err := db.DB.Query("SELECT id, name, email, phone_number FROM contacts WHERE user_id = $1 AND (name ILIKE '%' || $2 || '%' OR email ILIKE '%' || $2 || '%'", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var contact models.Contact
		if err := rows.Scan(&contact.ContactId, &contact.Name, &contact.Email, &contact.PhoneNumber); err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
	}
	return contacts, nil
}
