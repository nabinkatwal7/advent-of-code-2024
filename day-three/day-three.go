package daythree

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Part1() {
	file, err := os.Open("./day-three/input.txt")
	

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	functionRegex := `mul\((-?\d+),\s*(-?\d+)\)`
	

	re := regexp.MustCompile(functionRegex)

	scanner := bufio.NewScanner(file)
	totalSum := 0

	for scanner.Scan() {
		text := scanner.Text()

		matches := re.FindAllStringSubmatch(text, -1)
		
		for _, match := range matches {
			// Convert the captured strings to integers
			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])
	
			// Multiply and accumulate the result
			totalSum += x * y
		}
	}
	fmt.Println("Total sum:", totalSum)
}