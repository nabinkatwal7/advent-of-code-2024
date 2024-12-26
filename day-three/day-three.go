package daythree

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Part1() {
	// Open the input file
	file, err := os.Open("./day-three/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Regex for mul(x, y), do(), and don't()
	functionRegex := `mul\((-?\d+),\s*(-?\d+)\)`
	doPattern := `do\(\)`
	dontPattern := `don\'t\(\)`

	// Compile the regular expressions
	re := regexp.MustCompile(functionRegex)
	doRe := regexp.MustCompile(doPattern)
	dontRe := regexp.MustCompile(dontPattern)

	// Start with mul() instructions enabled by default
	enabled := true
	totalSum := 0

	// Scan through each line in the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		// Check for do() and don't() instructions to update enabled state
		if doRe.MatchString(text) {
			enabled = true // Enable mul() instructions
		} else if dontRe.MatchString(text) {
			enabled = false // Disable mul() instructions
		}

		// Find all mul(x, y) matches in the line
		matches := re.FindAllStringSubmatch(text, -1)

		// Process each match (mul(x, y)) only if enabled
		for _, match := range matches {
			// Only process multiplication if enabled
			if enabled {
				// Convert the captured numbers to integers
				x, errX := strconv.Atoi(match[1])
				y, errY := strconv.Atoi(match[2])

				// Ensure no errors in conversion
				if errX != nil || errY != nil {
					fmt.Println("Error converting numbers:", match[1], match[2])
					continue
				}

				// Multiply and accumulate the result
				result := x * y
				totalSum += result
				fmt.Println("Multiplying", x, "by", y, "Result:", result)
			} else {
				fmt.Println("Skipping mul() due to disabled state")
			}
		}
	}

	// Check for any scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	// Output the total sum of all enabled multiplications
	fmt.Println("Total sum:", totalSum)
}