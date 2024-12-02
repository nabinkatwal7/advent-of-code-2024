package daytwo

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkReportSafety(levels []int) bool {
	isIncreasing := true
	isDecreasing := true

	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]

		if diff == 0 {
			isIncreasing = false
			isDecreasing = false
			break
		}

		if diff < -3 || diff > 3 {
			isIncreasing = false
			isDecreasing = false
			break
		}

		if diff > 0 {
			isDecreasing = false
		}
		if diff < 0 {
			isIncreasing = false
		}
	}

	return isIncreasing || isDecreasing
}

func Part1() {
	file, err := os.Open("./day-two/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safeReportsCount := 0

	for scanner.Scan() {
		levelStrings := strings.Fields(scanner.Text())
		levels := make([]int, len(levelStrings))

		for i, levelStr := range levelStrings {
			level, err := strconv.Atoi(levelStr)
			if err != nil {
				fmt.Println("Error converting level:", err)
				return
			}
			levels[i] = level
		}

		if checkReportSafety(levels) {
			safeReportsCount++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("Number of safe reports:", safeReportsCount)
}
