package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// Sample text
	text := `Go is an open source programming language that makes it easy to build 
    simple, reliable, and efficient software. Go is expressive, concise, clean, 
    and efficient. Its concurrency mechanisms make it easy to write programs that 
    get the most out of multicore and networked machines, while its novel type 
    system enables flexible and modular program construction. Go compiles quickly 
    to machine code yet has the convenience of garbage collection and the power of 
    run-time reflection. It's a fast, statically typed, compiled language that feels 
    like a dynamically typed, interpreted language.`

	// 1. Split the text into words
	words := strings.Fields(strings.ToLower(text))
	// fmt.Println(words)

	// 2. Create a map to store word counts
	wordCounts := make(map[string]int)

	// 3. Count occurrences of each word

	for _, word := range words {
		word = strings.Trim(word, ".,!?")
		wordCounts[word]++
	}
	fmt.Println(wordCounts)

	// 4. Print the word counts
	fmt.Println("Word frequencies:")
	for word, count := range wordCounts {
		fmt.Printf("%s: %d\n", word, count)
	}

	// 5. Find the most frequent word
	maxCount := 0
	var mostFrequentWord string

	for word, count := range wordCounts {
		if count > maxCount {
			maxCount = count
			mostFrequentWord = word
		}
	}
	fmt.Printf("\nMost frequent word: '%s' (count: %d)\n", mostFrequentWord, maxCount)

	// 6. Sort words by frequency
	var wfSlice []wordFreq
	for word, count := range wordCounts {
		wfSlice = append(wfSlice, wordFreq{word, count})
	}

	sort.Slice(wfSlice, func(i, j int) bool {
		return wfSlice[i].count > wfSlice[j].count
	})

	fmt.Println("\nWords sorted by frequency:", wfSlice)

	// 7. Print top 5 most frequent words
	fmt.Println("\nTop 5 most frequent words:")
	for i := 0; i < 5 && i < len(wfSlice); i++ {
		fmt.Printf("%s: %d\n", wfSlice[i].word, wfSlice[i].count)
	}

}

type wordFreq struct {
	word  string
	count int
}
