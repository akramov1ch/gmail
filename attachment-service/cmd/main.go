package main

import (
	"log"
	"net"

	"attachment-service/internal/db"
	"attachment-service/internal/storage"

	"google.golang.org/grpc"
)

func main() {
	db.InitPostgres()
	storage.InitS3()

	lis, err := net.Listen("tcp", ":9003")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	log.Println("Attachment service is running on port 50053")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
