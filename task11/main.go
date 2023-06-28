package main

import "fmt"

func main() {
	set1 := map[string]bool{
		"c++":    true,
		"python": true,
		"golang": true,
		"sql":    true,
	}
	set2 := map[string]bool{
		"sql":    true,
		"c#":     true,
		"java":   true,
		"golang": true,
	}
	intersection := make(map[string]bool)
	for key := range set1 {
		if set2[key] {
			intersection[key] = true
		}
	}
	fmt.Println(intersection)
}
