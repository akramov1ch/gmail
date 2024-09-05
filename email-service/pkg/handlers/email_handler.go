package handlers

import (
	"email-service/internal/db"
	"email-service/internal/kafka"
	"email-service/pkg/models"
	"context"
	"encoding/json"

	"email-service/proto"
)

type EmailService struct {
	proto.UnimplementedEmailServiceServer
}

func (s *EmailService) ComposeEmail(ctx context.Context, req *proto.ComposeEmailRequest) (*proto.ComposeEmailResponse, error) {
	email := models.Email{
		Sender:      req.Sender,
		Recipients:  req.To,
		CC:          req.Cc,
		BCC:         req.Bcc,
		Subject:     req.Subject,
		Body:        req.Body,
		Attachments: req.Attachments,
	}

	emailJson, err := json.Marshal(email)
	if err != nil {
		return nil, err
	}

	err = kafka.SendEmailMessage(string(emailJson))
	if err != nil {
		return nil, err
	}

	_, err = db.DB.Exec(
		"INSERT INTO emails (sender, recipients, cc, bcc, subject, body, attachments) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		email.Sender, email.Recipients, email.CC, email.BCC, email.Subject, email.Body, email.Attachments,
	)
	if err != nil {
		return nil, err
	}

	return &proto.ComposeEmailResponse{MessageId: "some-id", Success: true}, nil
}

func (s *EmailService) GetInbox(ctx context.Context, req *proto.GetInboxRequest) (*proto.GetInboxResponse, error) {
	rows, err := db.DB.Query("SELECT id, sender, subject, body, attachments, created_at FROM emails WHERE recipients @> ARRAY[$1]", req.UserEmail)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var emails []*proto.EmailMessage
	for rows.Next() {
		email := proto.EmailMessage{}
		var attachments []string
		err := rows.Scan(&email.Id, &email.Sender, &email.Subject, &email.Body, &attachments, &email.CreatedAt)
		if err != nil {
			return nil, err
		}
		emails = append(emails, &email)
	}

	return &proto.GetInboxResponse{Emails: emails}, nil
}
