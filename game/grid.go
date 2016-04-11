package game

import (
	"math/rand"
)

// Grid represent a two dimensional grid of boolean values
type Grid struct {
	cells         [][]bool
	Width, Height int
}

// NewGrid returns a Grid of size width, height, with all values set to false
func NewGrid(width, height int) *Grid {
	cells := make([][]bool, height)
	c := make([]bool, width*height)
	for i := range cells {
		cells[i] = c[:width]
		c = c[width:]
	}
	return &Grid{cells, width, height}
}

// ResizeGrid returns a copy of grid g, resized to dimensions specified by width and height.
// If g is smaller than the given dimensions, cells created to make up the size will be false.
// If g is larger than the given dimensions, cells that don't fit in the returned grid will be truncated.
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

// TranslateGrid returns a copy of grid g, translated by values dx and dy.
// Edges are wrapped around such that a cell translated over one edge will appear on the opposite.
func TranslateGrid(g *Grid, dx, dy int) *Grid {
	newGrid := NewGrid(g.Width, g.Height)
	for x := 0; x < g.Width; x++ {
		for y := 0; y < g.Height; y++ {
			newGrid.Set(x, y, g.GetToroidal(x+dx, y+dy))
		}
	}
	return newGrid
}

// Randomise will set every cell in the grid to a random boolean value.
func (g *Grid) Randomise() {
	for x := 0; x < g.Width; x++ {
		for y := 0; y < g.Height; y++ {
			g.Set(x, y, rand.Intn(2) == 0)
		}
	}
}

// ActiveNeighbours returns the number of cells neighbouring cell at x, y, that are currently set to true.
// Edges are wrapped around such that a cell on an edge of the grid neighbours cells on the opposite edge.
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

// Get the value of cell at x, y.
func (g *Grid) Get(x, y int) bool {
	return g.cells[y][x]
}

// Set the value of cell at x, y.
func (g *Grid) Set(x, y int, s bool) {
	g.cells[y][x] = s
}

// Get the value of cell at x, y, toroidaly.
func (g *Grid) GetToroidal(x, y int) bool {
	return g.Get(g.toroidal(x, y))
}

// toroidal returns grid coordinates based on the toroidal value of coordinate x, y.
// If x or y are outside the range of the grid, the returned coordinate will be wrapped around.
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

// min returns the smaller of integers a and b.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

