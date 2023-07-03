package main

import (
	"fmt"
	"unicode"
)

func main() {
	s1 := "abcd"
	s2 := "abCdefAaf"
	s3 := "aabcd"
	s4 := "😀привет😀"
	s5 := "привет😀"
	s6 := "aAbB"
	fmt.Println("s1: ", IsUnique(s1))
	fmt.Println("s2: ", IsUnique(s2))
	fmt.Println("s3: ", IsUnique(s3))
	fmt.Println("s4: ", IsUnique(s4))
	fmt.Println("s5: ", IsUnique(s5))
	fmt.Println("s6: ", IsUnique(s6))
}

func IsUnique(s string) bool {
	dict := make(map[rune]bool)
	for _, val := range s {
		lower := unicode.ToLower(val)
		if found := dict[lower]; found {
			return false
		}
		dict[val] = true
	}
	return true
}
