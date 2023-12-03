package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// Replaces number words in a string with their corresponding digits
func replaceNumberWords(s string, numWordMap map[string]int) string {
	result := ""
	i := 0
	for i < len(s) {
		found := false
		for word, num := range numWordMap {
			if strings.HasPrefix(s[i:], word) {
				result += strconv.Itoa(num)
				i += len(word)
				found = true
				break
			}
		}
		if !found {
			result += string(s[i])
			i++
		}
	}
	return result
}

// Finds the first and last digit in a string
func findFirstAndLastDigit(s string) (int, int) {
	firstDigit, lastDigit := -1, -1
	for _, r := range s {
		if unicode.IsDigit(r) {
			if firstDigit == -1 {
				firstDigit = int(r - '0')
			}
			lastDigit = int(r - '0')
		}
	}
	return firstDigit, lastDigit
}

// Processes the file and calculates the total sum
func processFile(filepath string) int {
	numWordMap := map[string]int{
		"one": 1, "two": 2, "three": 3, "four": 4,
		"five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
	}

	totalSum := 0
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return -1
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		line = replaceNumberWords(line, numWordMap)
		firstDigit, lastDigit := findFirstAndLastDigit(line)
		if firstDigit != -1 && lastDigit != -1 {
			totalSum += firstDigit*10 + lastDigit
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return -1
	}

	return totalSum
}

func main() {
	filepath := "day1/input.txt" // Update with the actual file path
	sum := processFile(filepath)
	fmt.Println("Sum of calibration values:", sum)
}
