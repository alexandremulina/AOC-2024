package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("../day2-input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	lines := strings.Split(string(content), "\n")
	sum := 0
	for i, line := range lines {
		numbers := strings.Fields(line)
		if safeCheckPartTwo(numbers) {
			fmt.Printf("Line %d: SAFE - %v\n", i+1, numbers)
		} else {
			fmt.Printf("Line %d: UNSAFE - %v\n", i+1, numbers)
		}
		if safeCheckPartTwo(numbers) {
			sum++
		}
	}
	fmt.Println("SOMA:", sum)
}

func safeCheck(arrNumbers []string) bool {
	inc := false
	dec := false
	lastNum, _ := strconv.Atoi(arrNumbers[0])

	for i := 1; i < len(arrNumbers); i++ {
		currentNum, _ := strconv.Atoi(arrNumbers[i])

		if inc && dec {
			return false
		}

		if currentNum == lastNum {
			return false
		}

		if currentNum > lastNum {

			if dec {
				return false
			}

			if currentNum-lastNum > 3 {
				return false
			}
			inc = true
		}

		if currentNum < lastNum {

			if inc {
				return false
			}

			if lastNum-currentNum > 3 {
				return false
			}
			dec = true
		}

		lastNum = currentNum
	}

	return inc || dec
}

// PART TWO
func safeCheckPartTwo(arrNumbers []string) bool {
	if safeCheck(arrNumbers) {
		return true
	}

	for i := 0; i < len(arrNumbers); i++ {
		newArr := append([]string{}, arrNumbers[:i]...)
		newArr = append(newArr, arrNumbers[i+1:]...)

		if safeCheck(newArr) {
			return true
		}

	}
	return false

}
