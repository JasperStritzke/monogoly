package models

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

const logMaxEvents = 2048

type Game struct {
	GameId     string     `json:"gameId"`
	GameSecret string     `json:"-"`
	Players    []*Player  `json:"players"`
	LogEvents  []LogEvent `json:"logEvents"`
	LastAction time.Time  `json:"-"`
	Data       GameData   `json:"data"`
}

type GameData struct {
	Currency       string  `json:"currency"`
	IsLocked       bool    `json:"isLocked"`
	InitialBalance float32 `json:"initialBalance"`
	AmountGo       float32 `json:"amountGo"`
}

func (g *Game) SetPassword(password string) {
	hash := md5.Sum([]byte(g.GameId + "#" + password))
	g.GameSecret = hex.EncodeToString(hash[:])
}

func (g *Game) CheckPassword(password string) bool {
	hash := md5.Sum([]byte(g.GameId + "#" + password))
	return g.GameSecret == hex.EncodeToString(hash[:])
}

func (g *Game) LogEvent(t LogType, p []interface{}) {
	event := LogEvent{
		Timestamp: time.Now(),
		Type:      t,
		Payload:   p,
	}

	if len(g.LogEvents) >= logMaxEvents {
		//Truncate first event
		g.LogEvents = append(g.LogEvents[1:], event)
		return
	}

	g.LogEvents = append(g.LogEvents, event)
}

func (g *Game) GetPlayerWithName(name string) *Player {
	for _, player := range g.Players {
		if player.Name == name {
			return player
		}
	}

	return nil
}

func (g *Game) BroadcastRaw(msg any) {
	for _, player := range g.Players {
		player.SendRaw(msg)
	}
}

func (g *Game) BroadcastPlayerConnectionUpdate(name string, isConnected bool) {
	g.BroadcastRaw(RealtimePlayerConnectionStateChange{
		Name:        name,
		IsConnected: isConnected,
	})
}
