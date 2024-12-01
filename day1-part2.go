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

func RunDay1Part2() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	fmt.Println("Day 1 Part2 puzzle: Running")
	sum := 0

	// load file
	file, err := os.Open("./input/day1.txt")
	if err != nil {
		log.Fatal("Failed to fetch Day1 input")
	}
	defer file.Close()

	leftList := []int{}
	similarityMap := make(map[int]int)

	// process one line at a time
	r, _ := regexp.Compile(`-?\d+`) // regex to match numbers
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := r.FindAllString(scanner.Text(), -1)
		if len(data) == 2 {
			leftNum := utils.StrToInt(data[0])
			rightNum := utils.StrToInt(data[1])
			leftList = append(leftList, leftNum)

			similarity, ok := similarityMap[rightNum]
			if ok {
				similarityMap[rightNum] = similarity + 1
			} else {
				similarityMap[rightNum] = 1
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// calculate distance sum
	for i := 0; i < len(leftList); i++ {
		leftNum := leftList[i]
		if similarity, found := similarityMap[leftNum]; found {
			sum += leftList[i] * similarity
		}
	}

	fmt.Println("Day 1 Part2 puzzle: Result = ", sum)
}
