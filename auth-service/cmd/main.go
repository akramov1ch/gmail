package main

import (
	"log"
	"net"

	"auth-service/internal/db"
	"auth-service/internal/redis"

	"google.golang.org/grpc"
)

func main() {
	db.InitPostgres()
	redis.InitRedis()

	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	log.Println("Auth Service is running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
