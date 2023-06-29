package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "Главрыба"
	fmt.Println(Reverse(s1))
	fmt.Println(Reverse2(s1))
}

func Reverse(s string) string {
	runes := []rune(s)
	// Идем по строке с двух концов и меняем местами символы
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Reverse2(s string) string {
	builder := strings.Builder{}
	runes := []rune(s)
	for i := len(runes) - 1; i >= 0; i-- {
		builder.WriteRune(runes[i])
	}
	return builder.String()
}
