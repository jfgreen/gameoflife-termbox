package game

// Life represents a bounded 2D simulation of the game of life.
type Life struct {
	world, prev   *Grid
	Width, Height int
}

// NewLifeEmpty constructs and returns an empty Life, where all cells are dead.
func NewLifeEmpty(width, height int) *Life {
	world := NewGrid(width, height)
	temp := NewGrid(width, height)
	return &Life{world, temp, width, height}
}
// NewLifeFromGrid constructs and returns a Life from en existing grid.
func NewLifeFromGrid(g *Grid) *Life {
	temp := NewGrid(g.Width, g.Height)
	return &Life{g, temp, g.Width, g.Height}
}

// Step increments the simulation by a single step.
func (l *Life) Step() {
	for x := 0; x < l.world.Width; x++ {
		for y := 0; y < l.world.Height; y++ {
			l.prev.Set(x, y, l.nextState(x, y))
		}
	}
	l.world, l.prev = l.prev, l.world
}

// nextState returns the next state of the cell at the coordinate (x, y).
func (l *Life) nextState(x, y int) bool {
	switch l.world.ActiveNeighbours(x, y) {
	case 3:
		return true
	case 2:
		return l.world.Get(x, y)
	default:
		return false
	}
}

// Alive returns current state of the cell at the coordinate (x, y).
func (l *Life) Alive(x, y int) bool {
	return l.world.Get(x, y)
}

// Flip inverts the state of the cell at the coordinate (x, y).
func (l *Life) Flip(x, y int) {
	l.world.Set(x, y, !l.world.Get(x, y))
}

func (l *Life) Randomise() {
	l.world.Randomise()
}

// Resize destructively changes the size the Life.
// If the new size is smaller than the current size, the Life will be truncated.
func (l *Life) Resize(width, height int) {
	l.world = ResizeGrid(l.world, width, height)
	l.prev = NewGrid(width, height)
	l.Width = width
	l.Height = height
}

// ContainsCoordinate returns a boolean indicating if the given coordinate (x, y) is in the bound of the Life
func (l *Life) ContainsCoordinate(x, y int) bool {
	return x >= 0 && x < l.Width && y >= 0 && y < l.Height
}


