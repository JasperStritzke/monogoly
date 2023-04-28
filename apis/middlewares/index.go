package middlewares

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"monogoly/apis/services"
	"net/http"
	"strings"
)

type WrappedError struct {
	Message string `json:"message"`
}

func ErrorCatcher(c *gin.Context) {
	c.Next()
	if len(c.Errors) > 0 {
		status := c.Writer.Status()

		if status == 0 {
			status = http.StatusBadRequest
		}

		c.JSON(status, WrappedError{Message: c.Errors[0].Error()})
	}
}

func AuthMiddleware(gameService *services.GameService) func(c *gin.Context) {
	return func(c *gin.Context) {
		authBytes, err := base64.URLEncoding.DecodeString(c.GetHeader("Authorization"))

		auth := string(authBytes)

		if err != nil || len(auth) == 0 || !strings.Contains(auth, "#") {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		authSplit := strings.Split(auth, "#")
		secret, playerName := authSplit[0], authSplit[1]

		game := gameService.AuthBySecret(secret)

		if game == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		player := game.GetPlayerWithName(playerName)

		if player == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("game", game)
		c.Set("player", player)
	}
}
