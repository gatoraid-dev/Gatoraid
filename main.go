//A Text/Console based Dungeon Game for the 2021 Game Off!
//go run main.go

package main

import (
	"bufio"
	"fmt"
	"os"

	dun "github.com/csharpdf/Gatoraid/dungeon"
	"github.com/csharpdf/Gatoraid/visual"
)

var input = bufio.NewReader(os.Stdin)
var compass string
var d = dun.NewDungeonGen()

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
}

func main() {
	fmt.Print("Welcome to Gatoraid!\nThis is a text-based dungeon made for the 2021 Game Off!\n\nStarting adventure...\n\n\n")
	room := d.NewRoom()
	visual.WriteMap(room)
}
