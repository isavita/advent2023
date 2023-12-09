package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	// Read the file and convert it into a 2D slice of runes.
	matrix, err := readFileToMatrix("day3/input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Calculate the sum of all gear ratios.
	sum := sumOfGearRatios(matrix)
	fmt.Println("Sum of gear ratios:", sum)
}

// readFileToMatrix reads the file and converts it into a 2D slice of runes.
func readFileToMatrix(filePath string) ([][]rune, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var matrix [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matrix = append(matrix, []rune(scanner.Text()))
	}
	return matrix, scanner.Err()
}

// sumOfGearRatios calculates the sum of all gear ratios in the matrix.
func sumOfGearRatios(matrix [][]rune) int {
	sum := 0
	for y, row := range matrix {
		for x, char := range row {
			if char == '*' {
				if ratios, valid := calculateGearRatio(matrix, x, y); valid {
					sum += ratios
				}
			}
		}
	}
	return sum
}

// calculateGearRatio calculates the gear ratio of a '*' at the given position, if valid.
func calculateGearRatio(matrix [][]rune, x, y int) (int, bool) {
	var partNumbers []int
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			adjX, adjY := x+dx, y+dy
			if adjY >= 0 && adjY < len(matrix) && adjX >= 0 && adjX < len(matrix[adjY]) {
				number, length := extractNumber(matrix, adjX, adjY)
				if length > 0 {
					partNumbers = append(partNumbers, number)
				}
			}
		}
	}
	if len(partNumbers) == 2 {
		return partNumbers[0] * partNumbers[1], true
	}
	return 0, false
}

// extractNumber extracts the whole number starting from the given position and returns the number and its length.
func extractNumber(matrix [][]rune, x, y int) (int, int) {
	if !unicode.IsDigit(matrix[y][x]) {
		return 0, 0
	}
	numberStr := ""
	for x < len(matrix[y]) && unicode.IsDigit(matrix[y][x]) {
		numberStr += string(matrix[y][x])
		x++
	}
	number, _ := strconv.Atoi(numberStr)
	return number, len(numberStr)
}
