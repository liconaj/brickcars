package main
import (
	"time"
	"math/rand"
	tl "github.com/liconaj/brickcars/termloop"
)

type Game struct {
	*tl.Game
	frame Frame
	velocity float64
	boost bool
	lose bool
	state string
	distance float64
}


func NewGame() Game {
	rand.Seed(time.Now().UnixNano())
	game := Game{}
	game.distance = 0
	game.lose = false
	game.boost = false
	game.velocity = 10
	game.Game = tl.NewGame()
	game.Game.Screen().SetFps(30)
	game.Game.Screen().EnablePixelMode()
	return game
}

func (game *Game) Start() {
	game.Game.Start()
}

func (game *Game) SetState(state string) {
	game.state = state
	if state == "init" {
		game.frame = WelcomeScreen()
	} else if state == "over" {
		game.frame = OverScreen()
	} else if state == "play" {
		game.frame = PlayGame()
	}
	game.Game.Screen().SetLevel(&game.frame)
}


func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
