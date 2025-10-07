package handlers

import (
	"encoding/json"

	"example.com/control-tower-agent/internal/hub"
)

func HandleContract(data interface{}, h *hub.Hub) {
	// Implement the logic to handle environment contract
	response := EventResponse{
		Event: "billingContract",
		Data:  data,
	}
	msg, _ := json.Marshal(response)
	h.Broadcast <- msg
}
