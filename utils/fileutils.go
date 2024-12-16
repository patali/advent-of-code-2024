package utils

import (
	"bufio"
	"os"
)

func Load2DStringArray(filePath string) ([]string, int, int, error) {
	xCount := 0
	yCount := 0
	file, err := os.Open(filePath)
	if err != nil {
		return nil, 0, 0, err
	}
	defer file.Close()

	var result []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			result = append(result, line)
			xCount = len(line)
		}
		yCount++
	}

	if err := scanner.Err(); err != nil {
		return nil, 0, 0, err
	}

	return result, xCount, yCount, nil
}

func Load2DIntArray(filePath string) ([][]int, int, int, error) {
	xCount := 0
	yCount := 0
	file, err := os.Open(filePath)
	if err != nil {
		return nil, 0, 0, err
	}
	defer file.Close()

	var result [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineArray := make([]int, len(line))
		for i := 0; i < len(line); i++ {
			lineArray[i] = StrToInt(string(line[i]))
		}
		result = append(result, lineArray)
		xCount = len(line)
		yCount++
	}

	if err := scanner.Err(); err != nil {
		return nil, 0, 0, err
	}

	return result, xCount, yCount, nil
}
