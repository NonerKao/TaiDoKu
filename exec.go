package tdk

import (
	"bufio"
	"fmt"
	"os"
)

type TDK_OP uint8

const (
	OP_IN   TDK_OP = 0
	OP_OUT  TDK_OP = 1
	OP_SV   TDK_OP = 2
	OP_LD   TDK_OP = 3
	OP_JMP  TDK_OP = 4
	OP_JEQ  TDK_OP = 5
	OP_JL   TDK_OP = 6
	OP_JLU  TDK_OP = 7
	OP_ADD  TDK_OP = 8
	OP_AND  TDK_OP = 9
	OP_OR   TDK_OP = 10
	OP_XOR  TDK_OP = 11
	OP_SHL  TDK_OP = 12
	OP_SHRL TDK_OP = 13
	OP_SHRA TDK_OP = 14
	OP_META TDK_OP = 15
)

type TDK_META_OP uint8

const (
	META_MOVE      TDK_META_OP = 0
	META_NEW_RAND  TDK_META_OP = 1
	META_NEW_DUP   TDK_META_OP = 2
	META_DELETE    TDK_META_OP = 3
	META_FORK      TDK_META_OP = 4
	META_JOIN      TDK_META_OP = 5
	META_SYNC      TDK_META_OP = 6
	META_SHUFFLE   TDK_META_OP = 7
	META_COPY_FROM TDK_META_OP = 8
	META_COPY_TO   TDK_META_OP = 9
	META_HALT      TDK_META_OP = 15
)

func getHex() uint8 {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	h := uint8(input[0])
	if h >= uint8('0') && h <= uint8('9') {
		return h - uint8('0')
	} else if h >= uint8('A') && h <= uint8('F') {
		return h - uint8('A') + 10
	} else if h >= uint8('a') && h <= uint8('f') {
		return h - uint8('a') + 10
	}

	panic("Invalid input!")
}

func putHex(h uint8) {
	if h >= 0 && h <= 9 {
		fmt.Printf("%c", h+uint8('0'))
	} else {
		fmt.Printf("%c", h-10+uint8('A'))
	}
}

func (t *Tile) Execute(op TDK_OP) {
	x := t.a[(t.IP.x+(t.IP.y+1)/16)%16][(t.IP.y+1)%16] % 16
	y := t.a[(t.IP.x+(t.IP.y+2)/16)%16][(t.IP.y+2)%16] % 16
	op3 := t.a[(t.IP.x+(t.IP.y+3)/16)%16][(t.IP.y+3)%16] % 16
	op4 := t.a[(t.IP.x+(t.IP.y+4)/16)%16][(t.IP.y+4)%16] % 16
	switch op {
	case OP_IN:
		t.a[x][y] = getHex()
	case OP_OUT:
		putHex(t.a[x][y])
	case OP_SV:
		t.a[x][y] = op3
	case OP_LD:
		t.a[(t.IP.x+(y+3)/16)%16][(t.IP.y+3)%16] = t.a[x][y]
	case OP_JMP:
		t.IP.x = x
		t.IP.y = y
	case OP_JEQ:
		if op3 == op4 {
			t.IP.x = x
			t.IP.y = y
		}
	case OP_JL:
		sop3 := (op3 + 8) % 16
		sop4 := (op4 + 8) % 16
		if sop3 < sop4 {
			t.IP.x = x
			t.IP.y = y
		}
	case OP_JLU:
		if op3 < op4 {
			t.IP.x = x
			t.IP.y = y
		}
	case OP_ADD:
		t.a[x][y] = (op3 + op4) % 16
	case OP_AND:
		t.a[x][y] = op3 & op4
	case OP_OR:
		t.a[x][y] = op3 | op4
	case OP_XOR:
		t.a[x][y] = op3 ^ op4
	case OP_SHL:
		t.a[x][y] = (t.a[x][y] << (op3 % 4)) % 16
	case OP_SHRL:
		t.a[x][y] = t.a[x][y] >> (op3 % 4)
	case OP_SHRA:
		if t.a[x][y]&0x8 != 0 {
			for i := 0; i < int(op3%4); i += 1 {
				t.a[x][y] = t.a[x][y]/2 + 8
			}
		} else {
			t.a[x][y] = t.a[x][y] >> (op3 % 4)
		}
	case OP_META:
		metaOP := x
		meta1 := t.a[(t.IP.x+(t.IP.y+1)/16)%16][(t.IP.y+1)%16]
		meta2 := t.a[(t.IP.x+(t.IP.y+2)/16)%16][(t.IP.y+2)%16]
		meta3 := t.a[(t.IP.x+(t.IP.y+3)/16)%16][(t.IP.y+3)%16]
		meta(t, metaOP, meta1, meta2, meta3)
	}
}

func meta(t *Tile, op, arg1, arg2, arg3 uint8) {
}
