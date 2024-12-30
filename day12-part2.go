package main

import (
	"AdventOfCode/utils"
	"fmt"
	"time"
)

func calcAreaAndPerimeter2(inRegion []Plant) (int, int) {
	area := len(inRegion)
	edges := 0
	for i := 0; i < len(inRegion); i++ {
		edges += inRegion[i].cornerCount
	}
	return area, edges
}

func RunDay12Part2() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	fmt.Println("Day 12 Part 2 puzzle: Running")
	sum := 0

	data, cols, rows, _ := utils.Load2DStringArray("./input/day12.txt")
	seen := make(map[string]bool)
	regions := make(map[string][][]Plant)
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			plantPos := utils.Point{X: x, Y: y}
			plantType := getPlantAt(plantPos, data, rows, cols)
			plant := Plant{X: x, Y: y, score: 4}
			region := []Plant{plant}
			calcRegions(plantType, &region, data, rows, cols, &seen, false)
			if len(region) > 0 {
				regions[plantType] = append(regions[plantType], region)
			}
		}
	}

	for _, regions := range regions {
		for i := 0; i < len(regions); i++ {
			area, perimeter := calcAreaAndPerimeter2(regions[i])
			sum += area * perimeter
		}
	}

	fmt.Println("Day 12 Part 2 puzzle: Result = ", sum)
}
