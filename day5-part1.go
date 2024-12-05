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

type StringPair struct {
	a string
	b string
}

func isABeforeB(inRules *[]StringPair, inSeen *map[string]bool, inA, inB string) bool {
	if _, found := (*inSeen)[inA+"-"+inB]; found {
		return true
	}

	for i := 0; i < len(*inRules); i++ {
		rule := (*inRules)[i]
		if rule.a == inA && rule.b == inB {
			(*inSeen)[inA+"-"+inB] = true
			return true
		}
	}
	return false
}

func RunDay5Part1() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	fmt.Println("Day 5 Part 1 puzzle: Running")
	sum := 0
	seen := make(map[string]bool)

	// load file
	file, err := os.Open("./input/day5.txt")
	if err != nil {
		log.Fatal("Failed to fetch Day3 input")
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
		success := true
		for x := 0; x < len(update); x++ {
			for y := x; y < len(update); y++ {
				if x == y {
					continue // no need to check itself
				}
				pageA := update[x]
				pageB := update[y]
				if !isABeforeB(&orders, &seen, pageA, pageB) {
					success = false
					break
				}
			}
			if !success {
				break
			}
		}
		if success {
			sum += utils.StrToInt(update[len(update)/2])
		}
	}

	fmt.Println("Day 5 Part 1 puzzle: Result = ", sum)
}
