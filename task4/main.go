package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	var nWorkersFlag = flag.Int("n", 1, "number of workers")
	flag.Parse()
	nWorkers := *nWorkersFlag
	fmt.Println("number of workers:", nWorkers)
	ctx, cancel := context.WithCancel(context.Background())
	numChan := make(chan int)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Closing channel")
				close(numChan)
				return
			default:
				time.Sleep(time.Second)
				numChan <- rand.Intn(20)
			}
		}
	}(ctx)
	wg := sync.WaitGroup{}
	for i := 1; i <= nWorkers; i++ {
		wg.Add(1)
		go func(i int, ch chan int) {
			for val := range ch {
				fmt.Printf("Worker №%d: got %d\n", i, val)
			}
			wg.Done()
		}(i, numChan)
	}
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	sig := <-sigChan
	fmt.Println("Got signal: ", sig)
	fmt.Println("Stopping writing to main channel")
	cancel()
	fmt.Println("Waiting for workers")
	wg.Wait()
	fmt.Println("Workers finished the job")
}
