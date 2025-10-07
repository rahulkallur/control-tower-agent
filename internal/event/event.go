package event

import (
	"encoding/json"
	"log"

	"example.com/control-tower-agent/internal/crypto"
	"example.com/control-tower-agent/internal/event/handlers"
	"example.com/control-tower-agent/internal/hub"
)

func HandleEvent(data []byte, clientID string, h *hub.Hub) {
	// Unmarshal the incoming JSON to extract the encrypted data
	var incomingData struct {
		Data string `json:"data"`
	}
	if err := json.Unmarshal(data, &incomingData); err != nil {
		log.Printf("error unmarshaling incoming data: %v", err)
		return
	}

	// Decrypt the data
	decryptedData, err := crypto.Decrypt(incomingData.Data)
	if err != nil {
		log.Printf("error decrypting data: %v", err)
		return
	}

	// Unmarshal to IncomingEvent
	var incoming handlers.IncomingEvent
	if err := json.Unmarshal([]byte(decryptedData), &incoming); err != nil {
		log.Printf("error unmarshaling incoming event: %v", err)
		return
	}
	log.Printf("Received event: %s from client: %s", incoming.Event, clientID)

	// Handle Events
	switch incoming.Event {
	case "ping":
		handlers.HandlePing(h)
	case "saveUser":
		handlers.HandleSaveUser(incoming.Data, h)
	case "deleteUser":
		handlers.HandleDeleteUser(incoming.Data, h)
	case "billingContract":
		handlers.HandleContract(incoming.Data, h)
	case "meterEvent":
		handlers.HandleBillingEvent(incoming.Data, h)
	case "syncConnectors":
		handlers.HandleSyncConnectors(incoming.Data, h)
	case "syncTransforms":
		handlers.HandleSyncTransforms(incoming.Data, h)
	case "syncCommands":
		handlers.HandleSyncCommands(incoming.Data, h)
	case "assignRoles":
		handlers.HandleAssignRoles(incoming.Data, h)
	case "rotateKeys":
		handlers.HandleRotateKeys(incoming.Data, h)
	default:
		log.Printf("unknown event: %s", incoming.Event)
	}
}
