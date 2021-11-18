//A Text/Console based Dungeon Game for the 2021 Game Off!
//go run main.go

package main

import (
	"fmt"

	"strconv"
	"strings"
	"time"
	"math/rand"
	dun "github.com/csharpdf/Gatoraid/dungeon"
	//"github.com/csharpdf/Gatoraid/visual"
)

var compass string
var d = dun.NewDungeonGen(3)
var resp string
var dirs []string
var difficulty int
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
	fmt.Println("Welcome to Gatoraid!\nThis is a infinite text-based dungeon crawler made for the 2021 Game Off!\ngeneral commands are 'use [item]' and 'proceed'\n\n")
	//write("What difficulty would you like to play? You can choose 'easy' mode, 'normal' mode, and 'hard' mode! Hint: hard mode is way too hard") //make room counter
	fmt.Print("> ")
	getD()
}
func getD() {
	fmt.Scanln(&resp)
	switch strings.ToLower(resp) {
	case "easy":
		difficulty = 1
	case "normal":
		difficulty = 2
	case "hard":
		difficulty = 3
	default:
		write("\nInvalid gamemode given. Gamemodes are 'easy', 'normal', and 'hard'.")
		fmt.Print("> ")
		getD()
	}
}
func parseMsg() {
	rand.Seed(time.Now().UnixNano())
	fmt.Scanln(&resp)
	switch strings.ToLower(resp) {
	case "proceed":
		write("You proceed down the seemingly infinite hall.")
		y := rand.Intn(100+1) - 1
		var x int;
		switch difficulty {
		case 1:
			x = 20
		case 2:
			x = 45
		case 3:
			x = 80
		}
		if (y ) //TODO: make rand enemy generator based on difficulty
	}
}
func main() {
	//write("You walk down a stairwell that reminds you of your past. This is not the first dungeon you've explored, and it will certainly not be the last. You stumble into a dark, long hall as the doors behind you close, not knowing where the first monsters are hiding.")
	//! removed rooms and basically everything to actually make it "simple" 
	fmt.Println("You can choose to proceed, or check your inventory.\n> ")
}
