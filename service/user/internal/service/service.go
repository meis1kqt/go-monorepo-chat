package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/meis1kqt/go-monorepo-chat.git/service/user/internal/domain"
)

type Storage interface {
	SaveMessage(ctx context.Context, message *domain.Message) error
	GetMessage(ctx context.Context, dialog int64 ) ([]*domain.Message, error)
	GetUser(ctx context.Context, userID int64) (*domain.User, error)
	GetChats(ctx context.Context, userID int64) ([]*domain.Dialog, error)
	CreateDialog(ctx context.Context, creatorID, userID int64) (*domain.Dialog, error)
	GetDialog(ctx context.Context, creatorID, userID int64) (*domain.Dialog, error)
	SearchUser(ctx context.Context, username string) ([]*domain.User, error)
}

type ChatService struct {
	log *slog.Logger
	storage Storage
}

func New(log *slog.Logger, storage Storage) *ChatService {
	return &ChatService{log: log, storage: storage}
}


func (c *ChatService) SendMessage(ctx context.Context, creatorID, userID int64, text string) error {
	if text == "" {
		return fmt.Errorf("text is empty")
	}
	dialog, err := c.storage.GetDialog(ctx, creatorID, userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			dialog, err = c.storage.CreateDialog(ctx, creatorID, userID)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	msg := &domain.Message{
		DialogID: dialog.ID,
		From: creatorID,
		Text: text,
	}
	
	return c.storage.SaveMessage(ctx, msg)
}

func (c *ChatService) GetMessage(ctx context.Context, dialogID int64) ([]*domain.Message, error) {
	msgs, err := c.storage.GetMessage(ctx, dialogID)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}

func (c *ChatService) GetDialogs(ctx context.Context, userID int64) ([]*domain.Dialog, error) {
	chats, err := c.storage.GetChats(ctx, userID)
	if err != nil {
		return nil, err
	}

	return chats, nil
}

func (c *ChatService) GetUser(ctx context.Context, userID int64) (*domain.User, error) {
	user, err := c.storage.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}


func (c *ChatService) SearchUsers(ctx context.Context, username string) ([]*domain.User, error) {
	users, err := c.storage.SearchUser(ctx, username)
	if err != nil {
		return nil, err
	}

	return users, nil
}