package game

import (
	"fmt"
	"math/rand"
	"time"
	"workshop-go-guess-me/limit"
	"workshop-go-guess-me/util"
)

type GameParameters struct {
	mainLimit limit.Limit
	random    int
	loop      int
}

func getRandomNumber(limits limit.Limit) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(limits.Higher)
}

func initConditionNumber(nb int) bool {
	return nb > 0
}

func inGameConditionNmber(nb int) bool {
	return nb >= 0
}

func InitGame() GameParameters {
	mainLimits := limit.Limit{
		Lower:  0,
		Higher: util.ReadInt("Lower limit is 0, choose you're upper limit: ", initConditionNumber),
	}
	fmt.Printf("Limits defined to -> %s.\n", mainLimits)
	numberToFind := getRandomNumber(mainLimits)
	fmt.Println("Game initialized !\nLets's Golang!")
	return GameParameters{
		mainLimit: mainLimits,
		random:    numberToFind,
		loop:      0,
	}
}

func (params *GameParameters) StartGame() {
	var guess int
	for guess != params.random {
		params.loop++
		guess = util.ReadInt("Make your guess: ", inGameConditionNmber)
		if guess < params.random {
			fmt.Println("Higher !")
		} else if guess > params.random {
			fmt.Println("Lower !")
		}
	}
}

func (params *GameParameters) Result() {
	fmt.Println("HIT !")
	fmt.Printf("The limits were -> %s\n", params.mainLimit)
	fmt.Printf("You found %d in %d tries.\n", params.random, params.loop)
}
