package models

type WebsocketDataPayload struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Data any    `json:"data"`
}

var SocketDataPaylod = make(chan WebsocketDataPayload, 3)
