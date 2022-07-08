package Auth

import (
	"github.com/JacobRWebb/authentication-microservice/pb/auth"
	"google.golang.org/grpc"
)

func AttachServer(grpcServer *grpc.Server) {
	auth.RegisterAuthServiceServer(grpcServer, &AuthServer{})
}