package handlers

import (
	"encoding/json"

	"example.com/control-tower-agent/internal/hub"
)

func HandleDeleteUser(data interface{}, h *hub.Hub) {
	// Implement the logic to delete a user
	response := EventResponse{
		Event: "deleteUser",
		Data:  data,
	}
	msg, _ := json.Marshal(response)
	h.Broadcast <- msg
}

func HandleSuspendUser(data interface{}, h *hub.Hub) {
	// Implement the logic to suspend a user
	response := EventResponse{
		Event: "suspendUser",
		Data:  data,
	}
	msg, _ := json.Marshal(response)
	h.Broadcast <- msg
}
