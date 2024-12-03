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

func testFunc(inData []string) bool {
	isLevelsIncreasing := true
	for i := 0; i < len(inData)-1; i++ {
		a := utils.StrToInt(inData[i])
		b := utils.StrToInt(inData[i+1])
		diff := a - b
		diffVal := math.Abs(float64(diff))

		if diff == 0 || diffVal > 3 { // fail conditions
			return false
		}

		if i == 0 {
			isLevelsIncreasing = diff > 0
		} else {
			if isLevelsIncreasing != (diff > 0) { // trend mismatch
				return false
			}
		}
	}
	return true
}

func RunDay2Part2() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	fmt.Println("Day 2 Part2 puzzle: Running")
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
		success := testFunc(data)
		if success {
			sum++
		} else {
			for x := 0; x < len(data); x++ {
				// deep copy and remove the x'th element
				sub_data := make([]string, len(data))
				copy(sub_data, data)
				sub_data = append(sub_data[:x], sub_data[x+1:]...)

				sub_success := testFunc(sub_data)
				if sub_success {
					sum++
					break
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Day 2 Part2 puzzle: Result = ", sum)
}
