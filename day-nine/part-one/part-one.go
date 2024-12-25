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
	lines := strings.Split(string(content), "\n")
	fmt.Println("lines", lines)
	code := lines[0]
	fmt.Println("leng", len(code))
	var arrEachMemo [][]int
	for i := 0; i < len(code); i += 2 {
		fileLeng, _ := strconv.Atoi(string(code[i]))
		fileID := i / 2
		fileSpace := 0
		if i+1 < len(code) {
			fileSpace, _ = strconv.Atoi(string(code[i+1]))
		}
		arrEachMemo = append(arrEachMemo, []int{fileID, fileLeng, fileSpace})
		// fmt.Printf("\n File ID %d File Length %d, file space:%d", fileID, fileLeng, fileSpace)
	}
	fmt.Println("removeSpaces", arrEachMemo)
	removeSpaces(arrEachMemo)

}

func removeSpaces(arrMemo [][]int) {
	totalSize := 0
	for _, file := range arrMemo {
		totalSize += file[1]
	}
	fmt.Println("Total size:", totalSize)

	result := make([]int, totalSize)

	currentPos := totalSize - 1
	for i := len(arrMemo) - 1; i >= 0; i-- {
		fileID := arrMemo[i][0]
		fileLength := arrMemo[i][1]

		for j := 0; j < fileLength; j++ {
			result[currentPos] = fileID
			currentPos--
		}
	}

	fmt.Println("Final:", result)
}
