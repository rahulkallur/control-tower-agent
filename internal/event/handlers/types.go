package handlers

type IncomingEvent struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}

type EventResponse struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}
