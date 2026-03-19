package ws

import (
	"context"
	"encoding/json"

	"github.com/gorilla/websocket"
	"github.com/meis1kqt/go-monorepo-chat.git/service/user/internal/service"
)
type InComingMessage struct {
    To   int64  `json:"to"`
    Text string `json:"text"`
}

type Client struct {
	conn *websocket.Conn
	userID int64
	send chan[]byte
	hub *Hub
	service *service.ChatService
}

func (c *Client) ReadPump() {
	for {
		_, data, err := c.conn.ReadMessage()
		if err != nil {
			c.hub.unregister <- c
			break
		}
		var msg InComingMessage
		if err := json.Unmarshal(data, &msg); err != nil {
			break
		}
		ctx := context.Background()
		err = c.service.SendMessage(ctx, c.userID, msg.To, msg.Text)
		if err != nil {
			break
		}
		if recipient, ok := c.hub.clients[msg.To]; ok {
			recipient.send <- data
		}
	}
}

func (c *Client) WritePump() {
	for msg := range c.send {
		if err := c.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
}