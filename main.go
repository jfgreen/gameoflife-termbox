// An interactive terminal implementation of Conway's Game of Life.
package main

import (
	"flag"
	"fmt"
	"github.com/jfgreen/gameoflife-termbox/game"
	"log"
	"os"
)

func main() {
	fps := flag.Int("fps", 25, "Frames per second.")
	logpath := flag.String("log", os.DevNull, "Path of logfile to write debugging messages to.")
	flag.Parse()

	logfile := createLogfile(*logpath)
	defer logfile.Close()
	log.SetOutput(logfile)

	err := game.Begin(*fps)
	if err != nil {
		fmt.Println("Problem running game: ", err)
		os.Exit(1)
	}
}

func createLogfile(filepath string) *os.File {
	logfile, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Problem opening file for logging: ", err)
		os.Exit(1)
	}
	return logfile
}
