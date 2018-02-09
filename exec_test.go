package tdk

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func setupFile(t *testing.T) *os.File {
	fp, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}

	_, err = io.WriteString(fp, "4")
	if err != nil {
		t.Fatal(err)
	}

	_, err = fp.Seek(0, os.SEEK_SET)
	if err != nil {
		t.Fatal(err)
	}
	return fp
}

func TestExecute_IN(t *testing.T) {
	var tile Tile
	tile.IP = new(IP)

	in := setupFile(t)
	os.Stdin = in
	defer in.Close()

	tile.Execute(OP_IN)

	if tile.a[0][0] != 4 {
		t.Error("Execute() OP_IN failed.")
	}
	t.Log("Execute() OP_IN.")
}

func TestExecute_OUT(t *testing.T) {
	var tile Tile
	tile.IP = new(IP)
	tile.IP.x = 3
	tile.IP.y = 11
	tile.a[3][11] = 2
	tile.a[3][12] = 3
	tile.a[3][13] = 11

	out := setupFile(t)
	old := os.Stdout
	os.Stdout = out
	defer out.Close()

	tile.Execute(OP_OUT)

	var dest rune
	_, err := out.Seek(0, os.SEEK_SET)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Fscanf(out, "%c", &dest)

	os.Stdout = old

	if dest != '2' {
		t.Errorf("Execute() OP_OUT failed: %c.", dest)
	}
	t.Log("Execute() OP_OUT.")
}

func TestExecute_SV(t *testing.T) {
	var tile Tile
	tile.IP = new(IP)
	tile.IP.x = 14
	tile.IP.y = 13
	tile.a[14][14] = 1
	tile.a[14][15] = 1
	tile.a[15][0] = 8

	tile.Execute(OP_SV)

	if tile.a[1][1] != 8 {
		t.Errorf("Execute() OP_SV failed: %X.", tile.a[1][1])
	}
	t.Log("Execute() OP_SV.")
}

func TestExecute_LD(t *testing.T) {
	var tile Tile
	tile.IP = new(IP)
	tile.IP.x = 4
	tile.IP.y = 9
	tile.a[4][10] = 3
	tile.a[4][11] = 9
	tile.a[3][9] = 10

	tile.Execute(OP_LD)

	if tile.a[4][12] != 10 {
		t.Errorf("Execute() OP_LD failed: %X.", tile.a[4][12])
	}
	t.Log("Execute() OP_LD.")
}

func TestExecute_JMP(t *testing.T) {
	var tile Tile
	tile.IP = new(IP)
	tile.IP.x = 10
	tile.IP.y = 0
	tile.a[10][1] = 0
	tile.a[10][2] = 10

	tile.Execute(OP_JMP)

	if tile.IP.x != 0 || tile.IP.y != 10 {
		t.Errorf("Execute() OP_JMP failed: IP at (%X,%X).", tile.IP.x, tile.IP.y)
	}
	t.Log("Execute() OP_JMP.")
}

func TestExecute_JEQ(t *testing.T) {
	var tile Tile
	tile.IP = new(IP)
	tile.IP.x = 5
	tile.IP.y = 7
	tile.a[5][8] = 0
	tile.a[5][9] = 10
	tile.a[5][10] = 10
	tile.a[5][11] = 10

	tile.Execute(OP_JMP)

	if tile.IP.x != 0 || tile.IP.y != 10 {
		t.Errorf("Execute() OP_JEQ failed: IP at (%X,%X).", tile.IP.x, tile.IP.y)
	}
	t.Log("Execute() OP_JEQ.")
}

func TestExecute_JL_JLU(t *testing.T) {
	var tile Tile
	tile.IP = new(IP)
	tile.IP.x = 5
	tile.IP.y = 7
	tile.a[5][8] = 0
	tile.a[5][9] = 10
	tile.a[5][10] = 10
	tile.a[5][11] = 7

	tile.Execute(OP_JL)

	if tile.IP.x != 0 || tile.IP.y != 10 {
		t.Errorf("Execute() OP_JL failed: IP at (%X,%X).", tile.IP.x, tile.IP.y)
	}
	t.Log("Execute() OP_JL.")

	tile.a[0][11] = 2
	tile.a[0][12] = 1
	tile.a[0][13] = 2
	tile.a[0][14] = 15

	tile.Execute(OP_JLU)

	if tile.IP.x != 2 || tile.IP.y != 1 {
		t.Errorf("Execute() OP_JLU failed: IP at (%X,%X).", tile.IP.x, tile.IP.y)
	}
	t.Log("Execute() OP_JLU.")
}

func TestExecute_ALU1(t *testing.T) {
	var tile Tile
	tile.IP = new(IP)
	tile.IP.x = 4
	tile.IP.y = 9
	tile.a[4][10] = 13
	tile.a[4][11] = 1
	tile.a[4][12] = 10
	tile.a[4][13] = 7

	tile.Execute(OP_ADD)
	tile.Execute(OP_JMP)

	if tile.a[13][1] != 1 {
		t.Errorf("Execute() OP_ADD failed: %X.", tile.a[13][1])
	}
	t.Log("Execute() OP_ADD.")

	tile.a[13][2] = 2
	tile.a[13][3] = 5
	tile.a[13][4] = 2
	tile.a[13][5] = 6

	tile.Execute(OP_AND)
	tile.Execute(OP_JMP)

	if tile.a[2][5] != 2 {
		t.Errorf("Execute() OP_AND failed: %X.", tile.a[2][5])
	}
	t.Log("Execute() OP_AND.")

	tile.a[2][6] = 11
	tile.a[2][7] = 11
	tile.a[2][8] = 10
	tile.a[2][9] = 7

	tile.Execute(OP_OR)
	tile.Execute(OP_JMP)

	if tile.a[11][11] != 15 {
		t.Errorf("Execute() OP_OR failed: %X.", tile.a[11][11])
	}
	t.Log("Execute() OP_OR.")

	tile.a[11][12] = 0
	tile.a[11][13] = 12
	tile.a[11][14] = 11
	tile.a[11][15] = 13

	tile.Execute(OP_XOR)
	tile.Execute(OP_JMP)

	if tile.a[0][12] != 6 {
		t.Errorf("Execute() OP_XOR failed: %X.", tile.a[0][12])
	}
	t.Log("Execute() OP_XOR.")
}

func TestExecute_ALU2(t *testing.T) {
	var tile Tile
	tile.IP = new(IP)
	tile.IP.x = 4
	tile.IP.y = 9
	tile.a[13][1] = 15
	tile.a[4][10] = 13
	tile.a[4][11] = 1
	tile.a[4][12] = 10

	tile.Execute(OP_SHL)

	if tile.a[13][1] != 12 {
		t.Errorf("Execute() OP_SHL failed: %X.", tile.a[13][1])
	}
	t.Log("Execute() OP_SHL.")

	tile.a[4][10] = 13
	tile.a[4][11] = 1
	tile.a[4][12] = 9

	tile.Execute(OP_SHRL)

	if tile.a[13][1] != 6 {
		t.Errorf("Execute() OP_SHRL failed: %X.", tile.a[13][1])
	}
	t.Log("Execute() OP_SHRL.")

	tile.a[13][1] = 15
	tile.a[4][10] = 13
	tile.a[4][11] = 1
	tile.a[4][12] = 7

	tile.Execute(OP_SHRA)

	if tile.a[13][1] != 15 {
		t.Errorf("Execute() OP_SHRA failed: %X.", tile.a[13][1])
	}
	t.Log("Execute() OP_SHRA.")
}

func TestExecute_META_MOVE(t *testing.T) {
	var tile1 Tile
	var tile2 Tile

	tile1.IP = new(IP)
	tile1.IP.x = 6
	tile1.IP.y = 1
	tile1.a[6][2] = uint8(META_MOVE)
	tile1.a[6][3] = 13
	tile1.a[6][4] = 1
	tile1.a[6][5] = 3
	tile1.a[1][3] = 15

	tile2.IP = nil
	tile1.neighbor[1] = &tile2
	tile1.Execute(OP_META)

	if tile2.IP == nil || tile1.IP != nil {
		t.Errorf("Execute() META_MOVE failed.")
	}
	t.Log("Execute() META_MOVE.")
}

func TestExecute_META_NEW_DELETE(t *testing.T) {
	tile1 := new(Tile)

	ip := new(IP)
	tile1.IP = ip
	ip.Tile = tile1
	tile1.IP.x = 5
	tile1.IP.y = 10
	tile1.a[5][11] = uint8(META_NEW_RAND)
	tile1.a[5][12] = 11

	tile2 := NewRand(nil, tile1, tile1)
	tile3 := NewRand(nil, tile1, tile2)
	ip.Execute(OP_META)

	if tile1.neighbor[1].neighbor[1] != tile3 || tile3.neighbor[0].neighbor[0] != tile1 {
		t.Errorf("Execute() META_NEW_RAND failed.")
	}
	t.Log("Execute() META_NEW_RAND.")

	ip.Tile.a[5][11] = uint8(META_MOVE)
	ip.Tile.a[5][12] = 8
	ip.Tile.a[5][13] = 5
	ip.Tile.a[5][14] = 10
	ip.Execute(OP_META)
	ip.Tile.a[5][11] = uint8(META_DELETE)
	ip.Tile.a[5][12] = 7
	ip.Execute(OP_META)
	if ip.Tile != tile2 {
		t.Errorf("Execute() META_DELETE failed.")
	}
	t.Log("Execute() META_DELETE.")
}

func TestExecute_META_NEW_DUP(t *testing.T) {
	tile1 := new(Tile)

	tile1.IP = new(IP)
	tile1.IP.x = 3
	tile1.IP.y = 8
	tile1.a[3][9] = uint8(META_NEW_DUP)
	tile1.a[3][10] = 2

	var tile2 *Tile
	tile2 = NewRand(tile1, tile1, tile1)
	tile1.Execute(OP_META)

	if tile2.a[3][10] != 2 || tile1.neighbor[1] != tile2 {
		t.Errorf("Execute() META_NEW_DUP failed: %X.", tile2.a[3][10])
	}
	t.Log("Execute() META_NEW_DUP.")
}

func TestExecute_META_SHUFFLE(t *testing.T) {
	tile1 := new(Tile)

	tile1.IP = new(IP)
	tile1.IP.x = 3
	tile1.IP.y = 8
	tile1.a[3][9] = uint8(META_SHUFFLE)
	tile1.a[3][10] = 2
	tile1.a[3][11] = 3

	tile1.Execute(OP_META)

	if tile1.a[3][10] != 3 || tile1.a[3][11] != 2 {
		t.Errorf("Execute() META_SHUFFLE failed: %X, %X.", tile1.a[3][10], tile1.a[3][11])
	}
	t.Log("Execute() META_SHUFFLE.")
}

func TestExecute_META_COPY(t *testing.T) {
	ip := new(IP)
	tile1 := new(Tile)

	ip.Tile = tile1
	tile1.IP = ip
	tile1.IP.x = 5
	tile1.IP.y = 10
	tile1.a[5][11] = uint8(META_COPY)
	// from top, region 2
	tile1.a[5][12] = 6
	tile1.a[5][13] = 2

	tile2 := NewRand(nil, tile1, tile1)
	tile3 := NewRand(nil, tile2, tile1)
	tile3.a[0][8] = 15
	tile3.a[0][11] = 3
	tile3.a[3][11] = 2
	ip.Execute(OP_META)

	if tile1.a[0][8] != 15 || tile1.a[0][11] != 3 || tile1.a[3][11] != 2 {
		t.Errorf("Execute() META_COPY FROM %X,%X,%X.", tile1.a[0][8], tile1.a[0][11], tile1.a[3][11])
	}
	t.Log("Execute() META_COPY FROM.")

	tile1.a[5][11] = uint8(META_COPY)
	// to bottom, row 13
	tile1.a[5][12] = 13
	tile1.a[5][13] = 13
	tile1.a[13][4] = 15
	tile1.a[13][7] = 3
	tile1.a[13][11] = 2
	ip.Execute(OP_META)

	if tile2.a[13][4] != 15 || tile2.a[13][7] != 3 || tile2.a[13][11] != 2 {
		t.Errorf("Execute() META_COPY TO %X,%X,%X.", tile2.a[13][4], tile2.a[13][7], tile2.a[13][11])
	}
	t.Log("Execute() META_COPY TO.")
}

func TestExecute_META_HALT(t *testing.T) {
	ip := new(IP)
	tile := new(Tile)

	ip.Tile = tile
	tile.IP = ip
	tile.IP.x = 4
	tile.IP.y = 6
	tile.a[4][7] = uint8(META_HALT)
	ip.Execute(OP_META)

	if tile.IP != nil || ip.Tile != nil {
		t.Errorf("Execute() META_HALT.")
	}
	t.Log("Execute() META_HALT.")
}
