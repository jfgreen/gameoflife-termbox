package life

import (
	"math/rand"
	"bytes"
)

const alive rune = '█'
const dead rune = '░'

type Grid struct {
	cells         [][]bool
	Width, Height int
}

func NewGrid(width, height int) *Grid {
	cells := make([][]bool, height)
	c := make([]bool, width*height)
	for i := range cells {
		cells[i] = c[:width]
		c = c[width:]
	}
	return &Grid{cells, width, height}
}

func (g *Grid) Randomise() {
	for x := 0; x < g.Width; x++ {
		for y := 0; y < g.Height; y++ {
			g.Set(x, y, rand.Intn(2) == 0)
		}
	}
}

func (g *Grid) AliveNeighbours(x, y int) int {
	aliveNeighbours := 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			nx, ny := g.toroidal(x+dx, y+dy)
			if (dx != 0 || dy != 0) && g.Get(nx, ny) {
				aliveNeighbours++
			}
		}
	}
	return aliveNeighbours
}

func (g* Grid) Get(x, y int) bool {
	return g.cells[y][x]
}

func (g* Grid) Set(x, y int, s bool) {
	g.cells[y][x] = s
}

func (g *Grid) toroidal(x, y int) (tx, ty int) {
	tx = x % g.Width
	ty = y % g.Height
	if tx < 0 {
		tx += g.Width
	}
	if ty < 0 {
		ty += g.Height
	}
	return
}

func (g *Grid) String() string {
	var state bytes.Buffer
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			if g.Get(x, y) {
				state.WriteRune(alive)
			} else {
				state.WriteRune(dead)
			}
		}
		state.WriteRune('\n')
	}
	return state.String()
}
