package main
import (
	"strconv"
	tl "github.com/liconaj/brickcars/termloop"
)


func PlayGame() Frame {
	x, y := 28, 42
	sprite := assets.player
	player := Player{
		Entity: tl.NewEntity(x, y, len(*sprite), len((*sprite)[0])),
	}
	play := Frame{
		BaseLevel: tl.NewBaseLevel(tl.Cell{}),
		player: player,
	}
	play.skin = NewImage(assets.skin, 0, 0)
	player.ApplyCanvas(sprite)
	play.AddEntity(NewImage(assets.background, 10, 10))
	play.AddEntity(&player)
	play.AddEntity(play.skin)
	return play
}


func WelcomeScreen() Frame {
	title := Frame{
		BaseLevel: tl.NewBaseLevel(tl.Cell{}),
	}
	title.AddEntity(NewImage(assets.background, 10, 10))
	title.AddEntity(NewImage(assets.title, 10, 10))
	title.AddEntity(NewImage(assets.skin, 0, 0))
	return title
}


func OverScreen() Frame {
	over := Frame{
		BaseLevel: tl.NewBaseLevel(tl.Cell{}),
	}
	over.AddEntity(NewImage(assets.background, 10, 10))
	over.AddEntity(NewImage(assets.over, 10, 10))
	over.AddEntity(NewImage(assets.skin, 0, 0))
	ShowScore(22, 41, &over) //22, 41
	return over
}


type Frame struct {
	*tl.BaseLevel
	skin *Image
	player Player
}


func (frame *Frame) Draw(screen *tl.Screen) {
	offsetx, offsety := frame.Offset()
	frameWidth, frameHeight := len(*assets.skin), len((*assets.skin)[0])
	screenWidth, screenHeight := screen.Size()
	if screenWidth > frameWidth {
		offsetx = (screenWidth-frameWidth)/2
	}
	if screenHeight > frameHeight {
		offsety = (screenHeight-frameHeight)/2
	}
	frame.SetOffset(offsetx, offsety)
	frame.BaseLevel.Draw(screen)
}


func (frame *Frame) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		if game.state != "play"{
			if  event.Key == tl.KeySpace || event.Key == tl.KeyEnter {
				game = NewGame()
				game.SetState("play")
			}
		} else {
			if !game.lose {
				step := 8
				x, y := frame.player.Position()
				switch event.Key {
				case tl.KeyArrowRight:
					if (x + step <= 44) {x += step}
				case tl.KeyArrowLeft:
					if (x - step >= 11) {x -= step}
				}
				if event.Key == tl.KeySpace || event.Key == tl.KeyArrowUp {
					game.boost = true
				} else {
					game.boost = false
				}
				frame.player.SetPosition(x, y)
			}
		}
	}
	frame.BaseLevel.Tick(event)
}


func ShowScore(x, y int, frame *Frame) {
	width := 3
	digits, limit := 4, 9999
	padding := 2

	score := int(game.distance/10.0)
	if score > limit {score = limit}
	textscore := strconv.Itoa(score)

	length := len(textscore)

	var i int
	for i = 0; i < digits - length; i++ {
		// zeros
		xi := x + i * (padding + width)
		frame.AddEntity(NewImage(assets.numbers[0], xi, y))
	}
	x += i * (padding + width)
	for i = 0; i < length; i++ {
		// values
		xi := x + i * (padding + width)
		number,_ := strconv.Atoi(string(textscore[i]))
		frame.AddEntity(NewImage(assets.numbers[number], xi, y))
	}
}
