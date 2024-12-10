package utils

import "math"

type Point struct {
	X int
	Y int
}

func ThreePointsInLine(inA, inB, inC Point) bool {
	return (inA.X-inC.X)*(inC.Y-inB.Y) == (inC.X-inB.X)*(inA.Y-inC.Y)
}

func DistanceAB(inA, inB Point) int {
	xDelta := inB.X - inA.X
	yDelta := inB.Y - inA.Y
	return int(math.Sqrt(float64(xDelta*xDelta) + float64(yDelta*yDelta)))
}
