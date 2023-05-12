package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	var m = make(map[string]int)
	ss := strings.Fields(s)
	for _, value := range ss {
		m[value]++
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
