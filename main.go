// An interactive terminal implementation of Conway's Game of Life.
package main

import (
	"fmt"
	"math/rand"
	"time"
	"github.com/jfgreen/gameoflife-termbox/life"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	l := life.NewLife(40, 20)
	for i := 0; i < 300; i++ {
		l.Step()
		fmt.Print("\x0c", l) // Clear screen and print field.
		time.Sleep(time.Second / 10)
	}
}

