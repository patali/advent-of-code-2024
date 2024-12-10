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

type Day7Input struct {
	solution int
	digits   []int
}

type OperationType int

const (
	OPR_ADD OperationType = iota + 1
	OPR_MUL
	OPR_CONCAT
)

func loadDay7Input() []Day7Input {
	file, err := os.Open("./input/day7.txt")
	if err != nil {
		log.Fatal("Failed to fetch Day 7 input")
	}
	defer file.Close()

	r, _ := regexp.Compile(`-?\d+`)
	var data []Day7Input

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers := r.FindAllString(scanner.Text(), -1)
		solution := utils.StrToInt(numbers[0])
		var digits []int
		for i := 1; i < len(numbers); i++ {
			digits = append(digits, utils.StrToInt(numbers[i]))
		}
		input := Day7Input{
			solution: solution,
			digits:   digits,
		}
		data = append(data, input)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("File read error.")
	}

	return data
}

func applyOperation(inA, inB int, inType OperationType) int {
	switch inType {
	case OPR_ADD:
		return inA + inB
	case OPR_MUL:
		return inA * inB
	case OPR_CONCAT:
		return utils.StrToInt(fmt.Sprintf("%v%v", inA, inB))
	}
	return 0
}

func testExpression(inInput Day7Input, inCurrenTotal int, inIndex int, inType OperationType, inEnableConcat bool) bool {
	// termination condition
	if inIndex == len(inInput.digits) {
		return inInput.solution == inCurrenTotal
	}

	// update current total
	inCurrenTotal = applyOperation(inCurrenTotal, inInput.digits[inIndex], inType)

	// go further down the expression tree
	if testExpression(inInput, inCurrenTotal, inIndex+1, OPR_ADD, inEnableConcat) {
		return true
	} else if testExpression(inInput, inCurrenTotal, inIndex+1, OPR_MUL, inEnableConcat) {
		return true
	} else if inEnableConcat && testExpression(inInput, inCurrenTotal, inIndex+1, OPR_CONCAT, inEnableConcat) {
		return true
	}
	return false
}

func RunDay7Part1() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	fmt.Println("Day 7 Part 1 puzzle: Running")
	sum := 0

	// load the inputs
	data := loadDay7Input()

	for i := 0; i < len(data); i++ {
		initialTotal := data[i].digits[0]
		if testExpression(data[i], initialTotal, 1, OPR_ADD, false) {
			sum += data[i].solution
		} else if testExpression(data[i], initialTotal, 1, OPR_MUL, false) {
			sum += data[i].solution
		}
	}

	fmt.Println("Day 7 Part 1 puzzle: Result = ", sum)
}
