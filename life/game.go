package life

type Life struct {
	world, temp *Grid
}


func NewLife(width, height int) *Life {
	world := NewGrid(width, height)
	temp := NewGrid(width, height)
	world.Randomise()
	return &Life{world, temp}
}

func (l *Life) nextState(x, y int) bool {
	switch l.world.AliveNeighbours(x, y) {
	case 3:
		return true
	case 2:
		return l.world.Get(x,y)
	default:
		return false
	}
}

func (l *Life) Step() {
	for x := 0; x < l.world.Width; x++ {
		for y := 0; y < l.world.Height; y++ {
			l.temp.Set(x, y, l.nextState(x, y))
		}
	}
	l.world, l.temp = l.temp, l.world
}

func (l *Life) String() string {
	return l.world.String()
}
