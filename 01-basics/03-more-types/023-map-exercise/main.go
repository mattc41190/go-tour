// !!! MUST BE RUN IN TOUR OF GO !!!

package main

import (
	"fmt"
	"strings"
)

// WordCount returns a map
func WordCount(sentence string) map[string]int {
	result := make(map[string]int)
	words := strings.Fields(sentence)
	for _, word := range words {
		_, ok := result[word]
		if ok {
			result[word]++
		} else {
			result[word] = 1
		}
	}
	return result
}

func main() {
	sentence := "Where should we all go as we are everyones favorite people from Texas"
	result := WordCount(sentence)
	fmt.Println(result)
}
