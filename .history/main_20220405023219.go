package main

import (
	"fmt"
	"log"
	"net"

	"github.com/espitman/mygrpc-userService/user"
	"google.golang.org/grpc"
)

var PORT = "3300"

func main() {
	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := user.Server{}
	grpcServer := grpc.NewServer()

	user.RegisterUserServiceServer(grpcServer, &s)

	fmt.Println("Port: " + PORT)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}

}
