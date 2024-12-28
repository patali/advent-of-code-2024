package main

import (
	"AdventOfCode/utils"
	"fmt"
	"time"
)

type Plant struct {
	X     int
	Y     int
	score int
}

func getPlantAt(inPoint utils.Point, inData []string, inRows, inCols int) string {
	if inPoint.X < 0 || inPoint.X > inCols-1 || inPoint.Y < 0 || inPoint.Y > inRows-1 {
		return "."
	}
	return string(inData[inPoint.Y][inPoint.X])
}

func calcRegions(inType string, inRegion *[]Plant, inData []string, inRows, inCols int, inpSeen *map[string]bool) {
	testPlant := (*inRegion)[len(*inRegion)-1]

	// do some caching checks
	label := fmt.Sprintf("%v-%v", testPlant.X, testPlant.Y)
	if _, seen := (*inpSeen)[label]; seen {
		// remove last point
		*inRegion = (*inRegion)[:len(*inRegion)-1]
		return
	}
	(*inpSeen)[label] = true

	// check if we are at the edge
	if testPlant.X < 0 || testPlant.X > inCols-1 || testPlant.Y < 0 || testPlant.Y > inRows-1 {
		return
	}

	topPos := utils.Point{X: testPlant.X, Y: testPlant.Y - 1}
	topPlant := getPlantAt(topPos, inData, inRows, inCols)
	if topPlant == inType {
		(*inRegion)[len(*inRegion)-1].score-- // reduce score of testPlant by 1
		topPlant := Plant{X: topPos.X, Y: topPos.Y, score: 4}
		*inRegion = append(*inRegion, topPlant)
		calcRegions(inType, inRegion, inData, inRows, inCols, inpSeen)
	}

	bottomPos := utils.Point{X: testPlant.X, Y: testPlant.Y + 1}
	bottomPlant := getPlantAt(bottomPos, inData, inRows, inCols)
	if bottomPlant == inType {
		(*inRegion)[len(*inRegion)-1].score-- // reduce score of testPlant by 1
		bottomPlant := Plant{X: bottomPos.X, Y: bottomPos.Y, score: 4}
		*inRegion = append(*inRegion, bottomPlant)
		calcRegions(inType, inRegion, inData, inRows, inCols, inpSeen)
	}

	leftPos := utils.Point{X: testPlant.X - 1, Y: testPlant.Y}
	leftPlant := getPlantAt(leftPos, inData, inRows, inCols)
	if leftPlant == inType {
		(*inRegion)[len(*inRegion)-1].score-- // reduce score of testPlant by 1
		leftPlant := Plant{X: leftPos.X, Y: leftPos.Y, score: 4}
		*inRegion = append(*inRegion, leftPlant)
		calcRegions(inType, inRegion, inData, inRows, inCols, inpSeen)
	}

	rightPos := utils.Point{X: testPlant.X + 1, Y: testPlant.Y}
	rightPlant := getPlantAt(rightPos, inData, inRows, inCols)
	if rightPlant == inType {
		(*inRegion)[len(*inRegion)-1].score-- // reduce score of testPlant by 1
		rightPlant := Plant{X: rightPos.X, Y: rightPos.Y, score: 4}
		*inRegion = append(*inRegion, rightPlant)
		calcRegions(inType, inRegion, inData, inRows, inCols, inpSeen)
	}
}

func calcAreaAndPerimeter(inRegion []Plant) (int, int) {
	area := len(inRegion)
	perimeter := 0
	for i := 0; i < len(inRegion); i++ {
		perimeter += inRegion[i].score
	}
	return area, perimeter
}

func RunDay12Part1() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	fmt.Println("Day 12 Part 1 puzzle: Running")
	sum := 0

	data, rows, cols, _ := utils.Load2DStringArray("./input/day12.txt")
	seen := make(map[string]bool)
	regions := make(map[string][][]Plant)
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			plantPos := utils.Point{X: x, Y: y}
			plantType := getPlantAt(plantPos, data, rows, cols)
			plant := Plant{X: x, Y: y, score: 4}
			region := []Plant{plant}
			calcRegions(plantType, &region, data, rows, cols, &seen)
			if len(region) > 0 {
				regions[plantType] = append(regions[plantType], region)
			}
		}
	}

	for key, regions := range regions {
		for i := 0; i < len(regions); i++ {
			area, perimeter := calcAreaAndPerimeter(regions[i])
			fmt.Println(key, regions[i])
			fmt.Println(area, perimeter)
			sum += area * perimeter
		}
	}

	fmt.Println("Day 12 Part 1 puzzle: Result = ", sum)
}
