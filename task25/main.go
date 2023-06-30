package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("time.After function")
	fmt.Println("Before sleeping")
	sleepUsingAfter(1)
	fmt.Println("After sleeping")

	fmt.Println("time.NewTimer function")
	fmt.Println("Before sleeping")
	sleeUsingTimer(1)
	fmt.Println("After sleeping")

	fmt.Println("for loop function")
	fmt.Println("Before sleeping")
	sleepUsingLoop(1)
	fmt.Println("After sleeping")
}

// Создаем канал функцикй After и ждем сигнала из него
func sleepUsingAfter(x int) {
	<-time.After(time.Second * time.Duration(x))
}

// Создаем таймер и ждем сигнала из его канала
// Похоже на внутренее устройство функции time.After
func sleeUsingTimer(x int) {
	timer := time.NewTimer(time.Second * time.Duration(x))
	<-timer.C
}

// Используем пустой цикл for
func sleepUsingLoop(x int) {
	// Время, через которое надо "проснуться"
	needTime := time.Now().Add(time.Second * time.Duration(x))
	// Запускаем пустой цикл, который идет, пока текущее время меньше целевого
	for time.Now().Before(needTime) {
	}
}
