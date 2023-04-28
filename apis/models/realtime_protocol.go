package models

type RealtimeMessage struct {
	Action string
	Data   any
}

type RealtimePlayerConnectionStateChange struct {
	Name        string `json:"name"`
	IsConnected bool   `json:"isConnected"`
}

type RealtimePing struct{}
