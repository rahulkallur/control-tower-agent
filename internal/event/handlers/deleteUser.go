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

func HandleUpdatePassword(data interface{}, h *hub.Hub) {
	// Implement the logic to update a user's password
	response := EventResponse{
		Event: "updatePassword",
		Data:  data,
	}
	msg, _ := json.Marshal(response)
	h.Broadcast <- msg
}

func HandleUpdateRole(data interface{}, h *hub.Hub) {
	// Implement the logic to update a user's role
	response := EventResponse{
		Event: "updateRole",
		Data:  data,
	}
	msg, _ := json.Marshal(response)
	h.Broadcast <- msg
}

func HandleUpdatePersonalDetails(data interface{}, h *hub.Hub) {
	// Implement the logic to update a user's personal details
	response := EventResponse{
		Event: "updatePersonalDetails",
		Data:  data,
	}
	msg, _ := json.Marshal(response)
	h.Broadcast <- msg
}
