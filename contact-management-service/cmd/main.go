package main

import (
    "log"
    "net"

    "google.golang.org/grpc"
    "contact-management-service/internal/db"
    "contact-management-service/pkg/handlers"
    "contact-management-service/proto"
)

func main() {
    db.InitPostgres()

    lis, err := net.Listen("tcp", ":9005")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    proto.RegisterContactManagementServiceServer(grpcServer, &handlers.ContactManagementService{})

    log.Println("Contact Management service is running on port 50055")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
