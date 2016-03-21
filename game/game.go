package game

import (
	"github.com/nsf/termbox-go"
	"time"
	"fmt"
)

const alive rune = '█'
const aliveCol = termbox.ColorGreen
const bgCol = termbox.ColorDefault

type Game struct {
	life       *Life
	eventQueue chan termbox.Event
	frameDelay time.Duration
	lastTick   time.Time
	running, paused    bool
}

// TODO: Add config, debug, interactivity - take a look at gomatrix or termloop for some ideas
// TODO: Add some proper control channels to exit more gracefully
// TODO: Can we do better than using a clever wait to do a loop. Look up go game loops. Anything channel based?
// TODO: Does a really small terminal break things? 1*1

func Begin(fps int) {

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputAlt | termbox.InputMouse)

	w, h := termbox.Size()
	life := NewLife(w, h)
	life.Randomise()

	delay := time.Duration((float32(time.Second)/float32(fps)))
	fmt.Println(delay)
	eventQueue := make(chan termbox.Event)
	publishEvents(eventQueue)
	game := &Game{life: life, eventQueue: eventQueue, frameDelay: delay}
	game.Start()

}

func (g *Game) Start() {
	g.running = true
	g.paused = false
	for g.running {
		g.loop()
	}
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
		g.running = false
	case e.Ch == 'r':
		g.life.Randomise()
	case e.Key == termbox.KeySpace:
		g.paused = !g.paused
	}
}

func (g *Game) handleMouseEvent(e termbox.Event) {
	switch {
	case e.Key == termbox.MouseLeft:
		g.life.Flip(e.MouseX, e.MouseY)
	}
}

func (g *Game) handleResize(e termbox.Event) {
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
	time.Sleep(g.frameDelay - timeSinceLastTick)
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
			if l.Alive(x, y) {
				termbox.SetCell(x, y, alive, aliveCol, bgCol)
			}
		}
	}

	termbox.Flush()
}
