package handlers

import (
    "context"
    "notification-service/internal/notifier"
    "notification-service/proto"
)

type NotificationService struct {
    proto.UnimplementedNotificationServiceServer
}

func (s *NotificationService) SendNotification(ctx context.Context, req *proto.SendNotificationRequest) (*proto.SendNotificationResponse, error) {
    notifier.Notify(req.UserId, req.Message)
    return &proto.SendNotificationResponse{Status: "Notification Sent"}, nil
}
