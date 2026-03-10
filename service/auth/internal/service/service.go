package service

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/meis1kqt/go-monorepo-chat.git/pkg/jwt"
	"github.com/meis1kqt/go-monorepo-chat.git/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

type Storage interface{
	GetUser(ctx context.Context, email string)(*models.User, error)
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
	if email == "" {
		a.log.Error("email", "error")
		return  fmt.Errorf("wrong something")
	}
	if  password == "" {
		a.log.Error("password", "error")
		return  fmt.Errorf("wrong something")
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		a.log.Error("failed to generate password hash", "error", err)
		return  err
	}
	err = a.Storage.SaveUser(ctx, email, passHash)

	if err != nil {
  		return fmt.Errorf("failed to save user: %w", err)
	}
	return nil
}

func (a *AuthService) Login(ctx context.Context, email, password string)(string, error) {
	if email == "" {
		a.log.Error("in login service","email", "error")
		return  "", fmt.Errorf("wrong something")
	}
	if  password == "" {
		a.log.Error("in login service","password", "error")
		return  "", fmt.Errorf("wrong something")
	}
	
	user, err := a.Storage.GetUser(ctx, email)

	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword(user.PassHash, []byte(password)); err != nil {
		return "", err
	}

	token, err := jwt.GenerateToken(user, a.JwtSecret, int(a.TTL))

	if err != nil {
		return "", err
	}

	return token, nil
}