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
		if r.Directions[2] != 0 && i == (r.Size.Y-2)/2 {
			fmt.Print("|")
			for i := 0; i < r.Size.X-2; i++ {
				if i < r.Size.X-2 {
					fmt.Print("  ")
				}
			}
			if r.Directions[1] == 0 {
				fmt.Print(" *")
			}
			fmt.Print("\n")
		}
		if r.Directions[1] != 0 && i == (r.Size.Y-2)/2 {
			if r.Directions[2] == 0 {
				fmt.Print(" *")
			}
			for i := 0; i < r.Size.X-2; i++ {
				if i < r.Size.X-2 {
					fmt.Print("  ")
				}
			}
			fmt.Print("  |")
			fmt.Print("\n")
		}
	}
	for i := 0; i < r.Size.X; i++ { //writes last row
		if i < r.Size.X {
			fmt.Print("* ")
		} else if i == r.Size.X {
			fmt.Print("*")
		}
	}
}
