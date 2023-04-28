package services

import (
	"errors"
	"log"
	"monogoly/apis/forms"
	"monogoly/apis/models"
	"monogoly/util"
	"time"
)

type GameService struct {
	ticker        *time.Ticker
	games         map[string]*models.Game
	createQueue   chan *models.Game
	joinQueue     chan *joinQueueEntry
	realtimeQueue chan *realtimeQueueEntry //used for connection updates and disconnection updates
}

type realtimeQueueEntry struct {
	Game    *models.Game
	Player  *models.Player
	Channel chan any
}

type joinQueueEntry struct {
	Player *models.Player
	GameId string
}

func NewGameService() *GameService {
	gameService := &GameService{
		ticker:        time.NewTicker(time.Minute),
		games:         make(map[string]*models.Game),
		createQueue:   make(chan *models.Game),
		joinQueue:     make(chan *joinQueueEntry),
		realtimeQueue: make(chan *realtimeQueueEntry),
	}

	go gameService.run()

	return gameService
}

func (g *GameService) run() {
	for {
		select {
		case createdGame := <-g.createQueue:
			g.createGame(createdGame)
		case joinEntry := <-g.joinQueue:
			g.joinGame(joinEntry)
		case realtimeEntry := <-g.realtimeQueue:
			g.connectRealtime(realtimeEntry)
		case <-g.ticker.C:
			g.checkInactiveGamesAndPing()
		}
	}
}

func (g *GameService) AuthBySecret(secret string) *models.Game {
	for _, game := range g.games {
		if game.GameSecret == secret {
			return game
		}
	}

	return nil
}

const gameMaxInactive = time.Hour * 6

func (g *GameService) checkInactiveGamesAndPing() {
	for _, game := range g.games {
		if time.Now().Sub(game.LastAction) > gameMaxInactive {
			g.timeoutGame(game)
			continue
		}

		game.BroadcastRaw(models.RealtimePing{})
	}
}

func (g *GameService) timeoutGame(game *models.Game) {
	log.Println("Game", game.GameId, "timed out after", gameMaxInactive.Hours(), "hours.")

	//TODO disconnect all players with information, that game has been closed.

	/*for _, player := range game.Players {
		player.
	}*/

	delete(g.games, game.GameId)
	game = nil
}

func (g *GameService) createGame(game *models.Game) {
	g.games[game.GameId] = game
}

func (g *GameService) RequestGameStart(request forms.GameCreate) *models.Game {
	newGame := &models.Game{
		GameId:     util.RandomGameId(),
		Players:    nil,
		LastAction: time.Now(),
		Data: models.GameData{
			Currency:       request.Currency,
			IsLocked:       false,
			InitialBalance: request.InitialBalance,
			AmountGo:       request.AmountGo,
		},
	}
	newGame.SetPassword(request.Password)

	g.createQueue <- newGame

	return newGame
}

func (g *GameService) joinGame(entry *joinQueueEntry) {
	game, exists := g.games[entry.GameId]

	if !exists {
		return
	}

	if len(game.Players) == 0 {
		entry.Player.IsHost = true
	}

	game.Players = append(game.Players, entry.Player)

	log.Println(entry.Player.Name, "joined the game", entry.GameId)
}

func (g *GameService) RequestGameJoin(request forms.GameJoin) (*models.Game, *models.Player, error) {
	game, exists := g.games[request.GameId]

	if !exists {
		return nil, nil, errors.New("no game found with that id")
	}

	if !game.CheckPassword(request.Password) {
		return nil, nil, errors.New("password is invalid")
	}

	if game.Data.IsLocked {
		return nil, nil, errors.New("the game is locked")
	}

	if p := game.GetPlayerWithName(request.Name); p != nil {
		return game, p, nil
	} else {
		player := &models.Player{
			Name:    request.Name,
			Balance: game.Data.InitialBalance,
		}

		g.joinQueue <- &joinQueueEntry{
			Player: player,
			GameId: game.GameId,
		}

		gameCopy := *game
		gameCopy.Players = append(gameCopy.Players, player)

		return &gameCopy, player, nil
	}
}

func (g *GameService) connectRealtime(realtimeEntry *realtimeQueueEntry) {
	realtimeEntry.Player.Channel = realtimeEntry.Channel
	realtimeEntry.Game.BroadcastPlayerConnectionUpdate(realtimeEntry.Player.Name, realtimeEntry.Channel != nil)
}

func (g *GameService) RequestRealtimeChange(game *models.Game, player *models.Player, channel chan any) {
	g.realtimeQueue <- &realtimeQueueEntry{
		Game:    game,
		Player:  player,
		Channel: channel,
	}
}
