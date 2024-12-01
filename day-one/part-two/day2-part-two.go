package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("../day1-input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	lines := strings.Split(string(content), "\n")
	a := []int{}
	b := []int{}
	for _, line := range lines {
		numbers := strings.Fields(line)
		if len(numbers) > 1 {
			num1, _ := strconv.Atoi(numbers[0])
			num2, _ := strconv.Atoi(numbers[1])
			a = append(a, num1)
			b = append(b, num2)
		}
	}
	sum := 0
	for i := 0; i < len(a); i++ {
		appearsTime := 0
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				appearsTime++
			}
		}
		result := a[i] * appearsTime
		sum += result
	}
	fmt.Println(sum)

}
