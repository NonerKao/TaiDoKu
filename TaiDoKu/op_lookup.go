package tdk

// route map
//
//  +		: current position
//  numbers	: next position according to v (in hex)
//
/*
  0     1     2
     3  4  5
  6  7  +  8  9
     A  B  C
  D     E     F
*/
var routeMapX = [16]uint8{14, 0, 2, 15, 0, 1, 14, 15, 1, 2, 15, 0, 1, 14, 0, 2}
var routeMapY = [16]uint8{14, 14, 14, 15, 15, 15, 0, 0, 0, 0, 1, 1, 1, 2, 2, 2}

//find the operation to be executed in current cycle
func (t *Tile) OPLookup(x, y uint8) uint8 {
	var visited [16]uint8

	for visited[t.A[x][y]] == 0 {
		visited[t.A[x][y]] = 1

		// move the IP
		// if the next position is out of tile border,
		// the real x or y coordinate should modulo to 16
		v := t.A[x][y] + x + y
		t.A[x][y] += 1
		x = (routeMapX[v] + 16) % 16
		y = (routeMapY[v] + 16) % 16
	}

	return t.A[x][y]
}
