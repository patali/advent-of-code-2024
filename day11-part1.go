package main

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

func loadDay11Input(inFilename string) []string {
	// load file
	file, err := os.Open(inFilename)
	if err != nil {
		log.Fatal("Failed to fetch Day1 input")
	}
	defer file.Close()

	r, _ := regexp.Compile(`-?\d+`)
	var nums []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strNums := r.FindAllString(scanner.Text(), -1)
		nums = append(nums, strNums...)
	}
	return nums
}

func blink(inStone string, inDepth int, inpSeen *map[string]int, inMaxDepth int) int {
	noOfStones := 0
	if inDepth > inMaxDepth {
		return 1
	}
	seenKey := fmt.Sprintf("%v-%v", inStone, inDepth)
	if val, ok := (*inpSeen)[seenKey]; ok {
		return val
	}

	if inStone == "0" {
		noOfStones = blink("1", inDepth+1, inpSeen, inMaxDepth)
		(*inpSeen)[seenKey] = noOfStones
	} else if len(inStone)%2 == 0 {
		leftStone := inStone[:len(inStone)/2]
		leftStoneVal := utils.StrToInt(leftStone)
		leftStone = fmt.Sprintf("%v", leftStoneVal)

		rightStone := inStone[len(inStone)/2:]
		rightStoneVal := utils.StrToInt(rightStone)
		rightStone = fmt.Sprintf("%v", rightStoneVal)

		noOfStones = blink(leftStone, inDepth+1, inpSeen, inMaxDepth) + blink(rightStone, inDepth+1, inpSeen, inMaxDepth)
		(*inpSeen)[seenKey] = noOfStones
	} else {
		val := utils.StrToInt(inStone)
		strVal := fmt.Sprintf("%v", val*2024)
		strVal = strings.Trim(strVal, " ") // trim whitespace
		noOfStones = blink(strVal, inDepth+1, inpSeen, inMaxDepth)
		(*inpSeen)[seenKey] = noOfStones
	}

	return noOfStones
}

func RunDay11Part1() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	fmt.Println("Day 11 Part 1 puzzle: Running")
	sum := 0

	stones := loadDay11Input("./input/day11.txt")
	seen := make(map[string]int)
	for i := 0; i < len(stones); i++ {
		sum += blink(stones[i], 1, &seen, 25)
	}

	fmt.Println("Day 11 Part 1 puzzle: Result = ", sum)
}
