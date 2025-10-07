package hub

import (
	"encoding/json"
	"log"
	"sync"

	"example.com/control-tower-agent/internal/client"
	"example.com/control-tower-agent/internal/crypto"
	"example.com/control-tower-agent/internal/models"
)

type Hub struct {
	Clients    map[string]*client.Client   // Connected clients, keyed by client ID
	Broadcast  chan []byte                 // Channel for broadcasting messages to all clients
	Unicast    chan models.UnicastMessage  // Channel for sending messages to specific clients
	Register   chan *client.Client         // Channel for registering new clients
	Unregister chan *client.Client         // Channel for unregistering clients
	Incoming   chan models.IncomingMessage // Channel for incoming messages from clients
	Mutex      sync.RWMutex                // Synchronize access to the Clients map
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[string]*client.Client),
		Broadcast:  make(chan []byte),
		Unicast:    make(chan models.UnicastMessage),
		Register:   make(chan *client.Client),
		Unregister: make(chan *client.Client),
		Incoming:   make(chan models.IncomingMessage),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Mutex.Lock()
			h.Clients[client.ID] = client
			h.Mutex.Unlock()
		case client := <-h.Unregister:
			h.Mutex.Lock()
			if _, ok := h.Clients[client.ID]; ok {
				delete(h.Clients, client.ID)
				close(client.Send)
			}
			h.Mutex.Unlock()
		case message := <-h.Broadcast:
			encrypted, err := crypto.Encrypt(string(message))
			if err != nil {
				log.Printf("encrypt error: %v", err)
				continue
			}
			data := map[string]string{"data": encrypted}
			jsonData, _ := json.Marshal(data)
			h.Mutex.RLock()
			for _, client := range h.Clients {
				select {
				case client.Send <- jsonData:
				default:
					close(client.Send)
					delete(h.Clients, client.ID)
				}
			}
			h.Mutex.RUnlock()
		case unicast := <-h.Unicast:
			encrypted, err := crypto.Encrypt(string(unicast.Message))
			if err != nil {
				log.Printf("encrypt error: %v", err)
				continue
			}
			data := map[string]string{"data": encrypted}
			jsonData, _ := json.Marshal(data)
			h.Mutex.RLock()
			if client, ok := h.Clients[unicast.ClientID]; ok {
				select {
				case client.Send <- jsonData:
				default:
					close(client.Send)
					delete(h.Clients, client.ID)
				}
			}
			h.Mutex.RUnlock()
		}
	}
}
