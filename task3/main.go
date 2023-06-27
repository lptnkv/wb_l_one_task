package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	arr := [...]int32{2, 4, 6, 8, 10}
	var sum int32
	wg := sync.WaitGroup{}
	for _, val := range arr {
		wg.Add(1)
		go func(val int32) {
			// Атомарное прибавление, чтобы результат был точным
			atomic.AddInt32(&sum, val)
			wg.Done()
		}(val)
	}
	wg.Wait()
	fmt.Println(sum)
}
