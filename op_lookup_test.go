package tdk

import "testing"

func TestOPLookUp1(t *testing.T) {
	var tile Tile
	tile.IP = new(IP)

	op := tile.OPLookup()
	if op != OP_IN {
		t.Error("OPLookup() on null tile NOT return OP_IN.")
	}
	t.Log("OPLookup() on null tile returns OP_IN.")
}

func TestOPLookUp2(t *testing.T) {
	var tile Tile
	tile.IP = new(IP)
	tile.IP.x = 12
	tile.IP.y = 4
	tile.a[12][4] = 15
	tile.a[14][6] = 15

	op := tile.OPLookup()
	if op != OP_META {
		t.Error("OPLookup() NOT return OP_META.")
	} else if tile.IP.x != 14 || tile.IP.y != 6 {
		t.Errorf("Wrong IP position (%X,%X) after OPLookup().", tile.IP.x, tile.IP.y)
	}
	t.Log("OPLookup() succeeds.")
}
