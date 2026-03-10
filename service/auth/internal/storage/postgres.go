package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/meis1kqt/go-monorepo-chat.git/pkg/models"
)

type PostgresStorage struct {
	pool *pgxpool.Pool
}

func NewStorage(pool *pgxpool.Pool) *PostgresStorage {
	return &PostgresStorage{pool: pool}
}

func (s *PostgresStorage) SaveUser(ctx context.Context, email string, passHash []byte) error {
	query := `INSERT INTO users (email, password_hash) VALUES ($1, $2)`
	_, err := s.pool.Exec(ctx, query, email, passHash)
	if err != nil {
		return err	
	}
	return nil
}
func (s *PostgresStorage) GetUser(ctx context.Context, email string) (*models.User, error) {
	query := `SELECT id, email, password_hash FROM users WHERE email = $1`
	row := s.pool.QueryRow(ctx, query, email)

	var user models.User
	err := row.Scan(&user.ID, &user.Email, &user.PassHash)
	if err != nil {
		return nil, err
	}
	return &user, nil
}