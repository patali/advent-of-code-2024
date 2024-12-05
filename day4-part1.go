package main

import (
	"AdventOfCode/utils"
	"fmt"
	"log"
	"time"
)

type Data struct {
	data []string
	rows int
	cols int
}

type DirectionType int

const (
	TOP DirectionType = iota + 1
	TOP_RIGHT
	RIGHT
	BOTTOM_RIGHT
	BOTTOM
	BOTTOM_LEFT
	LEFT
	TOP_LEFT
)

type Direction struct {
	t DirectionType
	x int
	y int
}

var (
	DIRECTIONS = map[DirectionType]Direction{
		TOP:          {TOP, 0, -1},
		TOP_RIGHT:    {TOP_RIGHT, 1, -1},
		RIGHT:        {RIGHT, 1, 0},
		BOTTOM_RIGHT: {BOTTOM_RIGHT, 1, 1},
		BOTTOM:       {BOTTOM, 0, 1},
		BOTTOM_LEFT:  {BOTTOM_LEFT, -1, 1},
		LEFT:         {LEFT, -1, 0},
		TOP_LEFT:     {TOP_LEFT, -1, -1},
	}
)

func testBounds(inRow, inCol, inMaxRow, inMaxCol int, inDirection DirectionType) bool {
	switch inDirection {
	case TOP:
		return inRow >= 3
	case TOP_RIGHT:
		return inRow >= 3 && inCol <= inMaxCol-4
	case RIGHT:
		return inCol <= inMaxCol-4
	case BOTTOM_RIGHT:
		return inRow <= inMaxRow-4 && inCol <= inMaxCol-4
	case BOTTOM:
		return inRow <= inMaxRow-4
	case BOTTOM_LEFT:
		return inRow <= inMaxRow-4 && inCol >= 3
	case LEFT:
		return inCol >= 3
	case TOP_LEFT:
		return inRow >= 3 && inCol >= 3
	}
	return false
}

func matchXMAS(inData *Data, inRow, inCol int, inDirection Direction) int {
	// check bounds
	if testBounds(inRow, inCol, inData.rows, inData.cols, inDirection.t) {
		// check pattern in the direction
		if string(inData.data[inRow][inCol]) == "X" {
			if string(inData.data[inRow+inDirection.y][inCol+inDirection.x]) == "M" {
				if string(inData.data[inRow+inDirection.y*2][inCol+inDirection.x*2]) == "A" {
					if string(inData.data[inRow+inDirection.y*3][inCol+inDirection.x*3]) == "S" {
						return 1
					}
				}
			}
		}
	}

	return 0
}

func RunDay4Part1() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	fmt.Println("Day 4 Part 1 puzzle: Running")
	sum := 0
	block, cols, rows, err := utils.Load2DStringArray("./input/day4.txt")
	if err != nil {
		log.Fatal("Failed to open data from Day 4 input")
	}

	data := Data{
		data: block,
		rows: rows,
		cols: cols,
	}

	for y := 0; y < data.rows; y++ {
		for x := 0; x < data.cols; x++ {
			for _, val := range DIRECTIONS {
				sum += matchXMAS(&data, y, x, val)
			}
		}
	}

	fmt.Println("Day 4 Part 1 puzzle: Result = ", sum)
}
