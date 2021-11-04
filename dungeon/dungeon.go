package dungeon

import (
	_ "fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type dungeonGenerator struct {
	difficulty int //1-easy|2-normal|3-hard
}
type Room struct {
	Size       size //4x6, 8x8, 5x9, etc. (only squares)
	MobAmt     int
	Directions []rune //N, S, W, E
}
type size struct {
	X int
	Y int
}

func NewDungeonGen() dungeonGenerator {
	gen := dungeonGenerator{}
	d := rand.Intn(4-1) + 1
	gen.difficulty = d
	return gen
}

func (d *dungeonGenerator) NewRoom() Room {
	r := Room{}
	dirNum := rand.Intn(5-1) + 1 //creates random number of Directions
	var dirs int
	for i := 0; i < dirNum; i++ {
		dirs = rand.Intn(5-1) + 1 //amt of Directions                           //Error on switch statement that starts at line 38,   panic: runtime error: index out of range [x] with length 0
		switch dirs {
		case 1:
			r.Directions[0] = 'N'
		case 2:
			r.Directions[1] = 'S'
		case 3:
			r.Directions[2] = 'W'
		case 4:
			r.Directions[3] = 'E'
		}
	}
	switch d.difficulty { //creates room size and amt of mobs
	case 1:
		r.Size = size{X: rand.Intn(d.difficulty*4-1) + 1, Y: rand.Intn(d.difficulty*4-1) + 1}
		r.MobAmt = rand.Intn(d.difficulty*2-1) + 1 //num = multiplier / 2
	case 2:
		r.Size = size{X: rand.Intn(d.difficulty*7-1) + 1, Y: rand.Intn(d.difficulty*7-1) + 1}
		r.MobAmt = rand.Intn(d.difficulty*4-1) + 1
	case 3:
		r.Size = size{X: rand.Intn(d.difficulty*8-1) + 1, Y: rand.Intn(d.difficulty*8-1) + 1}
		r.MobAmt = rand.Intn(d.difficulty*5-1) + 1 //higher concentration of enemies
	default:
		return Room{}
	}
	return r
}
