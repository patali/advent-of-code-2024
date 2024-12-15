package main

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"time"
)

func loadDay8Input() (map[string][]utils.Point, int, int) {
	file, err := os.Open("./input/day8.txt")
	if err != nil {
		log.Fatal("Failed to fetch Day 8 input")
	}
	defer file.Close()

	data := make(map[string][]utils.Point)

	scanner := bufio.NewScanner(file)
	j := 0
	cols := 0
	for scanner.Scan() {
		line := scanner.Text()
		cols = len(line)
		for i := 0; i < len(line); i++ {
			freq := string(line[i])
			if freq == "." {
				continue
			}
			pos := utils.Point{
				X: i + 1,
				Y: j + 1,
			}
			if list, found := data[freq]; found {
				list = append(list, pos)
				data[freq] = list
			} else {
				data[freq] = make([]utils.Point, 0)
				data[freq] = append(data[freq], pos)
			}
		}
		j++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("File read error.")
	}
	return data, j, cols
}

func calcAntinodes(inA, inB utils.Point) (utils.Point, utils.Point) {
	x3 := 2*inA.X - inB.X
	y3 := 2*inA.Y - inB.Y

	x4 := 2*inB.X - inA.X
	y4 := 2*inB.Y - inA.Y
	return utils.Point{X: x3, Y: y3}, utils.Point{X: x4, Y: y4}
}

func RunDay8Part1() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	fmt.Println("Day 8 Part 1 puzzle: Running")
	sum := 0

	// load the inputs
	data, rows, cols := loadDay8Input()
	var seen []string

	// find antinodes
	for _, list := range data {
		for i := 0; i < len(list); i++ {
			for j := i + 1; j < len(list); j++ {
				if i == j {
					continue
				}
				a := list[i]
				b := list[j]
				antiNode1, antiNode2 := calcAntinodes(a, b)
				if antiNode1.X > 0 && antiNode1.X <= cols && antiNode1.Y > 0 && antiNode1.Y <= rows {
					label := fmt.Sprintf("%x-%x", antiNode1.X, antiNode1.Y)
					if !slices.Contains(seen, label) {
						seen = append(seen, label)
						sum++
					}
				}
				if antiNode2.X > 0 && antiNode2.X <= cols && antiNode2.Y > 0 && antiNode2.Y <= rows {
					label := fmt.Sprintf("%x-%x", antiNode2.X, antiNode2.Y)
					if !slices.Contains(seen, label) {
						seen = append(seen, label)
						sum++
					}
				}
			}
		}
	}

	fmt.Println("Day 8 Part 1 puzzle: Result = ", sum)
}
