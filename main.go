//A Text/Console based Dungeon Game for the 2021 Game Off!
//go run main.go

package main

import (
	"fmt"

	"strings"
	"time"

	dun "github.com/csharpdf/Gatoraid/dungeon"
	"github.com/csharpdf/Gatoraid/visual"
)

var compass string
var d = dun.NewDungeonGen(3)
var resp string

func write(t string) {
	x := []rune(t)
	for i := 0; i < len(x); i++ {
		if i == len(x)-1 {
			fmt.Println(string(x[i]))
		} else {
			fmt.Print(string(x[i]))
		}
		time.Sleep(time.Millisecond * 50)
	}
}
func init() {
	compass =
		`        _________________
                |        |        |
				| ~ N ~  |  ~ E ~ |
				|        |        |
				|--------|--------|
				|        |        |
				| ~ W ~  |  ~ S ~ |
				|        |        |
				-------------------  `
	fmt.Println("Welcome to Gatoraid!\nThis is a text-based dungeon made for the 2021 Game Off!")
	write("\nStarting adventure...\n\n")
	write("What difficulty would you like to play? You can choose 'easy' mode, 'normal' mode, and 'hard' mode! Hint: hard mode is way too hard") //make room counter
	fmt.Print("> ")
	getD()
}
func getD() {
	fmt.Scanln(&resp)
	switch strings.ToLower(resp) {
	case "easy":
		d.Difficulty = 1
	case "normal":
		d.Difficulty = 2
	case "hard":
		d.Difficulty = 3
	default:
		write("\nInvalid gamemode given. Gamemodes are 'easy', 'normal', and 'hard'.")
		fmt.Print("> ")
		getD()
	}
}
func main() {
	room := d.NewRoom()
	write("You walk down a stairwell that reminds you of your past. This is not the first dungeon you've explored, and it will certainly not be the last. You stumble into a dark room as the doors behind you close, not knowing where the first monsters are hiding.")
	write("\nCreating map of room...\n")
	//fmt.Println(room.Size.X, "x | ", room.Size.Y, "y\n\n", d.Difficulty)
	visual.WriteMap(&room)
}
