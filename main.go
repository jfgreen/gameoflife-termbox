// An interactive terminal implementation of Conway's Game of Life.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Grid struct {
	cells         [][]bool
	width, height int
}

func main() {
	g := NewGrid(4, 6)
	rand.Seed(time.Now().UnixNano())
	g.Randomise()
	fmt.Println(g)

}

func NewGrid(w, h int) *Grid {
	cells := make([][]bool, w)
	c := make([]bool, w*h)
	for i := range cells {
		cells[i], c = c[:h], c[h:]
	}
	return &Grid{cells, w, h}
}

func (g *Grid) State(x, y int) bool {
	return g.cells[x][y]
}

func (g *Grid) Randomise() {
	for x := 0; x < g.width; x++ {
		for y := 0; y < g.height; y++ {
			g.cells[x][y] = rand.Intn(2) == 0
		}
	}
}

func (g *Grid) aliveNeighbours(x, y int) int {
	aliveNeighbours := 0
	for dx := range []int{-1, 0, 1} {
		for dy := range []int{-1, 0, 1} {
			nx, ny := g.toroidal(x+dx, y+dy)
			if (nx != x || ny != y) && g.cells[nx][ny] {
				aliveNeighbours++
			}
		}
	}
	return aliveNeighbours
}

func (g *Grid) toroidal(x, y int) (tx, ty int) {
	tx = x % g.width
	ty = y % g.height
	if tx < 0 {
		tx += g.width
	}
	if ty < 0 {
		tx += g.height
	}
	return
}
