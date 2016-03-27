package game

import (
	"os"
	"bufio"
	"fmt"
	"errors"
	"strings"
	"strconv"
)

const (
	life16Header = "#Life 1.06"
	life16CoordinateSeparator = " "
)

type point struct {
	x, y int
}

//TODO: Refactor this, especially error handling

func LoadLife16File(filepath string) (*Grid, error) {
	statefile, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer statefile.Close()
	scanner := bufio.NewScanner(statefile)
	return parseLife16(scanner)
}

func parseLife16(scanner *bufio.Scanner) (*Grid, error) {
	err := consumeLine(scanner, life16Header)
	if err != nil {
		return nil, err
	}
	coordinates, err:= consumeCoordinates(scanner)
	if err != nil {
		return nil, err
	}
	return createGrid(coordinates), nil
}

func consumeCoordinates(scanner *bufio.Scanner) ([]point, error) {
	coordinates := make([]point, 0)
	for scanner.Scan() {
		p, err := parseCoordinate(scanner.Text())
		if err != nil {
			return nil, err
		}
		coordinates = append(coordinates, p)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return coordinates, nil
}

func parseCoordinate(s string) (p point, err error) {
	components := strings.Split(s, life16CoordinateSeparator)
	if len(components) != 2 {
		return point{0, 0}, errors.New(fmt.Sprintf("expected 2 components, got %v", components))
	}
	x, err := strconv.Atoi(components[0])
	y, err := strconv.Atoi(components[1])
	p = point{x, y}
	return
}

func consumeLine(scanner *bufio.Scanner, expected string) error {
	if scanner.Scan() {
		line := scanner.Text()
		if line != expected {
			return errors.New(fmt.Sprintf("expected %s, but got %s.", expected, line))
		}
		return nil
	} else  {
		if err := scanner.Err(); err != nil {
			return err
		}
		return errors.New("scanner unexpectadly ran out of input")
	}
}

func createGrid(coordinates []point) (*Grid) {
	if len(coordinates) == 0 {
		return NewGrid(0,0)
	}
	// Smallest point in coordinates will become the new origin in the grid.
	minPoint, maxPoint := getBoundingPoints(coordinates)
	width := maxPoint.x - minPoint.x + 1
	height := maxPoint.y - minPoint.y + 1
	grid := NewGrid(width, height)
	for _, c := range coordinates {
		grid.Set(c.x - minPoint.x, c.y-minPoint.y, true)
	}
	return grid
}

func getBoundingPoints(coordinates []point) (point, point) {
	xmin := coordinates[0].x
	xmax := coordinates[0].x
	ymin := coordinates[0].y
	ymax := coordinates[0].y
	for _, c := range coordinates[1:] {
		if c.x < xmin {
			xmin = c.x
		}
		if c.y < ymin {
			ymin = c.y
		}
		if c.x > xmax {
			xmax = c.x
		}
		if c.y > ymax {
			ymax = c.y
		}
	}
	return point{xmin, ymin}, point{xmax, ymax}
}


