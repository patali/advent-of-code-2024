package main

import (
	"AdventOfCode/utils"
	"fmt"
	"time"
)

func RunDay10Part2() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	fmt.Println("Day 10 Part 2 puzzle: Running")
	sum := 0

	data, trailHeads := loadDay10Input("./input/day10.txt")
	seen := make(map[string]bool)
	for i := 0; i < len(trailHeads); i++ {
		score := 0
		findTrail(data, trailHeads[i], trailHeads[i], 1, &score, &seen, true)
		sum += score
	}

	fmt.Println("Day 10 Part 2 puzzle: Result = ", sum)
}
