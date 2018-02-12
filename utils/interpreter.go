package main

import (
	tdk "github.com/NonerKao/TaiDoKu"
	"os"
)

func main() {
	var p *tdk.TDKProgram
	if len(os.Args) != 1 {
		p = tdk.InitProgram(os.Args[1])
	} else {
		p = tdk.InitRandProgram()
	}

	p.Run()
}
