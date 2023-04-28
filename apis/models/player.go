package models

type Player struct {
	Name    string   `json:"name"`
	Balance float32  `json:"balance"`
	IsHost  bool     `json:"isHost"`
	Channel chan any `json:"-"` //queue to send messages
}

func (p *Player) GenerateSecret(game *Game) string {
	return game.GameSecret + "#" + p.Name
}

func (p *Player) SendRaw(message any) {
	if p.Channel != nil {
		p.Channel <- message
	}
}
