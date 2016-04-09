package game

import (
	"github.com/nsf/termbox-go"
	"log"
	"math/rand"
	"time"
)

const alive rune = '●'
const aliveCol = termbox.ColorBlue
const birthCol = termbox.ColorCyan
const bgCol = termbox.ColorDefault

type Game struct {
	FrameDelay      time.Duration
	Producer        LifeProducer
	life            *Life
	eventQueue      chan termbox.Event
	lastTick        time.Time
	exiting, paused bool
	err             error
}

// TODO: Allow flag for exiting after a certain number of iterations
// TODO: Document expected format and supply some choice examples
// TODO: Flag for specifying alive
// TODO: Do we want to do some tests?
// TODO: Godoc?
// TODO: Rainbow mode.

type LifeProducer interface {
	produce(int, int) (*Life, error)
}

type RandomLifeProducer struct {
	Seed int64
}

func (c *RandomLifeProducer) produce(w, h int) (*Life, error) {
	l := NewEmptyLife(w, h)
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

func (g *Game) Init() {
	g.paused = false
	g.exiting = false
	g.err = termbox.Init()
	if g.err != nil {
		return
	}
	g.eventQueue = make(chan termbox.Event)
	publishEvents(g.eventQueue)
	termbox.SetInputMode(termbox.InputAlt | termbox.InputMouse)
	g.life, g.err = g.Producer.produce(termbox.Size())
}

func (g *Game) Run() {
	for g.running() {
		g.loop()
	}
	termbox.Close()
}

func (g *Game) Err() error {
	return g.err
}

func (g *Game) running() bool {
	return !g.exiting && g.err == nil
}

func (g *Game) loop() {
	select {
	case event := <-g.eventQueue:
		g.handleEvent(event)
	default:
		if !g.paused {
			g.life.Step()
		}
		g.life.Draw()
		g.wait()
	}
}

func (g *Game) handleEvent(e termbox.Event) {
	switch e.Type {
	case termbox.EventKey:
		g.handleKeyEvent(e)
	case termbox.EventMouse:
		g.handleMouseEvent(e)
	case termbox.EventResize:
		g.handleResize(e)
	}
}

func (g *Game) handleKeyEvent(e termbox.Event) {
	switch {
	case exitEvent(e):
		log.Println("Request to exit recieved, requesting termination of game loop")
		g.exiting = true
	case e.Ch == 'r':
		log.Println("Recreating game")
		g.life, g.err = g.Producer.produce(termbox.Size())
	case e.Key == termbox.KeySpace:
		log.Println("Toggling paused")
		g.paused = !g.paused
	}
}

func (g *Game) handleMouseEvent(e termbox.Event) {
	switch {
	case e.Key == termbox.MouseLeft:
		if g.life.ContainsCoordinate(e.MouseX, e.MouseY) {
			log.Printf("Mouse clicked, inverting state at (%d, %d)\n", e.MouseX, e.MouseY)
			g.life.Flip(e.MouseX, e.MouseY)
		} else {
			log.Printf("Ignoring mouse click (%d, %d) outside game area, (0,0) to (%d, %d)\n",
				e.MouseX, e.MouseY, g.life.Width, g.life.Height)
		}
	}
}

func (g *Game) handleResize(e termbox.Event) {
	log.Printf("Terminal resizing from (%d,%d) to (%d,%d)\n", g.life.Width, g.life.Height, e.Width, e.Height)
	g.life.Resize(e.Width, e.Height)
	g.life.Draw()
}

func exitEvent(e termbox.Event) bool {
	return e.Ch == 'q' ||
		e.Key == termbox.KeyEsc ||
		e.Key == termbox.KeyCtrlC ||
		e.Key == termbox.KeyCtrlD
}

func (g *Game) wait() {
	timeSinceLastTick := time.Since(g.lastTick)
	g.lastTick = time.Now()
	timeSinceLastTick.Hours()
	time.Sleep(g.FrameDelay - timeSinceLastTick)
}

func publishEvents(c chan termbox.Event) {
	go func() {
		for {
			c <- termbox.PollEvent()
		}
	}()
}

func (l *Life) Draw() {

	termbox.Clear(bgCol, bgCol)
	w, h := l.Width, l.Height
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if l.world.Get(x, y) {
				if l.prev.Get(x, y) {
					termbox.SetCell(x, y, alive, aliveCol, bgCol)
				} else {
					termbox.SetCell(x, y, alive, birthCol, bgCol)
				}
			}
		}
	}

	termbox.Flush()
}
