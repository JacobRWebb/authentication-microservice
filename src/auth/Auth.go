package Auth

import (
	"context"
	"fmt"

	"github.com/JacobRWebb/authentication-microservice/pb/auth"
)

type AuthServer struct {
	auth.UnimplementedAuthServiceServer
}

func (s *AuthServer) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	fmt.Println("Login request received")
	return &auth.LoginResponse{
		Token: "12345",
		Status: auth.ResponseStatus_OK,
	}, nil
}

func (s *AuthServer) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	fmt.Println("Register request received")
	return &auth.RegisterResponse{
		Status: auth.ResponseStatus_OK,
	}, nil
}