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
	fmt.Printf("Found %d segments\n", len(segments))
	fmt.Println("segments", segments)

	sum := 0
	mulPattern := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	firstSegment := strings.Split(segments[0], "don't()")[0]
	fmt.Printf("First segment before don't(): %s\n", firstSegment)

	matches := mulPattern.FindAllStringSubmatch(firstSegment, -1)
	fmt.Printf("Found %d matches in first segment\n", len(matches))

	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		fmt.Printf("Multiplying %d * %d = %d\n", num1, num2, num1*num2)
		sum += num1 * num2
	}

	for i, segment := range segments[1:] {
		fmt.Printf("\nProcessing segment %d\n", i+1)
		enabledPart := strings.Split(segment, "don't()")[0]
		fmt.Printf("Segment content before don't(): %s\n", enabledPart)

		matches := mulPattern.FindAllStringSubmatch(enabledPart, -1)
		fmt.Printf("Found %d matches in this segment\n", len(matches))

		for _, match := range matches {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			fmt.Printf("Multiplying %d * %d = %d\n", num1, num2, num1*num2)
			sum += num1 * num2
		}
	}

	fmt.Println("Sum:", sum)
}
