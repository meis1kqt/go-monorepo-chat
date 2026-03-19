package main

import (
	"context"
	"net"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/meis1kqt/go-monorepo-chat.git/pkg/slog"
	"github.com/meis1kqt/go-monorepo-chat.git/service/auth/internal/config"
	authgrpc "github.com/meis1kqt/go-monorepo-chat.git/service/auth/internal/grpc"
	"github.com/meis1kqt/go-monorepo-chat.git/service/auth/internal/service"
	"github.com/meis1kqt/go-monorepo-chat.git/service/auth/internal/storage"
	"google.golang.org/grpc"
)


func main(){
	
	Config := config.MustLoadConfig()

	logger := slog.NewLogger(Config.Environment)

	logger.Info("Starting auth service...")

	context := context.Background()

	db, err := pgxpool.New(context, Config.Database.DSN()) 
	if err != nil {
		logger.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	storage := storage.NewStorage(db)

	authService := service.New(logger, storage, Config.JWT.Secret, time.Duration(Config.JWT.Expiration))

	grpcServer := grpc.NewServer()

	authgrpc.Register(grpcServer, authService)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.Error("Failed to listen", "error", err)
		os.Exit(1)
	}
	logger.Info("Auth service is running on port 50051")

	if err := grpcServer.Serve(lis); err != nil {
		logger.Error("Failed to serve gRPC server", "error", err)
		os.Exit(1)
	}
}