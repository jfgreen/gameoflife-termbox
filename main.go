// An interactive terminal implementation of Conway's Game of Life.
package main

import "fmt"

type Grid struct {
	cells         [][]bool
	width, height int
}

func main() {
	//g := NewGrid(4, 6)
	//fmt.Println(g)
	i := -12
	w := 5
	i += w
	i %= w
	fmt.Println(i)
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

func (g *Grid) toroidal(x, y int) (tx, ty int) {
	//TODO: this
}

func (g *Grid) aliveNeighbours(x, y int) int {
	aliveNeighbours := 0
	for dx := range []int{-1, 0, 1} {
		for dy := range []int{-1, 0, 1} {
			nx, ny := g.toroidal(x+dx, y+dy)
			if (dx != 0 || dy != 0) && g.cells[nx][ny] {
				aliveNeighbours++
			}
		}
	}
	return aliveNeighbours
}
