package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	m := make(map[string]int)
	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}
	m["key"] = rand.Intn(100)
	fmt.Println("Initial value", m["key"])
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func(i int) {
			mutex.Lock()
			m["key"]++
			mutex.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(m)
}
