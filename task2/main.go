package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := [...]int{2, 4, 6, 8, 10}
	// Переменная типа WaitGroup для ожидания выполнения горутин
	wg := sync.WaitGroup{}
	for _, val := range arr {
		// Увеличиваем счетчик wg
		wg.Add(1)
		// В параметры передаем val, так как иначе использовалось бы последнее значение из массива (захват переменной)
		go func(val int) {
			fmt.Println(val * val)
			wg.Done() // Горутина выполнилась, уменьшаем счетчик wg
		}(val)
	}
	// Ждем завершения выполнения всех горутин
	wg.Wait()
}
