package tdk

// route map
//
//  +		: current position
//  numbers	: next position according to v (in hex)
//
/*
  0     6     D
     3  7  A
  1  4  +  B  E
     5  8  C
  2     9     F
*/
var routeMapX = [16]uint8{14, 0, 2, 15, 0, 1, 14, 15, 1, 2, 15, 0, 1, 14, 0, 2}
var routeMapY = [16]uint8{14, 14, 14, 15, 15, 15, 0, 0, 0, 0, 1, 1, 1, 2, 2, 2}

func (ip *IP) OPLookup() TDK_OP {
	return ip.Tile.OPLookup()
}

//find the operation to be executed in current cycle
func (t *Tile) OPLookup() TDK_OP {
	var visited [16]uint8
	x := t.IP.x
	y := t.IP.y

	for visited[t.a[x][y]%16] == 0 {
		visited[t.a[x][y]%16] = 1

		// move the IP
		// if the next position is out of tile border,
		// the real x or y coordinate should modulo to 16
		v := (t.a[x][y] + x + y) % 16

		t.a[x][y] = (t.a[x][y] + 1) % 16
		for !t.Fitted(x, y, t.a[x][y]) {
			t.a[x][y] = (t.a[x][y] + 1) % 16
		}
		t.a[x][y] += 16

		x = (routeMapX[v] + x) % 16
		y = (routeMapY[v] + y) % 16
	}

	t.IP.x = x
	t.IP.y = y
	return TDK_OP(t.a[x][y])
}
