package main

import "workshop-go-guess-me/game"

func main() {
	gameParams := game.InitGame()
	gameParams.StartGame()
	gameParams.Result()
}
