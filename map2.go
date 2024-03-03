package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Hello, world! Welcome to the world of Golang. Golang is great."

	// Convert the text to lowercase and split it into words.
	words := strings.Fields(strings.ToLower(text))

	// Create a map to hold the word frequencies.
	frequencies := make(map[string]int)

	// Iterate over the words and increment their count in the map.
	for _, word := range words {
		frequencies[word]++
	}

	// Print the word frequencies.
	for word, freq := range frequencies {
		fmt.Printf("%s: %d\n", word, freq)
	}
}
