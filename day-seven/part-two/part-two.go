package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("../day7-input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	lines := strings.Split(string(content), "\n")
	re := regexp.MustCompile(`(\d+):(.+)`)
	var sum int64 = 0

	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		if match == nil {
			continue
		}

		target, _ := strconv.ParseInt(match[1], 10, 64)
		values := strings.Fields(match[2])
		nums := make([]int64, len(values))
		for i, v := range values {
			nums[i], _ = strconv.ParseInt(v, 10, 64)
		}

		length := len(nums) - 1
		arrCombi := generateCombination([]string{"*", "+", "||"}, length)

		for _, combi := range arrCombi {
			if makeCalc(combi, target, nums) {
				fmt.Printf("Found valid equation: %d = %v with ops %v\n", target, nums, combi)
				sum += target
				break
			}
		}
	}

	fmt.Println("Final Target Sum:", sum)
}

func generateCombination(operators []string, length int) [][]string {
	if length == 0 {
		return [][]string{{}}
	}
	var result [][]string
	subCombinations := generateCombination(operators, length-1)

	for _, op := range operators {
		for _, sub := range subCombinations {
			result = append(result, append([]string{op}, sub...))
		}
	}
	return result
}

func makeCalc(arrComb []string, target int64, arrNum []int64) bool {
	result := arrNum[0]
	for i := 0; i < len(arrComb); i++ {
		result = operations[arrComb[i]](result, arrNum[i+1])
	}
	return result == target
}

var operations = map[string]func(int64, int64) int64{
	"*": func(a, b int64) int64 { return a * b },
	"+": func(a, b int64) int64 { return a + b },
	"||": func(a, b int64) int64 {
		bStr := strconv.FormatInt(b, 10)
		return a*int64(math.Pow10(len(bStr))) + b
	},
}
