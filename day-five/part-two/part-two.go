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
	fmt.Println(lines[0])
	pairs := [][2]int{}
	arraysNums := [][]int{}
	for _, line := range lines {
		nums := strings.Split(line, "|")
		if len(nums) == 2 {
			num1, _ := strconv.Atoi(nums[0])
			num2, _ := strconv.Atoi(nums[1])
			pairs = append(pairs, [2]int{num1, num2})
		}
		nums2 := strings.Split(line, ",")
		array := []int{}
		for _, numStr := range nums2 {
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
	// fmt.Println("Pairs", pairs)
	// fmt.Println("arraysNums", arraysNums)
	valid, invalid := checkRule(pairs, arraysNums)
	fmt.Println("valid", valid, "invalid", invalid)
	sum := processInvalidSequences(pairs, invalid)
	fmt.Println("Sum of middle numbers from corrected sequences:", sum)
}

func checkRule(pairs [][2]int, arrNums [][]int) ([][]int, [][]int) {
	validSeqs := [][]int{}
	invalidSeqs := [][]int{}
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
			validSeqs = append(validSeqs, arrNum)
		} else {
			invalidSeqs = append(invalidSeqs, arrNum)
		}
	}
	return validSeqs, invalidSeqs
}

func processInvalidSequences(pairs [][2]int, sequences [][]int) int {
	sum := 0

	for _, seq := range sequences {
		sortedSeq := make([]int, len(seq))
		copy(sortedSeq, seq)
		//Bubble sort implementation
		for i := 0; i < len(sortedSeq)-1; i++ {
			for j := 0; j < len(sortedSeq)-i-1; j++ {
				if shouldSwap(pairs, sortedSeq[j], sortedSeq[j+1]) {
					sortedSeq[j], sortedSeq[j+1] = sortedSeq[j+1], sortedSeq[j]
				}
			}
		}

		middleIdx := len(sortedSeq) / 2
		sum += sortedSeq[middleIdx]

		fmt.Printf("Original: %v -> Sorted: %v (middle: %d)\n",
			seq, sortedSeq, sortedSeq[middleIdx])
	}

	return sum
}

func shouldSwap(pairs [][2]int, a, b int) bool {
	for _, pair := range pairs {
		if pair[0] == b && pair[1] == a {
			return true
		}

		if pair[0] == a && pair[1] == b {
			return false
		}
	}
	return false
}
