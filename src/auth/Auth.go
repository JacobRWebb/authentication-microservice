package Auth

import (
	"context"
	"fmt"

	"github.com/JacobRWebb/authentication-microservice/database"
	"github.com/JacobRWebb/authentication-microservice/database/models"
	"github.com/JacobRWebb/authentication-microservice/pb/auth"
)

var count = 0

type AuthServer struct {
	auth.UnimplementedAuthServiceServer
}

func (s *AuthServer) Login(ctx context.Context, req *auth.LoginRequest) (*auth.Response, error) {
	count++
	fmt.Println("Login request received count:", count)

	user, err := models.GetUserByIdentifier(database.DB, req.Identifier)
	
	if err != nil {
		return &auth.Response{
			Status: auth.ResponseStatus_INVALID_CREDENTIALS,
			Metadata: &auth.Response_Message{Message: "Invalid credentials"},
		}, nil
	}

	return &auth.Response{
		Status: auth.ResponseStatus_INVALID_CREDENTIALS,
		Metadata: &auth.Response_Token{Token: user.Identifier},
	}, nil
}

func (s *AuthServer) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.Response, error) {
	fmt.Println("Register request received")
	return &auth.Response{
		Status: auth.ResponseStatus_OK,
	}, nil
}