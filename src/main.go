package main

import (
	"fmt"
	"net"

	authserver "auth/auth"
	"auth/proto"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Starting Auth Server")

	lis, err := net.Listen("tcp", ":9000")

	if err != nil {
		fmt.Println("Error listening:", err.Error())
	}

	dsn := "host=localhost user=postgres password=postgrespw dbname=XodiusAuth port=49153 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database:", err.Error())
	}

	db.Exec("")

	as := authserver.AuthServer{}
	grpcServer := grpc.NewServer()
	proto.RegisterAuthServiceServer(grpcServer, &as)

	if err := grpcServer.Serve(lis); err != nil {
		fmt.Println("Error serving:", err.Error())
	}
}
