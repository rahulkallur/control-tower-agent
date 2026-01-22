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

func HandleTffTokenResponse(data interface{}, h *hub.Hub) {
	// Implement the logic to handle TFF token responses
	response := EventResponse{
		Event: "tffTokenResponse",
		Data:  data,
	}
	msg, _ := json.Marshal(response)
	h.Broadcast <- msg
}
