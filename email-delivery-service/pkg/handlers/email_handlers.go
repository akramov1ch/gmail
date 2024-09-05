package handlers

import (
    "context"
    "email-delivery-service/internal/email"
    "email-delivery-service/proto"
)

type EmailDeliveryService struct {
    proto.UnimplementedEmailDeliveryServiceServer
}

func (s *EmailDeliveryService) SendEmail(ctx context.Context, req *proto.SendEmailRequest) (*proto.SendEmailResponse, error) {
    email.SendEmail(req.To, req.Subject, req.Body, req.Cc, req.Bcc, req.Attachment)
    return &proto.SendEmailResponse{Status: "Email Sent"}, nil
}
