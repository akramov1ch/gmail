package main

import (
    "log"
    "net"

    "google.golang.org/grpc"
    "email-service/internal/db"
    "email-service/internal/kafka"
    "email-service/pkg/handlers"
    "email-service/proto"
)

func main() {
    db.InitPostgres()
    kafka.InitKafkaProducer()

    lis, err := net.Listen("tcp", ":9002")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    proto.RegisterEmailServiceServer(grpcServer, &handlers.EmailService{})

    log.Println("Email service is running on port 50052")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
