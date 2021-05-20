package main
import tl "github.com/liconaj/brickcars/termloop"

type Player struct {
	*tl.Entity
	enemies []*Enemy
	timer float64
	prevtime float64
}


func (player *Player) Draw(screen *tl.Screen) {
	if game.lose {
		player.Explode(screen)
	} else {
		player.Update(screen)
	}
	player.Entity.Draw(screen)
}


func (player *Player) Update(screen *tl.Screen) {
	var newlist []*Enemy
	var distance float64
	if game.boost {
		distance = screen.TimeDelta() * 2
	} else {
		distance = screen.TimeDelta()
	}
	game.distance += distance
	for _,e := range player.enemies {
		_,y := e.Position()
		if y <= 54 {
			newlist = append(newlist, e)
		} else {
			game.frame.RemoveEntity(e)
		}
	}
	player.enemies = newlist
	if len(player.enemies) < 2 {
		enemies := SpawnEnemies()
		for _,e := range enemies {
			player.enemies = append(player.enemies, e)
			game.frame.AddEntity(e)
		}
		// Set skin in top of enemies
		game.frame.AddEntity(game.frame.skin)
	}
}


func (player *Player) Explode(screen *tl.Screen) {
	x, y := player.Position()
	player.prevtime = player.timer
	player.timer += screen.TimeDelta() * 15
	if player.timer < 8 {
		i := int(player.timer)
		if int(player.prevtime) != i {
			effect := NewImage(assets.explosion[i%5], x-2, y-2)
			game.frame.AddEntity(effect)
			game.frame.AddEntity(game.frame.skin)
		}
	} else {
		game.SetState("over")
	}
}


func (player *Player) Collide(collision tl.Physical) {
	if _, ok := collision.(*Enemy); ok {
		game.lose = true
	}
}
