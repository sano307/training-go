package main

import (
	"strings"

	"code.google.com/p/go-tour/wc"
)

func main() {
	wc.Test(WordCount)
}

func WordCount(word string) map[string]int {
	strings := strings.Fields(word)
	countList := make(map[string]int)

	for _, s := range strings {
		countList[s]++
	}

	return countList
}
