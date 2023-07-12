package logic

import (
	"sync"
)

var Hubs = NewHub()

type Hub struct {
	clients    sync.Map
	Broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
	lock       sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.lock.Lock()
			h.clients.Store(client.connId, client)
			h.lock.Unlock()
		case client := <-h.unregister:
			h.lock.Lock()
			h.clients.Delete(client.connId)
			h.lock.Unlock()
		case message := <-h.Broadcast:
			h.clients.Range(func(_, value interface{}) bool {
				go func() {
					client := value.(*Client)
					err := client.WriteMessage(message)
					if err != nil {
						h.lock.Lock()
						h.clients.Delete(client.connId)
						h.lock.Unlock()
					}
				}()
				return true
			})
		}
	}
}

func (h *Hub) FindOneByImId(connId string) *Client {
	h.lock.RLock()
	defer h.lock.RUnlock()
	value, ok := h.clients.Load(connId)
	if !ok {
		return nil
	}
	return value.(*Client)
}
