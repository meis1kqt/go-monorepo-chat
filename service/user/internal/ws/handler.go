package ws

import "github.com/meis1kqt/go-monorepo-chat.git/service/user/internal/service"


type Handler struct {
	hub *Hub
	service *service.ChatService
}

func New(hub *Hub, service *service.ChatService) *Handler {
	return &Handler{hub: hub, service: service}
}