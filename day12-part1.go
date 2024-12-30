package main

import (
	"AdventOfCode/utils"
	"fmt"
	"time"
)

type Plant struct {
	X           int
	Y           int
	score       int
	cornerCount int
}

func getPlantAt(inPoint utils.Point, inData []string, inRows, inCols int) string {
	if inPoint.X < 0 || inPoint.X > inCols-1 || inPoint.Y < 0 || inPoint.Y > inRows-1 {
		return "#"
	}
	return string(inData[inPoint.Y][inPoint.X])
}

func findCorners(inPlantIndex int, inRegion *[]Plant, inData []string, inRows, inCols int) int {
	plant := (*inRegion)[inPlantIndex]
	plantType := getPlantAt(utils.Point{X: plant.X, Y: plant.Y}, inData, inRows, inCols)

	topPlant := getPlantAt(utils.Point{X: plant.X, Y: plant.Y - 1}, inData, inRows, inCols)
	top := topPlant == plantType
	bottomPlant := getPlantAt(utils.Point{X: plant.X, Y: plant.Y + 1}, inData, inRows, inCols)
	bottom := bottomPlant == plantType
	leftPlant := getPlantAt(utils.Point{X: plant.X - 1, Y: plant.Y}, inData, inRows, inCols)
	left := leftPlant == plantType
	rightPlant := getPlantAt(utils.Point{X: plant.X + 1, Y: plant.Y}, inData, inRows, inCols)
	right := rightPlant == plantType
	topRightPlant := getPlantAt(utils.Point{X: plant.X + 1, Y: plant.Y - 1}, inData, inRows, inCols)
	topRight := topRightPlant == plantType
	bottomRightPlant := getPlantAt(utils.Point{X: plant.X + 1, Y: plant.Y + 1}, inData, inRows, inCols)
	bottomRight := bottomRightPlant == plantType
	bottomLeftPlant := getPlantAt(utils.Point{X: plant.X - 1, Y: plant.Y + 1}, inData, inRows, inCols)
	bottomLeft := bottomLeftPlant == plantType
	topLeftPlant := getPlantAt(utils.Point{X: plant.X - 1, Y: plant.Y - 1}, inData, inRows, inCols)
	topLeft := topLeftPlant == plantType

	if !top && !right && !bottom && !left {
		return 4
	}

	if bottom && !right && !top && !left {
		return 2
	}
	if top && !right && !bottom && !left {
		return 2
	}
	if right && !bottom && !left && !top {
		return 2
	}
	if left && !right && !bottom && !top {
		return 2
	}

	if top && right && left && bottom {
		count := 0
		if !topRight {
			count++
		}
		if !topLeft {
			count++
		}
		if !bottomRight {
			count++
		}
		if !bottomLeft {
			count++
		}

		if count > 0 {
			return count
		}
	}

	if top && right && bottom {
		count := 0
		if !topRight {
			count++
		}
		if !bottomRight {
			count++
		}
		return count
	}
	if top && left && bottom {
		count := 0
		if !topLeft {
			count++
		}
		if !bottomLeft {
			count++
		}
		return count
	}
	if right && bottom && left {
		count := 0
		if !bottomRight {
			count++
		}
		if !bottomLeft {
			count++
		}
		return count
	}
	if right && top && left {
		count := 0
		if !topRight {
			count++
		}
		if !topLeft {
			count++
		}
		return count
	}

	if left && top {
		if topLeft {
			return 1
		}
		return 2
	}
	if right && top {
		if topRight {
			return 1
		}
		return 2
	}
	if left && bottom {
		if bottomLeft {
			return 1
		}
		return 2
	}
	if right && bottom {
		if bottomRight {
			return 1
		}
		return 2
	}

	return 0
}

func calcRegions(inType string, inRegion *[]Plant, inData []string, inRows, inCols int, inpSeen *map[string]bool, inDisableCornerCount bool) {
	testPlantIndex := len(*inRegion) - 1
	testPlant := (*inRegion)[testPlantIndex]

	// do some caching checks
	label := fmt.Sprintf("%v-%v", testPlant.X, testPlant.Y)
	if _, seen := (*inpSeen)[label]; seen {
		// remove last point
		*inRegion = (*inRegion)[:testPlantIndex]
		return
	}
	(*inpSeen)[label] = true

	// check if we are at the edge
	if testPlant.X < 0 || testPlant.X > inCols-1 || testPlant.Y < 0 || testPlant.Y > inRows-1 {
		return
	}

	topPos := utils.Point{X: testPlant.X, Y: testPlant.Y - 1}
	topPlant := getPlantAt(topPos, inData, inRows, inCols)
	hasTopPlant := topPlant == inType
	if hasTopPlant {
		(*inRegion)[testPlantIndex].score-- // reduce score of testPlant by 1
		topPlant := Plant{X: topPos.X, Y: topPos.Y, score: 4}
		*inRegion = append(*inRegion, topPlant)
		calcRegions(inType, inRegion, inData, inRows, inCols, inpSeen, inDisableCornerCount)
	}

	bottomPos := utils.Point{X: testPlant.X, Y: testPlant.Y + 1}
	bottomPlant := getPlantAt(bottomPos, inData, inRows, inCols)
	hasBottomPlant := bottomPlant == inType
	if hasBottomPlant {
		(*inRegion)[testPlantIndex].score-- // reduce score of testPlant by 1
		bottomPlant := Plant{X: bottomPos.X, Y: bottomPos.Y, score: 4}
		*inRegion = append(*inRegion, bottomPlant)
		calcRegions(inType, inRegion, inData, inRows, inCols, inpSeen, inDisableCornerCount)
	}

	leftPos := utils.Point{X: testPlant.X - 1, Y: testPlant.Y}
	leftPlant := getPlantAt(leftPos, inData, inRows, inCols)
	hasLeftPlant := leftPlant == inType
	if hasLeftPlant {
		(*inRegion)[testPlantIndex].score-- // reduce score of testPlant by 1
		leftPlant := Plant{X: leftPos.X, Y: leftPos.Y, score: 4}
		*inRegion = append(*inRegion, leftPlant)
		calcRegions(inType, inRegion, inData, inRows, inCols, inpSeen, inDisableCornerCount)
	}

	rightPos := utils.Point{X: testPlant.X + 1, Y: testPlant.Y}
	rightPlant := getPlantAt(rightPos, inData, inRows, inCols)
	hasRightPlant := rightPlant == inType
	if hasRightPlant {
		(*inRegion)[testPlantIndex].score-- // reduce score of testPlant by 1
		rightPlant := Plant{X: rightPos.X, Y: rightPos.Y, score: 4}
		*inRegion = append(*inRegion, rightPlant)
		calcRegions(inType, inRegion, inData, inRows, inCols, inpSeen, inDisableCornerCount)
	}

	// lone region
	if !hasTopPlant && !hasBottomPlant && !hasLeftPlant && !hasRightPlant {
		(*inRegion)[testPlantIndex].cornerCount = 4
	} else if !hasTopPlant && hasBottomPlant && !hasLeftPlant && hasRightPlant {
		(*inRegion)[testPlantIndex].cornerCount = 1
	} else if !hasTopPlant && hasBottomPlant && hasLeftPlant && !hasRightPlant {
		(*inRegion)[testPlantIndex].cornerCount = 1
	} else if hasTopPlant && !hasBottomPlant && !hasLeftPlant && hasRightPlant {
		(*inRegion)[testPlantIndex].cornerCount = 1
	} else if hasTopPlant && !hasBottomPlant && hasLeftPlant && !hasRightPlant {
		(*inRegion)[testPlantIndex].cornerCount = 1
	} else if hasTopPlant && !hasBottomPlant && !hasLeftPlant && !hasRightPlant {
		(*inRegion)[testPlantIndex].cornerCount = 2
	} else if !hasTopPlant && hasBottomPlant && !hasLeftPlant && !hasRightPlant {
		(*inRegion)[testPlantIndex].cornerCount = 2
	} else if !hasTopPlant && !hasBottomPlant && hasLeftPlant && !hasRightPlant {
		(*inRegion)[testPlantIndex].cornerCount = 2
	} else if !hasTopPlant && !hasBottomPlant && !hasLeftPlant && hasRightPlant {
		(*inRegion)[testPlantIndex].cornerCount = 2
	}
	if !inDisableCornerCount {
		corners := findCorners(testPlantIndex, inRegion, inData, inRows, inCols)
		(*inRegion)[testPlantIndex].cornerCount = corners
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

	data, cols, rows, _ := utils.Load2DStringArray("./input/day12.txt")
	seen := make(map[string]bool)
	regions := make(map[string][][]Plant)
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			plantPos := utils.Point{X: x, Y: y}
			plantType := getPlantAt(plantPos, data, rows, cols)
			plant := Plant{X: x, Y: y, score: 4}
			region := []Plant{plant}
			calcRegions(plantType, &region, data, rows, cols, &seen, true)
			if len(region) > 0 {
				regions[plantType] = append(regions[plantType], region)
			}
		}
	}

	for _, regions := range regions {
		for i := 0; i < len(regions); i++ {
			area, perimeter := calcAreaAndPerimeter(regions[i])
			sum += area * perimeter
		}
	}

	fmt.Println("Day 12 Part 1 puzzle: Result = ", sum)
}
