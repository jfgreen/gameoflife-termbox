package game

import (
	"math/rand"
)

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

func ResizeGrid(g *Grid, width, height int) *Grid {
	newGrid := NewGrid(width, height)
	minWidth := min(g.Width, width)
	minHeight := min(g.Height, height)
	for x := 0; x < minWidth; x++ {
		for y := 0; y < minHeight; y++ {
			newGrid.Set(x, y, g.Get(x, y))
		}
	}
	return newGrid
}

func TranslateGrid(g *Grid, dx, dy int) *Grid {
	newGrid := NewGrid(g.Width, g.Height)
	for x := 0; x < g.Width; x++ {
		for y := 0; y < g.Height; y++ {
			newGrid.Set(x, y, g.GetToroidal(x+dx, y+dy))
		}
	}
	return newGrid
}

func (g *Grid) Randomise() {
	for x := 0; x < g.Width; x++ {
		for y := 0; y < g.Height; y++ {
			g.Set(x, y, rand.Intn(2) == 0)
		}
	}
}

func (g *Grid) ActiveNeighbours(x, y int) int {
	activeNeighbours := 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if (dx != 0 || dy != 0) && g.GetToroidal(x+dx, y+dy) {
				activeNeighbours++
			}
		}
	}
	return activeNeighbours
}

func (g *Grid) Get(x, y int) bool {
	return g.cells[y][x]
}

func (g *Grid) Set(x, y int, s bool) {
	g.cells[y][x] = s
}

func (g *Grid) GetToroidal(x, y int) bool {
	return g.Get(g.toroidal(x, y))
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



