package main

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

type SentryDir int

const (
	DIR_UP SentryDir = iota + 1
	DIR_RIGHT
	DIR_DOWN
	DIR_LEFT
)

type Sentry struct {
	pos       utils.Point
	direction SentryDir
}

func moveSentry(
	inSentry *Sentry,
	inObs []utils.Point,
	inMaxRows, inMaxCols int,
	inTotal *int,
	inVisited *map[string]bool,
) (bool, bool) {
	if inSentry.pos.X < 0 || inSentry.pos.Y < 0 || inSentry.pos.X >= inMaxCols || inSentry.pos.Y >= inMaxRows {
		return false, false // left the map
	}

	// check if sentry is over obstacle
	skipCount := false
	for i := 0; i < len(inObs); i++ {
		obs := inObs[i]
		if obs.X == inSentry.pos.X && obs.Y == inSentry.pos.Y { // over the obstacle
			skipCount = true
			// rotate
			switch inSentry.direction {
			case DIR_UP:
				inSentry.direction = DIR_RIGHT
				inSentry.pos.Y += 1
			case DIR_RIGHT:
				inSentry.direction = DIR_DOWN
				inSentry.pos.X -= 1
			case DIR_DOWN:
				inSentry.direction = DIR_LEFT
				inSentry.pos.Y -= 1
			case DIR_LEFT:
				inSentry.direction = DIR_UP
				inSentry.pos.X += 1
			}

			break
		}
	}

	// traverse
	if !skipCount {
		label := fmt.Sprintf("%v-%v", inSentry.pos.X, inSentry.pos.Y)
		if _, seen := (*inVisited)[label]; !seen {
			(*inTotal)++
			(*inVisited)[label] = true
		}

		switch inSentry.direction {
		case DIR_UP:
			inSentry.pos.Y -= 1
		case DIR_DOWN:
			inSentry.pos.Y += 1
		case DIR_LEFT:
			inSentry.pos.X -= 1
		case DIR_RIGHT:
			inSentry.pos.X += 1
		}
	}
	return true, !skipCount
}

func RunDay6Part1() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	fmt.Println("Day 6 Part 1 puzzle: Running")
	sum := 0

	// load map, find sentry location and direction
	cols := 0
	rows := 0
	var sentry Sentry
	var obstacles []utils.Point
	visted := make(map[string]bool)

	file, err := os.Open("./input/day6.txt")
	if err != nil {
		log.Fatal("Failed to fetch Day 6 input")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			cols = len(line)
			for i := 0; i < len(line); i++ {
				rune := line[i]
				if rune == '^' {
					sentry = Sentry{
						pos: utils.Point{
							X: i,
							Y: rows,
						},
						direction: DIR_UP,
					}
				} else if rune == '#' {
					obstacles = append(obstacles, utils.Point{
						X: i,
						Y: rows,
					})
				}
			}
		}
		rows++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("File read error.")
	}

	for {
		if goNext, _ := moveSentry(&sentry, obstacles, rows, cols, &sum, &visted); !goNext {
			break
		}
	}

	fmt.Println("Day 6 Part 1 puzzle: Result = ", sum)
}
