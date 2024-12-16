package main

import (
	"AdventOfCode/utils"
	"fmt"
	"log"
	"time"
)

func loadDay10Input(inFile string) ([][]int, []utils.Point) {
	data, _, _, err := utils.Load2DIntArray("./input/day10.txt")
	if err != nil {
		log.Fatal("Failed to open data from Day 10 input")
	}

	// find trail heads
	trailHeads := make([]utils.Point, 0)
	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data[0]); x++ {
			if data[y][x] == 0 {
				trailHeads = append(trailHeads, utils.Point{
					X: x,
					Y: y,
				})
			}
		}
	}

	return data, trailHeads
}

func findTrail(inData [][]int, inLastHead, inHead utils.Point, inNextMarker int, inpTotal *int, inpSeen *map[string]bool, inDisableSeen bool) {
	currLoc := inData[inHead.Y][inHead.X]
	if currLoc == 9 {
		if !inDisableSeen {
			label := fmt.Sprintf("%v:%v-%v:%v", inLastHead.X, inLastHead.Y, inHead.X, inHead.Y)
			if _, seen := (*inpSeen)[label]; seen {
				return
			}
			(*inpSeen)[label] = true
		}
		(*inpTotal)++
		return
	}
	// finds next marker around the given head
	top, right, bottom, left := -1, -1, -1, -1
	if inHead.Y > 0 && inHead.Y <= len(inData)-1 {
		top = inData[inHead.Y-1][inHead.X]
	}
	if inHead.Y >= 0 && inHead.Y < len(inData)-1 {
		bottom = inData[inHead.Y+1][inHead.X]
	}
	if inHead.X > 0 && inHead.X <= len(inData[0])-1 {
		left = inData[inHead.Y][inHead.X-1]
	}
	if inHead.X >= 0 && inHead.X < len(inData[0])-1 {
		right = inData[inHead.Y][inHead.X+1]
	}

	if top == inNextMarker {
		findTrail(inData, inLastHead, utils.Point{
			X: inHead.X,
			Y: inHead.Y - 1,
		}, inNextMarker+1, inpTotal, inpSeen, inDisableSeen)
	}
	if right == inNextMarker {
		findTrail(inData, inLastHead, utils.Point{
			X: inHead.X + 1,
			Y: inHead.Y,
		}, inNextMarker+1, inpTotal, inpSeen, inDisableSeen)
	}
	if bottom == inNextMarker {
		findTrail(inData, inLastHead, utils.Point{
			X: inHead.X,
			Y: inHead.Y + 1,
		}, inNextMarker+1, inpTotal, inpSeen, inDisableSeen)
	}
	if left == inNextMarker {
		findTrail(inData, inLastHead, utils.Point{
			X: inHead.X - 1,
			Y: inHead.Y,
		}, inNextMarker+1, inpTotal, inpSeen, inDisableSeen)
	}
}

func RunDay10Part1() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	fmt.Println("Day 10 Part 1 puzzle: Running")
	sum := 0
	data, trailHeads := loadDay10Input("./input/day10.txt")
	seen := make(map[string]bool)
	for i := 0; i < len(trailHeads); i++ {
		score := 0
		findTrail(data, trailHeads[i], trailHeads[i], 1, &score, &seen, false)
		sum += score
	}

	fmt.Println("Day 10 Part 1 puzzle: Result = ", sum)
}
