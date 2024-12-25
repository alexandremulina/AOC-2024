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

	if len(antennaList) >= 2 {
		for _, ant := range antennaList {
			key := fmt.Sprintf("%d,%d", ant[0], ant[1])
			antinodes[key] = true
		}
	}

	// Compare pairs of antennas
	for i := 0; i < len(antennaList); i++ {
		for j := i + 1; j < len(antennaList); j++ {
			ant1 := antennaList[i]
			ant2 := antennaList[j]

			dx := ant2[1] - ant1[1]
			dy := ant2[0] - ant1[0]

			gcd := gcd(abs(dx), abs(dy))
			if gcd > 0 {
				dx /= gcd
				dy /= gcd
			}

			// Check all points along the line in both directions
			curr := []int{ant1[0], ant1[1]}

			for isInBounds(float64(curr[0]), float64(curr[1]), maxRows, maxCols) {
				key := fmt.Sprintf("%d,%d", curr[0], curr[1])
				antinodes[key] = true
				curr[0] += dy
				curr[1] += dx
			}

			curr = []int{ant1[0] - dy, ant1[1] - dx}
			for isInBounds(float64(curr[0]), float64(curr[1]), maxRows, maxCols) {
				key := fmt.Sprintf("%d,%d", curr[0], curr[1])
				antinodes[key] = true
				curr[0] -= dy
				curr[1] -= dx
			}
		}
	}
	return antinodes
}

// Helper function to find GCD
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return abs(a)
}

// Helper function for absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isInteger(x float64) bool {
	return math.Floor(x) == x
}

func isInBounds(row, col float64, maxRows, maxCols int) bool {
	return row >= 0 && row < float64(maxRows) && col >= 0 && col < float64(maxCols)
}
