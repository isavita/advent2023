package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day3/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	partSum := 0
	gearRatioSum := 1

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			// Check for part numbers
			if isPartNumber(lines, i, j) {
				number, _ := strconv.Atoi(string(lines[i][j]))
				partSum += number
			}

			// Check for gears
			if lines[i][j] == '*' {
				adjacentNumbers := getAdjacentNumbers(lines, i, j)
				if len(adjacentNumbers) == 2 {
					gearRatioSum *= adjacentNumbers[0] * adjacentNumbers[1]
				}
			}
		}
	}

	fmt.Println("Sum of part numbers:", partSum)
	fmt.Println("Sum of gear ratios:", gearRatioSum)
}

func isPartNumber(lines []string, i, j int) bool {
	return isSymbol(lines[i][j-1]) || isSymbol(lines[i][j+1]) ||
		isSymbol(lines[i-1][j-1]) || isSymbol(lines[i-1][j]) ||
		isSymbol(lines[i-1][j+1]) || isSymbol(lines[i+1][j-1]) ||
		isSymbol(lines[i+1][j]) || isSymbol(lines[i+1][j+1])
}

func isSymbol(ch byte) bool {
	return ch != '.' && (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func getAdjacentNumbers(lines []string, i, j int) []int {
	var numbers []int
	if i > 0 && isNumber(lines[i-1][j]) {
		numbers = append(numbers, toInt(lines[i-1][j]))
	}
	if j > 0 && isNumber(lines[i][j-1]) {
		numbers = append(numbers, toInt(lines[i][j-1]))
	}
	if i < len(lines)-1 && isNumber(lines[i+1][j]) {
		numbers = append(numbers, toInt(lines[i+1][j]))
	}
	if j < len(lines[i])-1 && isNumber(lines[i][j+1]) {
		numbers = append(numbers, toInt(lines[i][j+1]))
	}
	return numbers
}

func isNumber(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func toInt(ch byte) int {
	num, _ := strconv.Atoi(string(ch))
	return num
}
