package main

import (
	"fmt"
	"net"

	Auth "github.com/JacobRWebb/authentication-microservice/auth"
	"github.com/JacobRWebb/authentication-microservice/database"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting Auth Server")

	lis, err := net.Listen("tcp", ":9000")

	if err != nil {
		fmt.Println("Error listening:", err.Error())
	}

	database.DB = database.CreateDatabaseConnection()
	
	grpcServer := grpc.NewServer()
	Auth.AttachServer(grpcServer)

	database.CreateDatabaseConnection()

	if err := grpcServer.Serve(lis); err != nil {
		fmt.Println("Error serving:", err.Error())
	}
}