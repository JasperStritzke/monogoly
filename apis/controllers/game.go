package controllers

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"io"
	"monogoly/apis/forms"
	"monogoly/apis/models"
	"monogoly/apis/services"
	"monogoly/util"
	"net/http"
	"reflect"
)

type GameController struct {
	gameService *services.GameService
}

func NewGameController(gameService *services.GameService) GameController {
	return GameController{gameService: gameService}
}

func (g GameController) Create(c *gin.Context) {
	var body forms.GameCreate

	if err := c.BindJSON(&body); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if !util.Currency.ExistsCurrency(body.Currency) {
		_ = c.AbortWithError(http.StatusBadRequest, errors.New("invalid currency"))
		return
	}

	game := g.gameService.RequestGameStart(body)

	c.JSON(http.StatusOK, forms.GameCreateResponse{
		GameId:   game.GameId,
		Password: body.Password,
	})
}

func (g GameController) Join(c *gin.Context) {
	var body forms.GameJoin

	if err := c.BindJSON(&body); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	game, player, err := g.gameService.RequestGameJoin(body)

	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.Header("Authorization", base64.URLEncoding.EncodeToString([]byte(player.GenerateSecret(game))))

	c.JSON(http.StatusOK, forms.GameJoinResponse{
		You:  *player,
		Game: *game,
	})
}

func (g GameController) Realtime(c *gin.Context) {
	gameAny := c.MustGet("game")
	playerAny := c.MustGet("player")

	player := playerAny.(*models.Player)
	game := gameAny.(*models.Game)

	channel := make(chan any)
	g.gameService.RequestRealtimeChange(game, player, channel)
	defer g.gameService.RequestRealtimeChange(game, player, nil)

	clientGone := c.Request.Context().Done()
	connected := c.Stream(func(w io.Writer) bool {
		select {
		case <-clientGone:
			return false
		case message := <-channel:
			msg, _ := json.Marshal(message)
			c.SSEvent(reflect.ValueOf(message).Type().Name(), string(msg))
			return true
		}
	})

	if !connected {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}
