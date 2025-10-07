package handlers

import (
	"encoding/json"

	"example.com/control-tower-agent/internal/hub"
)

func HandleBillingEvent(data interface{}, h *hub.Hub) {
	// Implement the logic to handle billing event
	response := EventResponse{
		Event: "meterEvent",
		Data:  data,
	}
	msg, _ := json.Marshal(response)
	h.Broadcast <- msg
}
