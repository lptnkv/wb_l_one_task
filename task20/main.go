package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "snow dog sun"
	fmt.Println(reverseWords(s1))
}

// С использованием strings.Split и strings.Join
func reverseWords(s string) string {
	words := strings.Split(s, " ")
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}
