package main

import (
	"fmt"
	"runtime"
	tl "github.com/liconaj/brickcars/termloop"
)

var game Game
var assets Assets

func main() {
	assets = loadAssets()
	game = NewGame()
	game.SetState("init")
	game.Start()
}

type Assets struct {
	skin *tl.Canvas
	background *tl.Canvas

	title *tl.Canvas
	over *tl.Canvas

	player *tl.Canvas
	enemy *tl.Canvas

	explosion [5]*tl.Canvas
	numbers [10]*tl.Canvas
}


func loadAssets() Assets {
	a := Assets{}
	a.title = tl.BackgroundCanvasFromFile("data/title.png")
	a.player = tl.BackgroundCanvasFromFile("data/player.png")
	a.enemy = tl.BackgroundCanvasFromFile("data/enemy.png")
	a.over = tl.BackgroundCanvasFromFile("data/over.png")
	for i := 0; i < 10; i++ {
		image := tl.BackgroundCanvasFromFile(fmt.Sprintf("data/numbers/%d.png", i))
			a.numbers[i] = image
	}
	if runtime.GOOS == "windows" {
		a.skin = tl.BackgroundCanvasFromFile("data/skin_windows.png")
		a.background = tl.BackgroundCanvasFromFile("data/background_windows.png")
		for i := 0; i < 5; i++ {
			effect := tl.BackgroundCanvasFromFile(fmt.Sprintf("data/explosion/%d_windows.png", i))
			a.explosion[i] = effect
		}
	} else {
		a.skin = tl.BackgroundCanvasFromFile("data/skin.png")
		a.background = tl.BackgroundCanvasFromFile("data/background.png")
		for i := 0; i < 5; i++ {
			effect := tl.BackgroundCanvasFromFile(fmt.Sprintf("data/explosion/%d.png", i))
			a.explosion[i] = effect
		}
	}
	return a
}
