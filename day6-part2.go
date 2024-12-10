package main

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func testSentryLoop(
	inSentry *Sentry,
	inObs []utils.Point,
	inMaxRows, inMaxCols int,
	inVisited *map[string]bool,
) bool {
	var visited []utils.Point
	visited = append(visited, inSentry.pos)
	total := 0
	edgeMap := make(map[string]bool)
	for {
		goNext, appendPos := moveSentry(inSentry, inObs, inMaxRows, inMaxCols, &total, inVisited)
		if !goNext {
			return false
		}

		if appendPos {
			// uses a edge test (using hashed string), could have also used cyclic vertices test
			lastPoint := visited[len(visited)-1]
			key := fmt.Sprintf("%v:%v-%v:%v", lastPoint.X, lastPoint.Y, inSentry.pos.X, inSentry.pos.Y)
			if _, seen := edgeMap[key]; !seen {
				edgeMap[key] = true
			} else {
				return true
			}
			visited = append(visited, inSentry.pos)
		}
	}
}

func RunDay6Part2() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	fmt.Println("Day 6 Part 2 puzzle: Running")
	sum := 0

	// load map, find sentry location and direction
	cols := 0
	rows := 0
	var sentry Sentry
	var sentryInitial Sentry
	var obstacles []utils.Point
	var path []Sentry
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
					sentryInitial = Sentry{
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

	// fetch all points of path the sentry takes
	for {
		currentSentry := Sentry{
			pos: utils.Point{
				X: sentry.pos.X,
				Y: sentry.pos.Y,
			},
			direction: sentry.direction,
		}
		goNext, appendPos := moveSentry(&sentry, obstacles, rows, cols, &sum, &visted)
		if appendPos {
			path = append(path, currentSentry)
		}
		if !goNext {
			break
		}
	}

	// for every point in path, add that point to obstables list and test for loop
	noOfLoops := 0
	successMap := make(map[string]bool)
	for i := 0; i < len(path); i++ {
		pos := path[i].pos
		if sentryInitial.pos.X == pos.X && sentryInitial.pos.Y == pos.Y {
			continue
		}

		// reset sentry position
		sentry.pos.X = sentryInitial.pos.X
		sentry.pos.Y = sentryInitial.pos.Y
		sentry.direction = sentryInitial.direction

		dupObstacles := make([]utils.Point, len(obstacles))
		copy(dupObstacles, obstacles)
		dupObstacles = append(dupObstacles, pos)
		if testSentryLoop(&sentry, dupObstacles, rows, cols, &visted) {
			label := fmt.Sprintf("%v-%v", pos.X, pos.Y)
			if _, found := successMap[label]; !found {
				successMap[label] = true
				noOfLoops++
			}
		}
	}

	fmt.Println("Day 6 Part 2 puzzle: Result = ", noOfLoops)
}
