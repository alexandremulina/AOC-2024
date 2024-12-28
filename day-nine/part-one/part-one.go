package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("../day9-input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Remove any whitespace or newlines
	input := strings.TrimSpace(string(content))

	// Create disk map from input
	diskMap := createDiskMap(input)
	fmt.Println("Initial disk map:", diskMap)

	moves := moveFiles(diskMap)
	checksum := calculateChecksum(moves[len(moves)-1])
	fmt.Printf("\nFinal Checksum: %d\n", checksum)
}

func createDiskMap(input string) string {
	var result strings.Builder
	fileID := 0

	// To make the input work
	digits := strings.Split(input, "")

	for i := 0; i < len(digits); i++ {
		length, _ := strconv.Atoi(digits[i])

		if i%2 == 0 {
			for j := 0; j < length; j++ {
				// Add 'length' number of fileID characters
				result.WriteRune(rune('0' + fileID))
			}
			fileID++
		} else {

			for j := 0; j < length; j++ {
				result.WriteRune('.')
			}
		}
	}

	return result.String()
}

func moveFiles(diskMap string) []string {
	result := []string{diskMap}
	currentMap := diskMap

	for {
		moved := false
		chars := []rune(currentMap)

		// Find the last item
		for i := len(chars) - 1; i > 0; i-- {
			if chars[i] != '.' {

				for j := 0; j < i; j++ {
					if chars[j] == '.' {
						chars[j] = chars[i]
						chars[i] = '.'
						moved = true
						break
					}
				}
				if moved {
					break
				}
			}
		}

		if !moved {
			break
		}

		newMap := string(chars)
		result = append(result, newMap)
		currentMap = newMap
	}

	fmt.Println("\nInitial state:")
	fmt.Println(result[0])
	fmt.Println("\nFinal state:")
	fmt.Println(result[len(result)-1])

	return []string{result[0], result[len(result)-1]}
}

func calculateChecksum(finalMap string) int64 {
	var checksum int64 = 0

	fmt.Println("\n=== Calculating Checksum ===")
	fmt.Println("Final Map:", finalMap)

	chars := []rune(finalMap)

	for pos, char := range chars {
		if char == '.' {
			continue
		}

		fileID := int64(char - '0')
		value := int64(pos) * fileID

		fmt.Printf("Position %d * File ID %d = %d\n", pos, fileID, value)
		checksum += value
	}

	fmt.Printf("\nFinal Checksum: %d\n", checksum)
	return checksum
}
