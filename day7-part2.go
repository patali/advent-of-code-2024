package main

import (
	"AdventOfCode/utils"
	"fmt"
	"time"
)

func RunDay7Part2() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	fmt.Println("Day 7 Part 2 puzzle: Running")
	sum := 0

	// load the inputs
	data := loadDay7Input()

	for i := 0; i < len(data); i++ {
		initialTotal := data[i].digits[0]
		if testExpression(data[i], initialTotal, 1, OPR_ADD, true) {
			sum += data[i].solution
		} else if testExpression(data[i], initialTotal, 1, OPR_MUL, true) {
			sum += data[i].solution
		} else if testExpression(data[i], initialTotal, 1, OPR_CONCAT, true) {
			sum += data[i].solution
		}
	}

	fmt.Println("Day 7 Part 2 puzzle: Result = ", sum)
}
