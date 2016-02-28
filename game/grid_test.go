package game

import "testing"

func TestNewGridNormalCases(t *testing.T) {
	cases := []struct{ x, y int }{
		{1, 1},
		{5, 5},
		{5, 10},
		{10, 5},
		{1000, 1000},
	}
	for _, c := range cases {
		got := NewGrid(c.x, c.y).cells
		correctSize := verifyGrid(got, c.x, c.y)
		if !correctSize {
			t.Errorf("NewGrid(%d, %d) created wrong size grid %v", c.y, c.y, got)
		}
	}
}

//TODO: Test abnormal grid size, maybe make grid return an error

func verifyGrid(grid [][]bool, x, y int) bool {
	if len(grid) != y {
		return false
	}
	for _, l := range grid {
		if len(l) != x {
			return false
		}
	}
	return true
}

func TestNeighbours(t *testing.T) {
	cases := []struct {
		grid *Grid
		x, y int
		want int
	}{
		{grid: &Grid{[][]bool{
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
		}, 5, 5}, x: 2, y: 2, want: 0},
		{grid: &Grid{[][]bool{
			{true, true, true, true, true},
			{true, true, true, true, true},
			{true, true, true, true, true},
			{true, true, true, true, true},
			{true, true, true, true, true},
		}, 5, 5}, x: 2, y: 2, want: 8},
		{grid: &Grid{[][]bool{
			{false, false, false, true, false},
			{false, true, true, false, false},
			{false, true, false, false, true},
			{false, false, true, false, true},
			{false, true, false, false, false},
		}, 5, 5}, x: 3, y: 2, want: 4},
		{grid: &Grid{[][]bool{
			{true, false, false, true, true},
			{false, false, false, false, false},
			{true, false, false, true, true},
			{true, false, false, true, false},
		}, 5, 4}, x: 4, y: 3, want: 8},
		{grid: &Grid{[][]bool{
			{false, true, false, false, true},
			{true, true, false, false, true},
			{false, false, false, false, false},
			{true, true, false, false, true},
		}, 5, 4}, x: 0, y: 0, want: 8},
	}
	for _, c := range cases {
		got := c.grid.ActiveNeighbours(c.x, c.y)
		if got != c.want {
			t.Errorf("Given Grid %q, AliveNeighbours(%d,%d) == %d, want %d",
				c.grid.cells, c.x, c.y, got, c.want)
		}
	}
}
