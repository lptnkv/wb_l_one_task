package main

import (
	"fmt"
	"sync"
)

func main() {
	in := make(chan int)
	out := make(chan int)
	arr := []int{1, 5, 18, 20, 2, -3, 14, 5, 8, 10, 17, 121, 27}
	wg := sync.WaitGroup{}
	// Пишем в канал in
	go func() {
		for _, val := range arr {
			in <- val
		}
		close(in)
	}()

	// Читаем из in, умножаем на 2 и пишем в out
	go func() {
		for val := range in {
			out <- val * 2
		}
		close(out)
	}()

	// Читаем из out, выводим в stdout
	wg.Add(1)
	go func() {
		for val := range out {
			fmt.Println(val)
		}
		wg.Done()
	}()
	wg.Wait()
}
