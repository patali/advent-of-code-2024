package main

import (
	"AdventOfCode/utils"
	"fmt"
	"time"
)

func RunDay11Part2() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	fmt.Println("Day 11 Part 2 puzzle: Running")
	sum := 0

	stones := loadDay11Input("./input/day11.txt")
	seen := make(map[string]int)
	for i := 0; i < len(stones); i++ {
		sum += blink(stones[i], 1, &seen, 75)
	}

	fmt.Println("Day 11 Part 2 puzzle: Result = ", sum)
}
