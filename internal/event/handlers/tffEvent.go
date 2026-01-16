package handlers

import (
	"encoding/json"

	"example.com/control-tower-agent/internal/hub"
)

func HandleTffTokens(data interface{}, h *hub.Hub) {
	// Implement the logic to handle TFF tokens
	response := EventResponse{
		Event: "tffTokens",
		Data:  data,
	}
	msg, _ := json.Marshal(response)
	h.Broadcast <- msg
}
