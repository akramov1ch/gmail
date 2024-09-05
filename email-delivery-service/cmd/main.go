package main

import (
	"log"
	"net"

	"email-delivery-service/pkg/handlers"
	"email-delivery-service/proto"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9007")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterEmailDeliveryServiceServer(grpcServer, &handlers.EmailDeliveryService{})

	log.Println("Email Delivery service is running on port 50057")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
