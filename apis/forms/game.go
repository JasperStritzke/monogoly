package forms

import "monogoly/apis/models"

type GameCreate struct {
	Currency       string  `json:"currency" binding:"required"`
	AmountGo       float32 `json:"amountGo" binding:"required"`
	InitialBalance float32 `json:"initialBalance" binding:"required"`
	Password       string  `json:"password" binding:"required"`
}

type GameCreateResponse struct {
	GameId   string `json:"gameId"`
	Password string `json:"password"`
}

type GameJoin struct {
	Name     string `json:"name" binding:"required"`
	GameId   string `json:"gameId" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type GameJoinResponse struct {
	You  models.Player `json:"you"`
	Game models.Game   `json:"game"`
}
