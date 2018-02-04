package tdk

import (
	"fmt"
	"math/rand"
	"time"
)

var text = []string{`０`, `１`, `２`, `３`, `４`, `５`, `６`, `７`, `８`, `９`, `Ａ`, `Ｂ`, `Ｃ`, `Ｄ`, `Ｅ`, `Ｆ`}

type IP struct {
	Tile *Tile
	x    uint8
	y    uint8
}

type Tile struct {
	A  [16][16]uint8
	IP *IP
}

type TDKProgram struct {
	Tile []*Tile
	IP   []*IP
}

func InitRandProgram() *TDKProgram {
	tdk := TDKProgram{
		Tile: make([]*Tile, 1),
		IP:   make([]*IP, 1),
	}
	ip := IP{
		Tile: Rand(),
		x:    0,
		y:    0,
	}

	tdk.IP[0] = &ip
	tdk.Tile[0] = tdk.IP[0].Tile
	tdk.Tile[0].IP = &ip
	return &tdk
}

func Rand() *Tile {
	var t Tile
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 16; i += 1 {
		for j := 0; j < 16; j += 1 {
			t.A[i][j] = uint8(rand.Int() % 16)
		}
	}

	return &t
}

func (t *Tile) Print() {
	top := "╔════════╤════════╤════════╤════════╗\n"
	mid := "╟╌╌╌╌╌╌╌╌┼╌╌╌╌╌╌╌╌┼╌╌╌╌╌╌╌╌┼╌╌╌╌╌╌╌╌╢\n"
	bot := "╚════════╧════════╧════════╧════════╝\n"

	fmt.Print(top)

	for i := 0; i < 16; i += 1 {
		fmt.Print(`║`)
		for j := 0; j < 16; j += 1 {
			fmt.Print(text[t.A[i][j]])
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
