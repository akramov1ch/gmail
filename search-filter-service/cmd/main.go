package main

import (
    "log"
    "net"

    "google.golang.org/grpc"
    "search-service/internal/db"
    "search-service/pkg/handlers"
    "search-service/proto"
)

func main() {
    db.InitPostgres()

    lis, err := net.Listen("tcp", ":9004")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    proto.RegisterSearchFilterServiceServer(grpcServer, &handlers.SearchFilterService{})

    log.Println("Search & Filter service is running on port 50054")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
