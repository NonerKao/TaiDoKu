package tdk

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

var text = []string{`０`, `１`, `２`, `３`, `４`, `５`, `６`, `７`, `８`, `９`, `Ａ`, `Ｂ`, `Ｃ`, `Ｄ`, `Ｅ`, `Ｆ`}

type IP struct {
	Tile *Tile
	x    uint8
	y    uint8
}

type Tile struct {
	IP *IP
	a  [16][16]uint8
}

func Rand() *Tile {
	var t Tile
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 16; i += 1 {
		for j := 0; j < 16; j += 1 {
			t.a[i][j] = uint8(rand.Int() % 16)
		}
	}

	return &t
}

func (t *Tile) Print() {
	top := "╔════════╤════════╤════════╤════════╗\n"
	mid := "╟╌╌╌╌╌╌╌╌┼╌╌╌╌╌╌╌╌┼╌╌╌╌╌╌╌╌┼╌╌╌╌╌╌╌╌╢\n"
	bot := "╚════════╧════════╧════════╧════════╝\n"

	fix := color.New(color.FgYellow, color.Bold)
	nf := color.New(color.Bold)

	fmt.Print(top)
	for i := 0; i < 16; i += 1 {
		fmt.Print(`║`)
		for j := 0; j < 16; j += 1 {
			if t.a[i][j] >= 16 {
				fix.Print(text[t.a[i][j]-16])
			} else {
				nf.Print(text[t.a[i][j]])
			}
			if j%4 == 3 && j != 15 {
				fmt.Print(`┆`)
			}
		}
		fmt.Print("║\n")
		if i%4 == 3 && i != 15 {
			fmt.Print(mid)
		}
	}

	fmt.Print(bot)
}
