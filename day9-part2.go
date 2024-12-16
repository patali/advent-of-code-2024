package main

import (
	"AdventOfCode/utils"
	"fmt"
	"time"
)

func defrag2(inFiles, inFreeBlocks []DiskBlock, inBlocks []int) {
	for i := len(inFiles) - 1; i >= 0; i-- {
		file := inFiles[i]
		for j := 0; j < len(inFreeBlocks); j++ {
			freeBlock := inFreeBlocks[j]
			if file.size <= freeBlock.size && file.index > freeBlock.index {
				// move data block to the free block
				for x := freeBlock.index; x < freeBlock.index+file.size; x++ {
					inBlocks[x] = file.id
				}
				// free the file block
				for x := file.index; x < file.index+file.size; x++ {
					inBlocks[x] = -1
				}
				// update free block
				freeBlock.index += file.size
				freeBlock.size -= file.size
				if freeBlock.size == 0 {
					inFreeBlocks = append(inFreeBlocks[:j], inFreeBlocks[j+1:]...)
				} else {
					inFreeBlocks[j] = freeBlock
				}
				break
			}
		}
	}
}

func RunDay9Part2() {
	start := time.Now()
	defer utils.PrintTimeElapsed(start, "This")

	fmt.Println("Day 9 Part 2 puzzle: Running")
	sum := 0

	// load the inputs
	data, files, free := loadDay9Input()
	defrag2(files, free, data)
	sum = checksum(data)

	fmt.Println("Day 9 Part 2 puzzle: Result = ", sum)
}
