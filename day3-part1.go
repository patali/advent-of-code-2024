package main

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"
)

func day3Part1ProcessLine(inEntries []string, inNumMatcher *regexp.Regexp) int {
	total := 0
	for i := 0; i < len(inEntries); i++ {
		nums := inNumMatcher.FindAllString(inEntries[i], -1)
		total += (utils.StrToInt(nums[0]) * utils.StrToInt(nums[1]))
	}
	return total
}

func RunDay3Part1() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	fmt.Println("Day 3 Part1 puzzle: Running")
	sum := 0

	// load file
	file, err := os.Open("./input/day3.txt")
	if err != nil {
		log.Fatal("Failed to fetch Day3 input")
	}
	defer file.Close()

	// regex to match mul numbers
	r, _ := regexp.Compile(`mul\([0-9]+,[0-9]+\)`)
	// regex to match numbers
	numMatcher, _ := regexp.Compile(`-?\d+`)

	// process one line at a time
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := r.FindAllString(scanner.Text(), -1)
		sum += day3Part1ProcessLine(data, numMatcher)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Day 3 Part1 puzzle: Result = ", sum)
}
