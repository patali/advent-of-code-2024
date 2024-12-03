package main

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"time"
)

func RunDay2Part1() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	fmt.Println("Day 2 Part1 puzzle: Running")
	sum := 0

	// load file
	file, err := os.Open("./input/day2.txt")
	if err != nil {
		log.Fatal("Failed to fetch Day2 input")
	}
	defer file.Close()

	// regex to match numbers
	r, _ := regexp.Compile(`-?\d+`)

	// process one line at a time
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := r.FindAllString(scanner.Text(), -1)
		success := 1
		isLevelsIncreasing := true
		for i := 0; i < len(data)-1; i++ {
			a := utils.StrToInt(data[i])
			b := utils.StrToInt(data[i+1])
			diff := a - b
			diffVal := math.Abs(float64(diff))

			if diff == 0 || diffVal > 3 { // fail conditions
				success = 0
				break
			}

			if i == 0 {
				isLevelsIncreasing = diff > 0
			} else {
				if isLevelsIncreasing != (diff > 0) { // trend mismatch
					success = 0
					break
				}
			}
		}

		sum += success
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Day 2 Part1 puzzle: Result = ", sum)
}
