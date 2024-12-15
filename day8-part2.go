package main

import (
	"AdventOfCode/utils"
	"fmt"
	"slices"
	"time"
)

func antiNodesInDirection(inInitPos utils.Point, inDir utils.Point, inRows, inCols int) []utils.Point {
	var antiNodes []utils.Point
	var seen []string
	initPos := utils.Point{
		X: inInitPos.X,
		Y: inInitPos.Y,
	}

	for {
		newPos := utils.Point{
			X: initPos.X + inDir.X,
			Y: initPos.Y + inDir.Y,
		}

		if newPos.X > 0 && newPos.X <= inCols && newPos.Y > 0 && newPos.Y <= inRows {
			label := fmt.Sprintf("%x-%x", newPos.X, newPos.Y)
			if !slices.Contains(seen, label) {
				seen = append(seen, label)
				antiNodes = append(antiNodes, newPos)
			}
			initPos.X = newPos.X
			initPos.Y = newPos.Y
		} else {
			break
		}
	}
	return antiNodes
}

func resonantAntiNodes(inA, inB utils.Point, inRows, inCols int) []utils.Point {
	dx := inA.X - inB.X
	dy := inA.Y - inB.Y

	var result []utils.Point
	result = append(result, antiNodesInDirection(inA, utils.Point{X: dx, Y: dy}, inRows, inCols)...)
	result = append(result, antiNodesInDirection(inB, utils.Point{X: -dx, Y: -dy}, inRows, inCols)...)

	result = append(result, inA)
	result = append(result, inB)
	return result
}

func RunDay8Part2() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	fmt.Println("Day 8 Part 2 puzzle: Running")
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

				antiNodes := resonantAntiNodes(a, b, rows, cols)

				for _, antiNode := range antiNodes {
					label := fmt.Sprintf("%x-%x", antiNode.X, antiNode.Y)
					if !slices.Contains(seen, label) {
						seen = append(seen, label)
						sum++
					}
				}
			}
		}
	}

	fmt.Println("Day 8 Part 2 puzzle: Result = ", sum)
}
