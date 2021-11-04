package visual

import (
	"fmt"

	dun "github.com/csharpdf/Gatoraid/dungeon"
)

func WriteMap(r dun.Room) {
	fmt.Println(r.Size.X, " | ", r.Size.Y)
	for i := 0; i <= r.Size.X; i++ { //writes first row
		fmt.Print("#")
	}
	for i := 0; i <= r.Size.Y; i++ { //writes columns
		fmt.Print("#")
		for i := 0; i <= r.Size.X-2; i++ {
			fmt.Print(" ")
		}
		fmt.Print("#")
	}
	for i := 0; i <= r.Size.X; i++ { //writes last row
		fmt.Print("#")
	}
}

func WriteMapWithChar(r dun.Room) { //used for moving the char (rewriting/updating the map)
	for i := 0; i <= r.Size.X; i++ { //writes first row
		fmt.Print("#")
	}
	for i := 0; i <= r.Size.Y; i++ { //writes columns
		if i == r.Size.Y-1 {
			fmt.Print("#")
			for i := 0; i <= r.Size.X-2; i++ {
				if i == r.Size.X/2 {
					fmt.Print("%")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Print("#")
		} else {
			fmt.Print("#")
			for i := 0; i <= r.Size.X-2; i++ {
				if i == r.Size.X/2 {
					fmt.Print("%")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Print("#")
		}
	}
	for i := 0; i <= r.Size.X; i++ { //writes last row
		fmt.Print("#")
	}
}
