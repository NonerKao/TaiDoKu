package tdk

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

var text = []string{`０`, `１`, `２`, `３`, `４`, `５`, `６`, `７`, `８`, `９`, `Ａ`, `Ｂ`, `Ｃ`, `Ｄ`, `Ｅ`, `Ｆ`}

type IP struct {
	State   IP_STATE
	Tile    *Tile
	x       uint8
	y       uint8
	Partner int
	Token   int
	Ch      chan bool
}

type Tile struct {
	IP       *IP
	neighbor [2]*Tile
	a        [16][16]uint8
}

func NewRand(dup, upper, lower *Tile) *Tile {
	var nt Tile
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 16; i += 1 {
		for j := 0; j < 16; j += 1 {
			if dup == nil {
				nt.a[i][j] = uint8(rand.Int() % 16)
			} else {
				nt.a[i][j] = dup.a[i][j]
			}
		}
	}

	if upper != nil {
		nt.neighbor[0] = upper
		nt.neighbor[1] = lower
		upper.neighbor[1] = &nt
		lower.neighbor[0] = &nt
	} else {
		nt.neighbor[0] = &nt
		nt.neighbor[1] = &nt
	}

	return &nt
}

func (t *Tile) Print() {
	top := "╔════════╤════════╤════════╤════════╗\n"
	mid := "╟╌╌╌╌╌╌╌╌┼╌╌╌╌╌╌╌╌┼╌╌╌╌╌╌╌╌┼╌╌╌╌╌╌╌╌╢\n"
	bot := "╚════════╧════════╧════════╧════════╝\n"

	fix := color.New(color.FgYellow, color.Bold)
	nf := color.New(color.Bold)
	ip := color.New(color.BgWhite, color.FgRed, color.Bold)

	t.a[t.IP.x][t.IP.y] += 32

	fmt.Print(top)
	for i := 0; i < 16; i += 1 {
		fmt.Print(`║`)
		for j := 0; j < 16; j += 1 {
			if t.a[i][j] >= 32 {
				ip.Print(text[t.a[i][j]-32])
				t.a[t.IP.x][t.IP.y] -= 32
			} else if t.a[i][j] >= 16 {
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
