package main

import (
	"AdventOfCode/utils"
	"fmt"
	"log"
	"time"
)

func testMASBounds(inRow, inCol, inMaxRow, inMaxCol int) bool {
	return inRow >= 1 && inRow <= inMaxRow-2 && inCol >= 1 && inCol <= inMaxCol-2
}

func matchMAS(inData *Data, inRow, inCol int) int {
	if !testMASBounds(inRow, inCol, inData.rows, inData.cols) {
		return 0
	}

	if string(inData.data[inRow][inCol]) == "A" {
		tl := string(inData.data[inRow-1][inCol-1])
		tr := string(inData.data[inRow-1][inCol+1])
		bl := string(inData.data[inRow+1][inCol-1])
		br := string(inData.data[inRow+1][inCol+1])
		if tl == "M" && br == "S" || tl == "S" && br == "M" {
			if tr == "M" && bl == "S" || tr == "S" && bl == "M" {
				return 1
			}
		}
	}

	return 0
}

func RunDay4Part2() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	fmt.Println("Day 4 Part 2 puzzle: Running")
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
			sum += matchMAS(&data, y, x)
		}
	}

	fmt.Println("Day 4 Part 2 puzzle: Result = ", sum)
}
