package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("../day3-input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	lines := strings.Split(string(content), "\n")
	totalSum := 0
	for _, line := range lines {
		lineSum := sanitizeMUL(line)
		totalSum += lineSum
	}
	fmt.Println("FINAL:", totalSum)
}

// Part One use
func sanitizeMUL(sentence string) int {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(sentence, -1)
	fmt.Println(matches)
	sum := 0
	for _, match := range matches {
		fmt.Println(match)
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		sum += num1 * num2

	}
	return sum
}
