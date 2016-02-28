// An interactive terminal implementation of Conway's Game of Life.
package main

import (
	"github.com/jfgreen/gameoflife-termbox/game"
	"github.com/nsf/termbox-go"
	"math/rand"
	"time"
)

type eventChanel chan termbox.Event

const alive rune = '█'
const aliveCol = termbox.ColorGreen
const bgCol = termbox.ColorDefault

func main() {
	rand.Seed(time.Now().UnixNano())
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	w, h := termbox.Size()
	l := game.NewLife(w, h)

	eventQueue := make(eventChanel)
	publishEvents(eventQueue)

	draw(l)
	running := true
	for running {
		select {
		case event := <-eventQueue:
			if event.Type == termbox.EventKey && event.Key == termbox.KeyEsc {
				running = false
			}
		default:
			l.Step()
			draw(l)
			// TODO: smooth updating technique, measuring time of each step
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func publishEvents(c eventChanel) {
	go func() {
		for {
			c <- termbox.PollEvent()
		}
	}()
}

func draw(l *game.Life) {

	termbox.Clear(bgCol, bgCol)
	w, h := l.Size()
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if l.State(x, y) {
				termbox.SetCell(x, y, alive, aliveCol, bgCol)
			}
		}
	}

	termbox.Flush()

}
