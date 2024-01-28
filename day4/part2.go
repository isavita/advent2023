package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	// Read input from input.txt
	data, err := ioutil.ReadFile("day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Split the input into individual lines
	lines := strings.Split(string(data), "\n")

	// Create a map to store the cards and their winning numbers
	cards := make(map[int][]int)

	// Parse the input and populate the cards map
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " | ")
		cardNum, err := strconv.Atoi(strings.Split(parts[0], ":")[1])
		if err != nil {
			log.Fatal(err)
		}
		nums := strings.Fields(parts[1])
		var cardNums []int
		for _, num := range nums {
			n, err := strconv.Atoi(num)
			if err != nil {
				log.Fatal(err)
			}
			cardNums = append(cardNums, n)
		}
		cards[cardNum] = cardNums
	}

	// Initialize a map to store the count of each card
	cardCounts := make(map[int]int)

	// Initialize a map to store the copies of cards won
	copies := make(map[int]int)

	// Process the original and copied scratchcards until no more scratchcards are won
	for len(cards) > 0 {
		for cardNum, card := range cards {
			cardCount := cardCounts[cardNum]
			copyCount := copies[cardNum]

			// If there are no copies left, process the original card
			if copyCount == 0 {
				for _, num := range card {
					if _, ok := cardCounts[num]; ok {
						cardCounts[num] += 1 << cardCount
					}
				}
				copies[cardNum] = 1
			} else {
				// Process the copies of the card
				for i := 1; i <= copyCount; i++ {
					for _, num := range card {
						if _, ok := cardCounts[num]; ok {
							cardCounts[num] += 1 << (cardCount + i)
						}
					}
					copies[cardNum] = 0
				}
			}

			// Remove the processed card from the original cards map
			delete(cards, cardNum)
		}
	}

	// Calculate the total scratchcards won
	totalScratchcards := 0
	for _, count := range cardCounts {
		totalScratchcards += count
	}

	fmt.Println("Total scratchcards won:", totalScratchcards)
}
