package ws

import (
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/meis1kqt/go-monorepo-chat.git/service/user/internal/service"
)


type Handler struct {
	hub *Hub
	service *service.ChatService
}

var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func New(hub *Hub, service *service.ChatService) *Handler {
	return &Handler{hub: hub, service: service}
}

func (h *Handler) ServeWs(w http.ResponseWriter, r *http.Request) {
	userIDSTR := r.URL.Query().Get("user_id")
	userID, err := strconv.ParseInt(userIDSTR, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	conn, err := upgrade.Upgrade(w,r,nil)
	if err != nil {
		return
	}

	client := &Client{
        conn:    conn,
        userID:  userID,
        send:    make(chan []byte, 256),
        hub:     h.hub,
        service: h.service,
    }

	h.hub.register <- client

	go client.ReadPump()
	go client.WritePump()
	
}