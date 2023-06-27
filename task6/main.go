package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	stopCodeFlag := flag.Int("c", 1, "Number of way to stop goroutine")
	flag.Parse()
	stopCode := *stopCodeFlag
	wg := sync.WaitGroup{}
	switch stopCode {
	case 1:
		// Остановка с помощью контекста
		// Можно также использовать WithCancel или WithDeadline
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		wg.Add(1)
		go func(ctx context.Context) {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("Stopping goroutine with context")
					wg.Done()
					return
				default:
					time.Sleep(600 * time.Millisecond)
					fmt.Println("Waiting...")
				}
			}
		}(ctx)
	case 2:
		// С помощью канала, сообщающего о завершении
		stopChan := make(chan interface{})
		wg.Add(1)
		go func(stopChan chan interface{}) {
			for {
				select {
				case <-stopChan:
					fmt.Println("Stopping goroutine with stop channel")
					wg.Done()
					return
				default:
					time.Sleep(600 * time.Millisecond)
					fmt.Println("Waiting...")
				}
			}
		}(stopChan)
		time.Sleep(2 * time.Second)
		stopChan <- struct{}{}
	case 3:
		go func() {
			for {
				time.Sleep(600 * time.Millisecond)
				fmt.Println("Waiting...")
			}
		}()
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt)
		signal.Notify(sigChan, os.Kill)
		sig := <-sigChan
		fmt.Println("Got signal: ", sig)
		fmt.Println("Stopping main goroutine")
	default:
		fmt.Println("unkown flag value")
	}
	wg.Wait()
}
