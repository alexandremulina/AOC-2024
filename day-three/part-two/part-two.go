package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Read input from file
	content, err := os.ReadFile("../day3-input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	input := string(content)

	segments := strings.Split(input, "do()")

	sum := 0
	mulPattern := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	firstSegment := strings.Split(segments[0], "don't()")[0]
	matches := mulPattern.FindAllStringSubmatch(firstSegment, -1)
	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		sum += num1 * num2
	}

	for _, segment := range segments[1:] {
		enabledPart := strings.Split(segment, "don't()")[0]
		matches := mulPattern.FindAllStringSubmatch(enabledPart, -1)
		for _, match := range matches {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			sum += num1 * num2
		}
	}

	fmt.Println("Sum:", sum)
}
