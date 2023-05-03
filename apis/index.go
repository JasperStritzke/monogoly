package apis

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"math/rand"
	"monogoly/apis/controllers"
	"monogoly/apis/middlewares"
	"monogoly/apis/services"
	"net/http"
	"strings"
)

func InitApi() *gin.Engine {
	e := gin.New()

	e.Use(gin.Recovery())
	e.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	api := e.Group("/api")
	api.GET("/status", status)

	//game routes
	gameService := services.NewGameService()
	gameController := controllers.NewGameController(gameService)

	game := api.Group("/game", middlewares.ErrorCatcher)
	game.POST("/create", gameController.Create)
	game.POST("/join", gameController.Join)

	game.GET("/realtime", middlewares.AuthMiddleware(gameService), gameController.Realtime)

	return e
}

func status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Im fine. Y" + strings.Repeat("E", rand.Intn(8)+1) + "T"})
}
