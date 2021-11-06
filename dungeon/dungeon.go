package dungeon

import (
	_ "fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Coordinates size
type dungeonGenerator struct {
	Difficulty int //1-easy|2-normal|3-hard
}
type Room struct {
	Size       size //4x6, 8x8, 5x9, etc. (only squares)
	MobAmt     int
	Directions [4]rune //N, E, W, S (in that order)
}
type size struct {
	X int
	Y int
}

func NewDungeonGen(difficulty int) dungeonGenerator {
	gen := dungeonGenerator{}
	gen.Difficulty = difficulty
	return gen
}
func (d *dungeonGenerator) NewRoom() Room {
	r := Room{}
	dirNum := rand.Intn(5-1) + 1 //creates random number of Directions
	var dirs int
	for i := 0; i < dirNum; i++ {
		dirs = rand.Intn(5-1) + 1 //amt of Directions
		switch dirs {
		case 1:
			if r.Directions[0] == rune(0) {
				r.Directions[0] = 'N'
				//append(r.Directions, 'N')
			}
		case 2:
			if r.Directions[1] == rune(0) {
				r.Directions[1] = 'E'
				//append(r.Directions, 'N')
			}
		case 3:
			if r.Directions[2] == rune(0) {
				r.Directions[2] = 'W'
				//append(r.Directions, 'N')
			}
		case 4:
			if r.Directions[3] == rune(0) {
				r.Directions[3] = 'S'
				//append(r.Directions, 'N')
			}
		}
	}
	switch d.Difficulty { //creates room size and amt of mobs
	case 1:
		mu := 16 //multiplier
		r.Size = size{X: rand.Intn((d.Difficulty*mu)-1) + 1, Y: rand.Intn((d.Difficulty*mu)-1) + 1}
		r.MobAmt = rand.Intn((d.Difficulty*(mu/2))-1) + 1 //num = multiplier / 2
		for (r.Size.X < 5 || r.Size.Y < 5) || (r.Size.X < 5 && r.Size.Y < 5) || (r.Size.X >= r.Size.Y*4) || (r.Size.Y >= r.Size.X*4) {
			r.Size = size{X: rand.Intn((d.Difficulty*mu)-1) + 1, Y: rand.Intn((d.Difficulty*mu)-1) + 1}
			r.MobAmt = rand.Intn((d.Difficulty*(mu/2))-1) + 1
		}
	case 2:
		mu := 26
		r.Size = size{X: rand.Intn((d.Difficulty*mu)-1) + 1, Y: rand.Intn((d.Difficulty*mu)-1) + 1}
		r.MobAmt = rand.Intn((d.Difficulty*(mu/2))-1) + 1
		for (r.Size.X < 5 || r.Size.Y < 5) || (r.Size.X < 5 && r.Size.Y < 5) {
			r.Size = size{X: rand.Intn((d.Difficulty*mu)-1) + 1, Y: rand.Intn((d.Difficulty*mu)-1) + 1}
			r.MobAmt = rand.Intn((d.Difficulty*(mu/2))-1) + 1
		}
	case 3:
		mu := 36
		r.Size = size{X: rand.Intn((d.Difficulty*mu)-1) + 1, Y: rand.Intn((d.Difficulty*mu)-1) + 1}
		r.MobAmt = rand.Intn((d.Difficulty*(mu/2))-1) + 1 //higher concentration of enemies
		for (r.Size.X < 5 || r.Size.Y < 5) || (r.Size.X < 5 && r.Size.Y < 5) {
			r.Size = size{X: rand.Intn((d.Difficulty*mu)-1) + 1, Y: rand.Intn((d.Difficulty*mu)-1) + 1}
			r.MobAmt = rand.Intn((d.Difficulty*(mu/2))-1) + 1
		}
	default:
		return Room{}
	}
	return r
}
