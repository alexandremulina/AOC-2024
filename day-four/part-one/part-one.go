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
	fmt.Println(lines)

	counter := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			counter += calculeAndCheckDir(lines, i, j, 0, 1)   // right
			counter += calculeAndCheckDir(lines, i, j, 1, 0)   // down
			counter += calculeAndCheckDir(lines, i, j, 1, 1)   // diagonal right-down
			counter += calculeAndCheckDir(lines, i, j, -1, 1)  // diagonal right-up
			counter += calculeAndCheckDir(lines, i, j, 0, -1)  // left
			counter += calculeAndCheckDir(lines, i, j, -1, 0)  // up
			counter += calculeAndCheckDir(lines, i, j, -1, -1) // diagonal left-up
			counter += calculeAndCheckDir(lines, i, j, 1, -1)  // diagonal left-down
		}
	}
	fmt.Printf("The count is %d.\n", counter)
}

func calculeAndCheckDir(lines []string, row, col, dirRow, dirCol int) int {
	word := "XMAS"
	if row < 0 || col < 0 || row >= len(lines) || col >= len(lines[0]) {
		return 0
	}

	if !isInBounds(lines, row, col, dirRow, dirCol, len(word)) {
		return 0
	}

	for i := 0; i < len(word); i++ {
		if lines[row+dirRow*i][col+dirCol*i] != word[i] {
			return 0
		}
	}
	return 1

}

func isInBounds(grid []string, row, col, dRow, dCol, length int) bool {
	lastRow := row + dRow*(length-1)
	lastCol := col + dCol*(length-1)
	return lastRow >= 0 && lastRow < len(grid) && lastCol >= 0 && lastCol < len(grid[0])
}
