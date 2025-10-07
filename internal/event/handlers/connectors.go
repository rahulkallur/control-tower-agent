package handlers

import (
	"encoding/json"

	"example.com/control-tower-agent/internal/hub"
)

func HandleSyncConnectors(data interface{}, h *hub.Hub) {
	// Implement the logic to sync connectors
	response := EventResponse{
		Event: "syncConnectors",
		Data:  data,
	}
	msg, _ := json.Marshal(response)
	h.Broadcast <- msg
}

func HandleSyncTransforms(data interface{}, h *hub.Hub) {
	// Implement the logic to sync transforms
	response := EventResponse{
		Event: "syncTransforms",
		Data:  data,
	}
	msg, _ := json.Marshal(response)
	h.Broadcast <- msg
}

func HandleSyncCommands(data interface{}, h *hub.Hub) {
	// Implement the logic to sync commands
	response := EventResponse{
		Event: "syncCommands",
		Data:  data,
	}
	msg, _ := json.Marshal(response)
	h.Broadcast <- msg
}
