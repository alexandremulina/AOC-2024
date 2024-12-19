package main

import (
	"fmt"
	"os"
	"strings"
)

type Position struct {
	row, col int
}

type Direction struct {
	deltaRow, deltaCol int
}

func main() {
	content, err := os.ReadFile("../day6-input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	lines := strings.Split(string(content), "\n")

	var startRow, startCol int
	for row, line := range lines {
		if col := strings.Index(line, "^"); col != -1 {
			startRow = row
			startCol = col
			break
		}
	}

	directions := []Direction{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	currentDir := 0 // Start facing up
	visited := make(map[Position]bool)

	current := Position{startRow, startCol}
	visited[current] = true

	for {
		nextRow := current.row + directions[currentDir].deltaRow
		nextCol := current.col + directions[currentDir].deltaCol

		if nextRow < 0 || nextRow >= len(lines) ||
			nextCol < 0 || nextCol >= len(lines[0]) {
			break
		}

		// Print current position and direction
		fmt.Printf("Position: (%d,%d), Direction: %d\n", current.row, current.col, currentDir)

		if lines[nextRow][nextCol] == '#' {
			currentDir = (currentDir + 1) % 4
		} else {
			current = Position{nextRow, nextCol}
			visited[current] = true
		}
	}

	for row := 0; row < len(lines); row++ {
		for col := 0; col < len(lines[0]); col++ {
			pos := Position{row, col}
			if visited[pos] {
				fmt.Print("X")
			} else if lines[row][col] == '#' {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	fmt.Printf("Number of positions visited: %d\n", len(visited))
}
