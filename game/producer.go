package game

import "math/rand"

type LifeProducer interface {
	produce(int, int) (*Life, error)
}

type RandomLifeProducer struct {
	Seed int64
}

func (c *RandomLifeProducer) produce(w, h int) (*Life, error) {
	l := NewLifeEmpty(w, h)
	rand.Seed(c.Seed)
	l.Randomise()
	return l, nil
}

type SaveFileLifeProducer struct {
	Filepath string
}

func (s *SaveFileLifeProducer) produce(w, h int) (*Life, error) {
	grid, err := LoadLife16File(s.Filepath)
	if err != nil {
		return nil, err
	}
	tx := (w / 2) + grid.Width/2
	ty := (h / 2) + grid.Height/2
	grid = TranslateGrid(ResizeGrid(grid, w, h), tx, ty)
	return NewLifeFromGrid(grid), nil
}