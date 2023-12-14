package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	startingX, startingY int
	stk                  []*tile
	g                    *grid
)

func main() {
	parseInput()

	g.ConnectTiles()

	t := g.Tile(startingX, startingY)

	stk = append(stk, t)
	visited := make(map[int]map[int]bool)

	var (
		cur     *tile
		maxDist int
		c int
	)
	for len(stk) > 0 {
		fmt.Print(c,",")
		c++
		cur, stk = stk[0], stk[1:]
		if cur.dist > maxDist {
			maxDist = cur.dist
		}

		if hasVisited(cur, visited) {
			continue
		}

		if cur.conns == 1 {
			fmt.Println(cur.x, cur.y)
		}

		check(cur.nconn, cur, visited)
		check(cur.econn, cur, visited)
		check(cur.wconn, cur, visited)
		check(cur.sconn, cur, visited)

		if _, ok := visited[cur.x]; !ok {
			visited[cur.x] = make(map[int]bool)
		}
		visited[cur.x][cur.y] = true
	}

	fmt.Println("Max dist:", maxDist)
}

func check(conn *tile, cur *tile, visited map[int]map[int]bool) {
	if conn == nil {
		return
	}
	// If we have already set the distance, then check if the distance
	// from current node is shorter and update
	if conn.dist != 0 {
		if conn.dist > cur.dist+1 {
			conn.dist = cur.dist + 1
		}
	} else {
		conn.dist = cur.dist + 1
	}

	if !hasVisited(conn, visited) {
		stk = append(stk, conn)
	}
}

func hasVisited(t *tile, v map[int]map[int]bool) bool {
	var ok bool
	if _, ok = v[t.x]; !ok {
		return false
	}
	_, ok = v[t.x][t.y]
	return ok
}

func parseInput() {
	b, err := os.ReadFile("inputs.txt")
	if err != nil {
		panic(err)
	}

	split := strings.Split(string(b), "\n")

	g = newGrid(len(split)-1, len(split[0]))

	for i, l := range split {
		if len(l) == 0 {
			continue
		}

		for j := range l {
			r := rune(l[j])

			if r == 'S' {
				startingX, startingY = i, j
			}

			g.AddTile(r, i, j)
		}
	}
}

type grid struct {
	x, y  int
	tiles [][]*tile
}

func (g *grid) Tile(i, j int) *tile {
	return g.tiles[i][j]
}

func (g *grid) ConnectTiles() {
	for i := range g.tiles {
		for j := range g.tiles[i] {
			t := g.tiles[i][j]

			switch t.val {
			case '.':
				// . is ground; there is no pipe in this tile.
				continue
			case '|':
				// | is a vertical pipe connecting north and south.
				if i-1 >= 0 {
					t.AddNconn(g.tiles[i-1][j])
				}
				if i+1 < g.x {
					t.AddSconn(g.tiles[i+1][j])
				}
			case '-':
				// - is a horizontal pipe connecting east and west.
				if j+1 < g.y {
					t.AddEconn(g.tiles[i][j+1])
				}
				if j-1 >= 0 {
					t.AddWconn(g.tiles[i][j-1])
				}
			case 'L':
				// L is a 90-degree bend connecting north and east.
				if i-1 >= 0 {
					t.AddNconn(g.tiles[i-1][j])
				}
				if j+1 < g.y {
					t.AddEconn(g.tiles[i][j+1])
				}
			case 'J':
				// J is a 90-degree bend connecting north and west.
				if i-1 >= 0 {
					t.AddNconn(g.tiles[i-1][j])
				}
				if j-1 >= 0 {
					t.AddWconn(g.tiles[i][j-1])
				}
			case '7':
				// 7 is a 90-degree bend connecting south and west.
				if i+1 < g.x {
					t.AddSconn(g.tiles[i+1][j])
				}
				if j-1 >= 0 {
					t.AddWconn(g.tiles[i][j-1])
				}
			case 'F':
				// F is a 90-degree bend connecting south and east.
				if i+1 < g.x {
					t.AddSconn(g.tiles[i+1][j])
				}
				if j+1 < g.y {
					t.AddEconn(g.tiles[i][j+1])
				}
			}
		}
	}
}

func (g *grid) AddTile(r rune, i, j int) {
	t := &tile{val: r, x: i, y: j}
	g.tiles[i][j] = t
}

func newGrid(i, j int) *grid {
	ts := make([][]*tile, i)
	for k := range ts {
		ts[k] = make([]*tile, j)
	}
	return &grid{x: i, y: j, tiles: ts}
}

type tile struct {
	x, y        int
	val         rune
	dist, conns int
	nconn       *tile
	econn       *tile
	sconn       *tile
	wconn       *tile
}

func (t *tile) AddNconn(other *tile) {
	if t.nconn == nil {
		t.conns++
		t.nconn = other
	}
	if other.sconn == nil {
		other.conns++
		other.sconn = t
	} else if other.sconn != t {
		log.Fatalf("Fuck man, %v,%v has connection to %v,%v", other.x, other.y, other.sconn.x, other.sconn.y)
	}
}

func (t *tile) AddSconn(other *tile) {
	if t.sconn == nil {
		t.conns++
		t.sconn = other
	}
	if other.nconn == nil {
		other.conns++
		other.nconn = t
	} else if other.nconn != t {
		log.Fatalf(
			"Fuck man, %v,%v has connection to %v,%v",
			other.x, other.y, other.nconn.x, other.nconn.y,
		)
	}
}

func (t *tile) AddEconn(other *tile) {
	if t.econn == nil {
		t.conns++
		t.econn = other
	}
	if other.wconn == nil {
		other.conns++
		other.wconn = t
	} else if other.wconn != t {
		log.Fatalf(
			"Fuck man, %v,%v has connection to %v,%v",
			other.x, other.y, other.wconn.x, other.wconn.y,
		)
	}
}

func (t *tile) AddWconn(other *tile) {
	if t.wconn == nil {
		t.conns++
		t.wconn = other
	}

	if other.econn == nil {
		other.conns++
		other.econn = t
	} else if other.econn != t {
		log.Fatalf(
			"Fuck man, %v,%v has connection to %v,%v",
			other.x, other.y, other.econn.x, other.econn.y,
		)
	}
}
