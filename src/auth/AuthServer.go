package Auth

import (
	"github.com/JacobRWebb/authentication-microservice/pb/auth"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Identifier string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
}

func AttachServer(grpcServer *grpc.Server) {
	auth.RegisterAuthServiceServer(grpcServer, &AuthServer{})
}