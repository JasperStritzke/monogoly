package main

import (
	"monogoly/env"
)

import (
	"log"
	"monogoly/apis"
	"monogoly/ui"
)

func main() {
	log.Println("Starting Monogoly v" + env.Version + ".")

	e := apis.InitApi()

	ui.BindUiRoutes(e)

	log.Fatal(e.Run(":8080"))
}
