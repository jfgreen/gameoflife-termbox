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
	fps := flag.Int("fps", 10, "Frames per second.")
	logpath := flag.String("log", os.DevNull, "Path of logfile to write debugging messages to.")
	seed := flag.Int64("seed", 0, "Seed to be used in initialisation of random life.")
	savefile := flag.String("file", "",
		"Path of pattern file to initialise game with. Takes precedence over --seed.")
	alive := flag.String("alive", "●", "Character to use to render alive cells.")
	flag.Parse()

	logfile := createLogfile(*logpath)
	defer logfile.Close()
	log.SetOutput(logfile)
	producer := createLifeProducer(*seed, *savefile)
	delay := fpsToDelay(*fps)
	aliveCell := aliveRune(*alive)

	log.Println("Starting game of life.")
	game := &game.Game{FrameDelay: delay, Producer: producer, Alive: aliveCell}
	game.Init()
	game.Run()
	log.Println("Exiting game of life.")

	if game.Err() != nil {
		fmt.Println("Problem running game: ", game.Err())
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

func fpsToDelay(fps int) time.Duration {
	if fps < 1 || fps > 60 {
		fmt.Println("Error: fps not within range 1-60")
		os.Exit(1)
	}
	return time.Duration((float32(time.Second) / float32(fps)))
}

func aliveRune(s string) rune {
	runes := []rune(s)
	if len(runes) != 1 {
		fmt.Println("Error: alive is not a single rune")
		os.Exit(1)
	}
	return runes[0]
}
