// An interactive terminal implementation of Conway's Game of Life.
package main

import (
	"flag"
	"fmt"
	"github.com/jfgreen/gameoflife-termbox/game"
	"log"
	"os"
	"time"
)

func main() {
	fps := flag.Int("fps", 15, "Frames per second.")
	logpath := flag.String("log", os.DevNull, "Path of logfile to write debugging messages to.")
	seed := flag.Int64("seed", 0, "Seed to be used in initialisation of random life.")
	savefile := flag.String("file", "", "Path of pattern file to initalise game with. Takes precedence over --seed.")
	flag.Parse()

	logfile := createLogfile(*logpath)
	defer logfile.Close()
	log.SetOutput(logfile)
	producer := createLifeProducer(*seed, *savefile)
	err := game.Begin(*fps, producer)
	if err != nil {
		fmt.Println("Problem running game: ", err)
		os.Exit(1)
	}
}

func createLifeProducer(seed int64, filepath string) game.LifeProducer {
	if filepath != "" {
		return &game.SaveFileLifeProducer{filepath}
	}
	if seed != 0 {
		return &game.RandomLifeProducer{seed}
	}
	return &game.RandomLifeProducer{time.Now().UnixNano()}
}

func createLogfile(filepath string) *os.File {
	logfile, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Problem opening file for logging: ", err)
		os.Exit(1)
	}
	return logfile
}
