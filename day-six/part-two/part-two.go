package main

import (
	"fmt"
	"os"
	"strings"
)

type Position struct {
	row, col int
}

type State struct {
	pos Position
	dir int
}

type Direction struct {
	deltaRow, deltaCol int
}

func simulateGuard(lines []string, startPos Position, obstaclePos Position) bool {
	directions := []Direction{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	visited := make(map[State]bool)
	current := startPos
	currentDir := 0

	tempMap := make([]string, len(lines))
	copy(tempMap, lines)
	if obstaclePos.row >= 0 {
		row := []rune(tempMap[obstaclePos.row])
		row[obstaclePos.col] = '#'
		tempMap[obstaclePos.row] = string(row)
	}

	for {
		state := State{current, currentDir}
		if visited[state] {
			return true
		}
		visited[state] = true

		nextRow := current.row + directions[currentDir].deltaRow
		nextCol := current.col + directions[currentDir].deltaCol

		if nextRow < 0 || nextRow >= len(tempMap) ||
			nextCol < 0 || nextCol >= len(tempMap[0]) {
			return false
		}

		if tempMap[nextRow][nextCol] == '#' {
			currentDir = (currentDir + 1) % 4
		} else {
			current = Position{nextRow, nextCol}
		}
	}
}

func main() {
	content, err := os.ReadFile("../day6-input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	lines := strings.Split(string(content), "\n")

	var startPos Position
	for row, line := range lines {
		if col := strings.Index(line, "^"); col != -1 {
			startPos = Position{row, col}
			break
		}
	}

	loopCount := 0
	for row := 0; row < len(lines); row++ {
		for col := 0; col < len(lines[0]); col++ {
			if lines[row][col] == '.' && (row != startPos.row || col != startPos.col) {
				if simulateGuard(lines, startPos, Position{row, col}) {
					loopCount++
				}
			}
		}
	}

	fmt.Printf("Number of possible obstruction positions: %d\n", loopCount)
}
