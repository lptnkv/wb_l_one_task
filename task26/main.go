package main

import "fmt"

func main() {
	s1 := "abcd"
	s2 := "abCdefAaf"
	s3 := "aabcd"
	s4 := "😀привет😀"
	s5 := "привет😀"
	fmt.Println("s1: ", IsUnique(s1))
	fmt.Println("s2: ", IsUnique(s2))
	fmt.Println("s3: ", IsUnique(s3))
	fmt.Println("s4: ", IsUnique(s4))
	fmt.Println("s5: ", IsUnique(s5))
}

func IsUnique(s string) bool {
	dict := make(map[rune]bool)
	for _, val := range s {
		if found := dict[val]; found {
			return false
		}
		dict[val] = true
	}
	return true
}
