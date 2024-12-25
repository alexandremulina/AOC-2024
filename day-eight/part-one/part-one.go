package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("../day8-input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	lines := strings.Split(string(content), "\n")

	// Create a map to store antennas by frequency
	antennasByFreq := make(map[rune][][]int)

	for row, line := range lines {
		for col, char := range line {
			if char != '.' {

				antennasByFreq[char] = append(antennasByFreq[char], []int{row, col})
			}
		}
	}

	maxRows := len(lines)
	maxCols := len(lines[0])
	fmt.Println("antennasByFreq", antennasByFreq)
	allAntinodes := make(map[string]bool)
	for freq, antennaList := range antennasByFreq {
		if len(antennaList) >= 2 {
			fmt.Printf("\nProcessing frequency %c with %d antennas\n", freq, len(antennaList))
			antinodes := findAntinodes(antennaList, maxRows, maxCols)
			fmt.Printf("Found %d antinodes for frequency %c\n", len(antinodes), freq)
			for k := range antinodes {
				allAntinodes[k] = true
			}
		}
	}

	fmt.Printf("\nTotal unique antinode locations: %d\n", len(allAntinodes))
}

func findAntinodes(antennaList [][]int, maxRows, maxCols int) map[string]bool {
	antinodes := make(map[string]bool)

	// Compare pairs
	for i := 0; i < len(antennaList); i++ {
		for j := i + 1; j < len(antennaList); j++ {
			ant1 := antennaList[i]
			ant2 := antennaList[j]

			dx := float64(ant2[1] - ant1[1])
			dy := float64(ant2[0] - ant1[0])

			x1 := float64(ant2[1]) + dx
			y1 := float64(ant2[0]) + dy

			x2 := float64(ant1[1]) - dx
			y2 := float64(ant1[0]) - dy

			fmt.Printf("Antenna pair: (%d,%d) and (%d,%d)\n", ant1[0], ant1[1], ant2[0], ant2[1])
			fmt.Printf("Potential antinode 1: (%.2f,%.2f)\n", y1, x1)
			fmt.Printf("Potential antinode 2: (%.2f,%.2f)\n", y2, x2)

			if isInteger(x1) && isInteger(y1) && isInBounds(y1, x1, maxRows, maxCols) {
				key := fmt.Sprintf("%d,%d", int(y1), int(x1))
				antinodes[key] = true
				fmt.Printf("Added antinode 1: %s\n", key)
			}
			if isInteger(x2) && isInteger(y2) && isInBounds(y2, x2, maxRows, maxCols) {
				key := fmt.Sprintf("%d,%d", int(y2), int(x2))
				antinodes[key] = true
				fmt.Printf("Added antinode 2: %s\n", key)
			}
		}
	}
	return antinodes
}

func isInteger(x float64) bool {
	return math.Floor(x) == x
}

func isInBounds(row, col float64, maxRows, maxCols int) bool {
	return row >= 0 && row < float64(maxRows) && col >= 0 && col < float64(maxCols)
}
