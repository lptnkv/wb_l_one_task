package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	nSecondsFlag := flag.Int("n", 3, "Number of seconds to work")
	flag.Parse()
	nSeconds := *nSecondsFlag
	ch := make(chan int)
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(nSeconds)*time.Second)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				close(ch)
				fmt.Println("Closing channel")
				fmt.Println("Waiting for reading goroutine to finish")
				return
			default:
				time.Sleep(500 * time.Millisecond)
				ch <- rand.Intn(100)
			}
		}
	}(ctx)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(ch chan int) {
		for val := range ch {
			fmt.Println(val)
		}
		wg.Done()
	}(ch)
	wg.Wait()
	fmt.Println("Finished")
}
