package tdk

import (
	"fmt"
	"os"
)

type TDKProgram struct {
	Tile []*Tile
	IP   []*IP
}

func InitProgram(fn string) *TDKProgram {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var nt uint64

	fmt.Fscan(f, &nt)
	tdk := TDKProgram{
		Tile: make([]*Tile, nt),
		IP:   make([]*IP, 0),
	}

	for k := 0; k < int(nt); k++ {
		tdk.Tile[k] = Rand()

		var hasip int
		fmt.Fscan(f, &hasip)
		if hasip == 1 {
			ip := IP{
				Tile: nil,
				x:    0,
				y:    0,
			}
			tdk.IP = append(tdk.IP, &ip)
			ip.Tile = tdk.Tile[k]
			tdk.Tile[k].IP = &ip
		}

		for i := 0; i < 16; i += 1 {
			for j := 0; j < 16; j += 1 {
				fmt.Fscanf(f, "%X", &tdk.Tile[k].a[i][j])
			}
		}
	}

	return &tdk
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
