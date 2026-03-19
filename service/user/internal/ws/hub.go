package ws


type Hub struct {
	clients map[int64]*Client
	register chan *Client
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		clients: make(map[int64]*Client),
		register: make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client.userID] = client
		case client := <- h.unregister:
			delete(h.clients, client.userID)
		} 
	}
}