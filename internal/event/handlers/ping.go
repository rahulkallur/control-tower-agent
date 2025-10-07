package handlers

import (
	"encoding/json"

	"example.com/control-tower-agent/internal/hub"
)

func HandlePing(h *hub.Hub) {
	response := EventResponse{
		Event: "pong",
		Data:  "pong",
	}
	msg, _ := json.Marshal(response)
	h.Broadcast <- msg
}
