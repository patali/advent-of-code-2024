package main

import (
	"AdventOfCode/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

type DiskBlock struct {
	id    int
	index int
	size  int
}

func loadDay9Input() ([]int, []DiskBlock, []DiskBlock) {
	file, err := os.Open("./input/day9.txt")
	if err != nil {
		log.Fatal("Failed to fetch Day 9 input")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var data string
	for scanner.Scan() {
		data = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("File read error.")
	}

	var blocks []int
	var usedBlocks []DiskBlock
	var freeBlocks []DiskBlock
	index := 0
	for i := 0; i < len(data); i++ {
		size := utils.StrToInt(string(data[i]))
		if i%2 == 0 {
			usedBlocks = append(usedBlocks, DiskBlock{id: index, index: len(blocks), size: size})
			blocks = append(blocks, utils.InitArray(size, index)...)
			index++
		} else {
			freeBlocks = append(freeBlocks, DiskBlock{id: -1, index: len(blocks), size: size})
			blocks = append(blocks, utils.InitArray(size, -1)...)
		}
	}

	return blocks, usedBlocks, freeBlocks
}

func defrag(inBlocks []int) {
	count := len(inBlocks)
	x := 0
	y := count - 1
	for i := 0; i < count; i++ {
		a := inBlocks[x]
		b := inBlocks[y]
		if a == -1 && b == -1 {
			y--
		} else if a == -1 && b > -1 {
			inBlocks[x] = b
			inBlocks[y] = a
			x++
			y--
		} else if b == -1 {
			y--
		} else {
			x++
		}
		if x > y {
			break
		}
	}
}

func checksum(inBlocks []int) int {
	sum := 0
	for i := 0; i < len(inBlocks); i++ {
		if inBlocks[i] == -1 {
			continue
		}
		sum += i * inBlocks[i]
	}
	return sum
}

func RunDay9Part1() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	fmt.Println("Day 9 Part 1 puzzle: Running")
	sum := 0

	// load the inputs
	blocks, _, _ := loadDay9Input()
	defrag(blocks)
	sum = checksum(blocks)

	fmt.Println("Day 9 Part 1 puzzle: Result = ", sum)
}
