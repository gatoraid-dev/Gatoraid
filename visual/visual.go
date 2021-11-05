package visual

import (
	"fmt"

	dun "github.com/csharpdf/Gatoraid/dungeon"
)

func WriteMap(r *dun.Room) {
	for i := 0; i < r.Size.X; i++ { //writes first row
		if i < r.Size.X {
			fmt.Print("* ")
		} else if i == r.Size.X {
			fmt.Print("*")
		}
	}
	fmt.Print("\n")
	for i := 0; i < r.Size.Y-2; i++ { //writes columns
		fmt.Print("*")
		for i := 0; i < r.Size.X-2; i++ {
			if i < r.Size.X-2 {
				fmt.Print("  ")
			} else if i == r.Size.X-2 {
				fmt.Print(" ")
			}
		}
		fmt.Print(" *")
		fmt.Print("\n")

	}
	for i := 0; i < r.Size.X; i++ { //writes last row
		if i < r.Size.X {
			fmt.Print("* ")
		} else if i == r.Size.X {
			fmt.Print("*")
		}
	}
}

func WriteMapWithChar(r dun.Room) { //used for moving the char (rewriting/updating the map)
	for i := 0; i < r.Size.X; i++ { //writes first row
		if i < r.Size.X {
			fmt.Print("* ")
		} else if i == r.Size.X {
			fmt.Print("*")
		}
	}
	fmt.Print("\n")
	for i := 0; i < r.Size.Y-2; i++ { //writes columns
		fmt.Print("*")
		for i := 0; i < r.Size.X-2; i++ {
			if i < r.Size.X-2 {
				fmt.Print("  ")
			} else if i == r.Size.X-2 {
				fmt.Print(" ")
			}
		}
		fmt.Print(" *")
		fmt.Print("\n")

	}
	for i := 0; i < r.Size.X; i++ { //writes last row
		if i < r.Size.X {
			fmt.Print("* ")
		} else if i == r.Size.X {
			fmt.Print("*")
		}
	}
}
