package grpc

import "context"


type Auth interface{
	RegisterUser(ctx context.Context, email , password string) error
	Login(ctx context.Context, email, password string)(string, error)
}

type ServerApi struct {
	
}


