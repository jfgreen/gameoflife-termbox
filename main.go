// An interactive terminal implementation of Conway's Game of Life.
package main

import (
	"github.com/jfgreen/gameoflife-termbox/game"
	"flag"
)

func main() {

	fps := flag.Int("fps", 25, "Frames per second.")
	flag.Parse()

	game.Begin(*fps);
}