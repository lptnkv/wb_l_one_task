package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// С использованием мьютекса
type MutexCounter struct {
	val int
	mu  *sync.Mutex
}

// Функция инкремента значения
func (c *MutexCounter) Increment() {
	c.mu.Lock()
	c.val += 1
	c.mu.Unlock()
}

// Вывод на экран
func (c *MutexCounter) PrintValue() {
	fmt.Println(c.val)
}

// С использованием типа atomic.Int64
type AtomicCounter struct {
	val atomic.Int64
}

func (c *AtomicCounter) Increment() {
	c.val.Add(1)
}

func (c *AtomicCounter) PrintValue() {
	fmt.Println(c.val.Load())
}

func main() {
	counter := MutexCounter{val: 0, mu: &sync.Mutex{}}
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			counter.Increment()
			wg.Done()
		}()
	}
	wg.Wait()
	counter.PrintValue()

	counter2 := AtomicCounter{val: atomic.Int64{}}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			counter2.Increment()
			wg.Done()
		}()
	}
	wg.Wait()
	counter2.PrintValue()

}
