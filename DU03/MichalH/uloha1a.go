// https://tour.golang.org/moretypes/23

package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	wordcount := make(map[string]int)
	splitwords := strings.Fields(s)

	for _, word := range splitwords {
		elem, ok := wordcount[word]
		if ok {
			wordcount[word] = elem + 1
		} else {
			wordcount[word] = 1
		}
	}

	return wordcount
}

func main() {
	wc.Test(WordCount)
}

