package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("../day4-input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	lines := strings.Split(string(content), "\n")

	count := 0
	// Check each possible center point of the X
	for i := 1; i < len(lines)-1; i++ {
		for j := 1; j < len(lines[i])-1; j++ {
			if checkXPattern(lines, i, j) {
				count++
			}
		}
	}

	fmt.Printf("Found X-MAS %d times\n", count)
}

func checkXPattern(grid []string, centerRow, centerCol int) bool {
	// Check if center is 'A'
	if grid[centerRow][centerCol] != 'A' {
		return false
	}

	// Check all four possible combinations:
	// 1. Top-left to bottom-right: MAS and top-right to bottom-left: MAS
	// 2. Top-left to bottom-right: SAM and top-right to bottom-left: MAS
	// 3. Top-left to bottom-right: MAS and top-right to bottom-left: SAM
	// 4. Top-left to bottom-right: SAM and top-right to bottom-left: SAM

	topLeft := string([]byte{
		grid[centerRow-1][centerCol-1],
		grid[centerRow][centerCol],
		grid[centerRow+1][centerCol+1],
	})

	topRight := string([]byte{
		grid[centerRow-1][centerCol+1],
		grid[centerRow][centerCol],
		grid[centerRow+1][centerCol-1],
	})

	return isValidArm(topLeft) && isValidArm(topRight)
}

func isValidArm(s string) bool {
	return s == "MAS" || s == "SAM"
}
