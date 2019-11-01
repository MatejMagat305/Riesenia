package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	mp := make(map[string]int)
	for _, fld := range strings.Fields(s){
		mp[fld]++
	}
	return mp
}

func main() {
	wc.Test(WordCount)
}
