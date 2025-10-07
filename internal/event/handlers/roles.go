package handlers

import (
	"encoding/json"

	"example.com/control-tower-agent/internal/hub"
)

func HandleAssignRoles(data interface{}, h *hub.Hub) {
	// Implement the logic to assign roles
	response := EventResponse{
		Event: "assignRoles",
		Data:  data,
	}
	msg, _ := json.Marshal(response)
	h.Broadcast <- msg
}
