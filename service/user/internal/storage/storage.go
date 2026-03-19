package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/meis1kqt/go-monorepo-chat.git/service/user/internal/domain"
)

type Storage struct {
	db *pgxpool.Pool
}

func NewStorage(db *pgxpool.Pool) *Storage {
	return &Storage{db: db}
}


func (s *Storage) SaveMessage(ctx context.Context, message *domain.Message) error {
	query := `INSERT INTO messages (dialog_id, from_id, text, created_at) VALUES ($1, $2, $3, $4)`
	_, err := s.db.Exec(ctx, query, message.DialogID, message.From, message.Text, message.CreatedAt)	
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetMessage(ctx context.Context, dialog int64 ) ([]*domain.Message, error) {

	query := `SELECT id, dialog_id, from_id, text, created_at FROM messages WHERE dialog_id = $1`

	var message []*domain.Message

	rows, err := s.db.Query(ctx, query, dialog)

	if err != nil {
		return nil, err
	}	
	defer rows.Close()

	for rows.Next() {
		msg := &domain.Message{}
		if err := rows.Scan(&msg.ID, &msg.DialogID, &msg.From, &msg.Text, &msg.CreatedAt); err != nil {
			return nil, err
		}
		message = append(message, msg)
	}
	return message, nil
}
func (s *Storage) GetUser(ctx context.Context, userID int64) (*domain.User, error) {
	query := `SELECT id, username, avatar, is_online From users WHERE id = $1`

	row := s.db.QueryRow(ctx, query, userID)

	user := &domain.User{}

	err := row.Scan(&userID, &user.Username, &user.Avatar, &user.IsOnline)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Storage) GetChats(ctx context.Context, userID int64) ([]*domain.Dialog, error) {
	query := `SELECT id, creator_id, user_id, created_at from dialogs WHERE creator_id = $1 or user_id = $1`

	var dialogs []*domain.Dialog

	rows, err := s.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		dialog := &domain.Dialog{}
		if err := rows.Scan(&dialog.ID, &dialog.CreatorID, &dialog.UserID, &dialog.CreatedAt); err != nil {
			return nil, err
		}
		dialogs = append(dialogs, dialog)
	}
	return dialogs, nil

}

func (s *Storage) CreateDialog(ctx context.Context, creatorID, userID int64) (*domain.Dialog, error) {
	query := `INSERT INTO dialogs (creator_id, user_id) VAlUES ($1,$2) RETURNING id`

	dialog := &domain.Dialog{}

	row := s.db.QueryRow(ctx, query, creatorID, userID)

	if err := row.Scan(&dialog.ID); err != nil {
		return nil, err
	}
	dialog.CreatorID = creatorID
	dialog.UserID = userID

	return dialog, nil
}

func (s *Storage) GetDialog(ctx context.Context, creatorID, userID int64) (*domain.Dialog, error) {
	query := `SELECT id, creator_id, user_id, created_at from dialogs WHERE (creator_id = $1 and user_id = $2) or (creator_id = $2 and user_id = $1)`

	dialog := &domain.Dialog{}

	row := s.db.QueryRow(ctx, query, creatorID, userID)

	if err := row.Scan(&dialog.ID, &dialog.CreatorID, &dialog.UserID, &dialog.CreatedAt); err != nil {
		return nil, err
	}

	return dialog, nil
}

func (s *Storage) SearchUser(ctx context.Context, username string) ([]*domain.User, error) {
	query := `SELECT id, username, avatar, is_online from users where username like $1`

	var users []*domain.User

	rows, err := s.db.Query(ctx, query, "%"+username+"%")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		user := &domain.User{}
		if err := rows.Scan(&user.ID, &user.Username, &user.Avatar, &user.IsOnline); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}