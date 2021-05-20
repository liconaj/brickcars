package main
import (
	//"fmt"
	"strings"
	"math/rand"
	tl "github.com/liconaj/brickcars/termloop"
)


var patterns = [10]string{
	"-------x-------x-----------x-------x-----------x-------x-----------x-------------------xx------x---x-------x----x--xx------x---x-----x",
	"----x-----x----------x--------x-------x--x--x-----x----------x--------x-------x-----x-----x----------x--------x-------x---------------",
	"--------x------x--------x--xx-------------------x------x--------x--xx-------------------x------x--------x--xx------------------------x",
	"---x-------------x--------------x----x-----x--------x------x------------x----x--------------x----x---x---------x-------x--------------",
	"-------x--x-----------x----x--x----x--------------------------x----x-------x-----------x--x-----------x----x-------x-----------x--x---",
	"x---------------x-----------------------x--------x------x-----------------------x--x------------x---------------x-------x--x----------",
	"-----------x-------------------x-------------------x-------------x-----x-------------------x---------x---------------x--------x-------",
	"---x--------------x----x--------------x----x--------------x-----------------------------x---------x-------------------x---------------",
	"------------x-------------------xx------------------x-------------------xx------------------x----------x--------xx-------------x----x-",
	"x--------x---------xx--x------------x------x-----x---------xx--x------------x------x-----x--------xx--------x---------x----x-----x----",
}


type Enemy struct {
	*tl.Entity
	sprite int
	visible bool
	position float64
}

func (enemy *Enemy) Draw(screen *tl.Screen) {
	x, y := enemy.Position()
	if game.velocity < 30 {
		game.velocity += 0.01 * screen.TimeDelta()
	}
	velocity := game.velocity
	if game.boost {
		velocity = 60
	}
	if !game.lose {
		enemy.position += velocity * screen.TimeDelta()
	}
	y = int(enemy.position)
	if y % 2 != 0 {
		y -= 1
	}
	enemy.visible = true
	if y <= 0 || y >= 54 {
		enemy.visible = false
	}
	enemy.SetPosition(x, y)
	if enemy.visible {
		enemy.Entity.Draw(screen)
	}
}


func NewEnemy(x, y int) Enemy {
	sprite := assets.enemy
	enemy := Enemy{
		Entity: tl.NewEntity(x, y, len(*sprite), len((*sprite)[0])),
	}
	enemy.visible = true
	enemy.position = float64(y)
	return enemy
}


func SpawnEnemies() []*Enemy {
	var taken []int
	var enemies []*Enemy
	for i := 0; i < 5; i++ {
		n := rand.Intn(9)
		for Find(taken, n) {
			n = rand.Intn(9)
		}
		taken = append(taken, n)
		pattern := strings.Split(patterns[n], "")
		for j,p := range pattern {
			if p == "x" {
				enemy := NewEnemy(12 + i*8, -10 * len(pattern) + j*10)
				enemy.ApplyCanvas(assets.enemy)
				enemies = append(enemies, &enemy)
			}
		}
	}
	return enemies
}


func Find(slice []int, val int) (bool) {
    for _, item := range slice {
        if item == val {
            return true
        }
    }
    return false
}
