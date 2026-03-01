package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/meis1kqt/go-monorepo-chat.git/pkg/models"
)



func GenerateToken(user *models.User, jwtSecret string, duration int) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Duration(duration) * time.Minute).Unix(),
	})
	
	return token.SignedString([]byte(jwtSecret))
}