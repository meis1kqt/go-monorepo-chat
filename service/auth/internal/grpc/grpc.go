package authgrpc

import (
	"context"

	pb "github.com/meis1kqt/go-monorepo-chat.git/protos/auth/sso"
	"google.golang.org/grpc"
)



type Auth interface{
	RegisterUser(ctx context.Context, email , password string) error
	Login(ctx context.Context, email, password string)(string, error)
}

type ServerApi struct {
	pb.UnimplementedAuthServiceServer
	auth Auth
}

func Register(grpc *grpc.Server, auth Auth) {
	pb.RegisterAuthServiceServer(grpc, &ServerApi{auth: auth})
}

func (s *ServerApi) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	err := s.auth.RegisterUser(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &pb.RegisterResponse{}, nil
}

func (s *ServerApi) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	token, err := s.auth.Login(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &pb.LoginResponse{Token: token}, nil
}	

