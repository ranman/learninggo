package main

import (
	"fmt"
	"strings"
	"tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Split(s, ' ')
	wordCounts := make(map[string]int)
	for word := range words {
		wordCounts[word]++
	}
	return wordCounts
}

func main() {
	wc.Test(WordCount)
}
