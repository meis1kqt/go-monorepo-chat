package main

import (
	"github.com/meis1kqt/go-monorepo-chat.git/pkg/slog"
	"github.com/meis1kqt/go-monorepo-chat.git/service/auth/internal/config"
)


func main(){
	
	Config := config.MustLoadConfig()

	logger := slog.NewLogger(Config.Environment)

	logger.Info("Starting auth service...")

}