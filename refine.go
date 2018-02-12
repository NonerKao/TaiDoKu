package tdk

func (ip *IP) Refine() {
	ip.Tile.Refine()
}

// Refine refine the tile according to the value of t.a[x][y]
func (t *Tile) Refine() {
	norm := []uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	gray := []uint8{0, 1, 3, 2, 6, 7, 5, 4, 12, 13, 15, 14, 10, 11, 9, 8}
	seqs := [2][]uint8{norm, gray}

	x := t.IP.x
	y := t.IP.y
	v := t.a[x][y] + x + y

	if v&0x8 != 0 {
		for i := 0; i < 16; i += 1 {
			for j := 0; j < 8; j += 1 {
				temp := t.a[i][j]
				t.a[i][j] = t.a[i][15-j]
				t.a[i][15-j] = temp
			}
		}
	}

	if v&0x4 != 0 {
		for i := 0; i < 8; i += 1 {
			for j := 0; j < 16; j += 1 {
				temp := t.a[i][j]
				t.a[i][j] = t.a[15-i][j]
				t.a[15-i][j] = temp
			}
		}
	}

	if v&0x2 != 0 {
		for i := 1; i < 16; i += 1 {
			for j := 0; j < i; j += 1 {
				temp := t.a[i][j]
				t.a[i][j] = t.a[j][i]
				t.a[j][i] = temp
			}
		}
	}

	seq := seqs[v%2]
	var status [16][16]uint8

	for i := 0; i < 16; i += 1 {
		for j := 0; j < 16; j += 1 {
			k := getIndex(seq, t.a[i][j])
		backtrace:
			if t.notFixed(uint8(i), uint8(j)) {
				for ; status[i][j] < 16; status[i][j] += 1 {
					if !t.Fitted(uint8(i), uint8(j), uint8(seq[k])) {
						k = (k + 1) % 16
						continue
					}
					t.a[i][j] = seq[k] + 16
					break
				}

				// no cadidate, need backtrace
				if status[i][j] == 16 {
					status[i][j] = 0
					i, j = t.prevCell(i, j)

					k = getIndex(seq, t.a[i][j]-16)
					k = (k + 1) % 16
					t.a[i][j] = seq[k]
					goto backtrace
				}
			}
		}
	}

	for i := 0; i < 16; i += 1 {
		for j := 0; j < 16; j += 1 {
			t.a[i][j] = t.a[i][j] - 16
		}
	}
}

func (t *Tile) prevCell(x, y int) (int, int) {
	if y == 0 {
		return x - 1, 15
	} else {
		return x, y - 1
	}
}

func (t *Tile) Fitted(x, y uint8, v uint8) bool {
	for i := 0; i < 16; i += 1 {
		if t.a[i][y] == v+16 {
			return false
		}
	}

	for j := 0; j < 16; j += 1 {
		if t.a[x][j] == v+16 {
			return false
		}
	}

	for i := 4 * (x / 4); i < 4*(x/4+1); i += 1 {
		for j := 4 * (y / 4); j < 4*(y/4+1); j += 1 {
			if t.a[i][j] == v+16 {
				return false
			}
		}
	}

	return true
}

func (t *Tile) notFixed(x, y uint8) bool {
	return t.a[x][y] < 16
}

func getIndex(seq []uint8, v uint8) uint8 {
	for i, n := range seq {
		if n == v {
			return uint8(i)
		}
	}
	return 255
}
