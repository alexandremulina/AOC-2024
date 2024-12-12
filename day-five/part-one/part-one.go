package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("../day5-input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	lines := strings.Split(string(content), "\n")
	pairs := [][2]int{}
	arraysNums := [][]int{}
	for _, line := range lines {
		if strings.Contains(line, "|") {
			parts := strings.Split(line, "|")
			num1, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
			num2, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
			pairs = append(pairs, [2]int{num1, num2})
			continue
		}
		if line == "" {
			continue
		}
		nums := strings.Split(line, ",")
		array := []int{}
		for _, numStr := range nums {
			numStr = strings.TrimSpace(numStr)
			if numStr == "" {
				continue
			}
			num, err := strconv.Atoi(numStr)
			if err == nil {
				array = append(array, num)
			}
		}
		if len(array) > 0 {
			arraysNums = append(arraysNums, array)
		}
	}
	checkRule(pairs, arraysNums)
}

func checkRule(pairs [][2]int, arrNums [][]int) {
	results := [][]int{}
	for _, arrNum := range arrNums {
		isValid := true
		fmt.Println("Checking array:", arrNum)
		for _, pair := range pairs {
			first := pair[0]
			second := pair[1]
			idxFirst := slices.Index(arrNum, first)
			idxSecond := slices.Index(arrNum, second)

			if idxFirst != -1 && idxSecond != -1 {
				if idxFirst >= idxSecond {
					isValid = false
					break
				}
			}
		}

		if isValid {
			results = append(results, arrNum)
		}
	}

	sum := 0
	for _, result := range results {
		middleIdx := len(result) / 2
		sum += result[middleIdx]
	}

	fmt.Println("Valid sequences:", results)
	fmt.Println("Sum of middle numbers:", sum)
}
