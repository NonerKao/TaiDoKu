package main

import (
	tdk "github.com/NonerKao/TaiDoKu/TaiDoKu"
)

func main() {
	p := tdk.InitRandProgram()

	p.Tile[0].Print()
	p.Tile[0].Refine(0, 0)
	p.Tile[0].Print()
}
