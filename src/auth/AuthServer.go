package authserver

import (
	"auth/proto"
	"context"
	"fmt"
)

type AuthServer struct {
	proto.UnimplementedAuthServiceServer
}

func (s *AuthServer) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	fmt.Println("Login request received")
	return &proto.LoginResponse{
		Token: "12345",
		Status: proto.ResponseStatus_OK,
	}, nil
}

func (s *AuthServer) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	fmt.Println("Register request received")
	return &proto.RegisterResponse{
		Status: proto.ResponseStatus_OK,
	}, nil
}