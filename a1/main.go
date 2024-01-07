package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortByA(words []string) []string {
	// Custom sorting function
	sort.Slice(words, func(i, j int) bool {
		// Count 'a' characters in each word
		countA_i := strings.Count(words[i], "a")
		countA_j := strings.Count(words[j], "a")

		// Sort by the number of 'a' characters in decreasing order
		if countA_i != countA_j {
			return countA_i > countA_j
		}

		// If the number of 'a' characters is the same, sort by word length in decreasing order
		return len(words[i]) > len(words[j])
	})

	return words
}

func main() {
	// Example input
	wordList := []string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}

	// Sorting the words
	sortedResult := sortByA(wordList)

	// Printing the result
	fmt.Println(sortedResult)
}
