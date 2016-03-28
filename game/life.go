package game

type Life struct {
	world, temp   *Grid
	Width, Height int
}

func NewEmptyLife(width, height int) *Life {
	world := NewGrid(width, height)
	temp := NewGrid(width, height)
	return &Life{world, temp, width, height}
}

func NewLifeFromGrid(g *Grid) *Life {
	temp := NewGrid(g.Width, g.Height)
	return &Life{g, temp, g.Width, g.Height}
}

func (l *Life) Step() {
	for x := 0; x < l.world.Width; x++ {
		for y := 0; y < l.world.Height; y++ {
			l.temp.Set(x, y, l.nextState(x, y))
		}
	}
	l.world, l.temp = l.temp, l.world
}

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

func (l *Life) Alive(x, y int) bool {
	return l.world.Get(x, y)
}

func (l *Life) Flip(x, y int) {
	l.world.Set(x, y, !l.world.Get(x, y))
}

func (l *Life) Randomise() {
	l.world.Randomise()
}

func (l *Life) Resize(width, height int) {
	l.world = ResizeGrid(l.world, width, height)
	l.temp = NewGrid(width, height)
	l.Width = width
	l.Height = height
}

func (l *Life) Translate() {

}

func (l *Life) ContainsCoordinate(x, y int) bool {
	return x >= 0 && x < l.Width && y >= 0 && y < l.Height
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
