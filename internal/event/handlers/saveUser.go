package handlers

import (
	"encoding/json"

	"example.com/control-tower-agent/internal/hub"
)

func HandleSaveUser(data interface{}, h *hub.Hub) {
	// Implement the logic to save a user
	response := EventResponse{
		Event: "saveUser",
		Data:  data,
	}
	msg, _ := json.Marshal(response)
	h.Broadcast <- msg
}

func HandleRotateKeys(data interface{}, h *hub.Hub) {
	// Implement the logic to rotate keys
	response := EventResponse{
		Event: "rotateKeys",
		Data:  data,
	}
	msg, _ := json.Marshal(response)
	h.Broadcast <- msg
}
