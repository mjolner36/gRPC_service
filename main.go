package main

import (
	"google.golang.org/grpc"
	"log"
	"mjolner36/gRPC_service/pb"
	"mjolner36/gRPC_service/server"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, server.NewAuthService())

	log.Println("gRPC server running at :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
