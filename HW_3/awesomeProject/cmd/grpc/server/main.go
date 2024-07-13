package main

import (
	"log"
	"net"

	"awesomeProject/accounts/grpc/handler"

	pb "awesomeProject/proto"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":4567")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	handler := handler.New()

	s := grpc.NewServer()
	pb.RegisterAccountsServer(s, handler)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
