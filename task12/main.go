package main

import "fmt"

// Пустая структура для оптимизации использования памяти
type void struct{}

var member void

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]void)
	for _, val := range words {
		set[val] = member
	}
	for key := range set {
		fmt.Println(key)
	}
}
