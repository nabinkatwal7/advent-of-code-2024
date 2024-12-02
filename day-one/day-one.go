package dayone

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func read_file(filepath string) ([]int, []int, error) {
	data, err := os.Open(filepath)

	if err != nil {
		return nil, nil, err
	}
	defer data.Close()

	var column1 []int
	var column2 []int

	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		if len(line) != 2 {
			fmt.Println("Skipping invalid line:", scanner.Text())
			continue
		}

		num1, err1 := strconv.Atoi(line[0])
		num2, err2 := strconv.Atoi(line[1])

		if err1 != nil || err2 != nil {
			fmt.Println("Error converting line to integers:", line)
			continue
		}

		column1 = append(column1, num1)
		column2 = append(column2, num2)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return column1, column2, nil
}

func sumArray(arr []float64) float64 {
	sum := 0.0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func Part1() {
	filepath := "./day-one/input.txt"

	var distance []float64

	column1, column2, err := read_file(filepath)

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	sort.Ints(column1)
	sort.Ints(column2)

	for i := 0; i < len(column1); i++ {
		distance = append(distance, math.Abs(float64(column1[i]-column2[i])))
	}

	totalDistance := sumArray(distance)

	fmt.Printf("total distance: %.0f\n ", totalDistance)
}

func calculateSimilarityScore(column1, column2 []int) int {
	countMap := make(map[int]int)

	for _, num := range column2 {
		countMap[num]++
	}

	var totalScore int
	for _, num := range column1 {
		totalScore += num * countMap[num]
	}

	return totalScore
}

func Part2() {

	filepath := "./day-one/input.txt"

	column1, column2, err := read_file(filepath)

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	totalScore := calculateSimilarityScore(column1, column2)

	fmt.Println("total similarity score:", totalScore)
}
