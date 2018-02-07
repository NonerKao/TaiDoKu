package main
import (
	tdk "github.com/NonerKao/TaiDoKu/TaiDoKu"
	"os"
)
func main() {
	var p *tdk.TDKProgram
	if len(os.Args) != 1 {
		p = tdk.InitProgram(os.Args[1])
	} else {
		p = tdk.InitRandProgram()
	}
	p.Tile[0].Print()
	op := p.Tile[0].OPLookup()
	p.Tile[0].Print()
	p.Tile[0].Execute(op)
	p.Tile[0].Print()
	p.Tile[0].Refine()
	p.Tile[0].Print()
}
