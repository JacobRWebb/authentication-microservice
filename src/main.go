package main

import (
	"fmt"
	"net"

	Auth "github.com/JacobRWebb/authentication-microservice/auth"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting Auth Server")

	lis, err := net.Listen("tcp", ":9000")

	if err != nil {
		fmt.Println("Error listening:", err.Error())
	}

	grpcServer := grpc.NewServer()
	Auth.AttachServer(grpcServer)

	grpcServer.Serve(lis)

	if err := grpcServer.Serve(lis); err != nil {
		fmt.Println("Error serving:", err.Error())
	}
}