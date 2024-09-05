package main

import (
	"log"
	"net"

	"notification-service/internal/kafka"
	"notification-service/pkg/handlers"
	"notification-service/proto"

	"google.golang.org/grpc"
)

func main() {
	kafka.InitKafka()
	defer kafka.Close()

	lis, err := net.Listen("tcp", ":9006")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterNotificationServiceServer(grpcServer, &handlers.NotificationService{})

	log.Println("Notification service is running on port 50056")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
