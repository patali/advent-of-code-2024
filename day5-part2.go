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

func swapPages(inArr []string, inI, inJ int) {
	inArr[inI], inArr[inJ] = inArr[inJ], inArr[inI]
}

func RunDay5Part2() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	fmt.Println("Day 5 Part 2 puzzle: Running")
	sum := 0
	seen := make(map[string]bool)

	// load file
	file, err := os.Open("./input/day5.txt")
	if err != nil {
		log.Fatal("Failed to fetch Day 2 input")
	}
	defer file.Close()

	numMatcher, _ := regexp.Compile(`-?\d+`)

	// load orders and updates
	orders := make([]StringPair, 0)
	updates := make([][]string, 0)
	orderFetch := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			orderFetch = false
			continue
		}

		if orderFetch {
			pair := numMatcher.FindAllString(line, -1)
			orders = append(orders, StringPair{pair[0], pair[1]})
		} else {
			nums := numMatcher.FindAllString(line, -1)
			updates = append(updates, nums)
		}
	}

	for i := 0; i < len(updates); i++ {
		update := updates[i]
		fullClean := true
		for {
			success := true
			retry := false
			for x := 0; x < len(update); x++ {
				for y := x; y < len(update); y++ {
					if x == y {
						continue // no need to check itself
					}
					pageA := update[x]
					pageB := update[y]
					if !isABeforeB(&orders, &seen, pageA, pageB) {
						swapPages(update, x, y)
						success = false
						retry = true
						fullClean = false
						break
					}
				}
				if !success {
					break
				}
			}

			if success && !fullClean && !retry {
				sum += utils.StrToInt(update[len(update)/2])
			}

			if !retry {
				break
			}
		}
	}

	fmt.Println("Day 5 Part 2 puzzle: Result = ", sum)
}
