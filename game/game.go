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

// TODO: Click to toggle
// TODO: Add config, debug, interactivity - take a look at gomatrix for some ideas
// TODO: Command line argument for FPS
// TODO: Add some proper control channels to exit more gracefully
// TODO: Handle resizing

func Begin(fps int) {

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

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
	if e.Type == termbox.EventKey {
		switch {
		case exitEvent(e):
			g.running = false
		case e.Ch == 'r':
			g.life.Randomise()
		case e.Key == termbox.KeySpace:
			g.paused = !g.paused
		}
	}
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
