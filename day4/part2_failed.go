package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day4/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var cards [][]int
	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		count++
		parts := strings.Split(line, "|")
		winNumbers := parseNumbers(parts[0])
		yourNumbers := parseNumbers(parts[1])
		cards = append(cards, append(winNumbers, yourNumbers...))
	}

	totalCards := processCards(cards)
	fmt.Println("Total scratchcards:", totalCards+count)
}

func parseNumbers(s string) []int {
	var numbers []int
	for _, str := range strings.Fields(s) {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Error parsing number:", err)
			continue
		}
		numbers = append(numbers, num)
	}
	return numbers
}

func processCards(cards [][]int) int {
	cardCounts := make([]int, len(cards))
	for i := range cardCounts {
		cardCounts[i] = 1 // initialize with 1 for each original card
	}

	for i := 0; i < len(cards); i++ {
		winNumbers := cards[i][:len(cards[i])/2]
		yourNumbers := cards[i][len(cards[i])/2:]
		matchCount := countMatches(winNumbers, yourNumbers)

		for j := 1; j <= matchCount && i+j < len(cards); j++ {
			cardCounts[i+j] += cardCounts[i] // add copies based on the current count of this card
		}
	}

	total := 0
	for _, count := range cardCounts {
		total += count
	}
	return total
}

func countMatches(winNumbers, yourNumbers []int) int {
	matchCount := 0
	for _, num := range winNumbers {
		if contains(yourNumbers, num) {
			matchCount++
		}
	}
	return matchCount
}

func contains(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
