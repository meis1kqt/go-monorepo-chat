package service

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/meis1kqt/go-monorepo-chat.git/service/auth/internal/dopmain"
	"golang.org/x/crypto/bcrypt"
)

type Storage interface{
	GetUser(ctx context.Context, email string)(*dopmain.User, error)
	SaveUser(ctx context.Context, email string, passHash []byte) error
}

type AuthService struct {
	log *slog.Logger
	Storage Storage
	JwtSecret string
	TTL time.Duration
}

func New(log *slog.Logger, Storage Storage, JwtSecret string, TTL time.Duration) *AuthService {
	return &AuthService{log: log ,Storage: Storage, JwtSecret: JwtSecret, TTL: TTL}
}


func (a *AuthService) RegisterUser(ctx context.Context, email , password string) error {
	if email == "" || password == "" {
		a.log.Error("password or email", "error")
		return  fmt.Errorf("wrong something")
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		a.log.Error("failed to generate password hash", "error", err)
		return  err
	}
	err = a.Storage.SaveUser(ctx, email, passHash)

	if err != nil {
		defer a.log.Error("storage", "error")
		return fmt.Errorf("failed to save user: %w", err)
	}
	return nil
}