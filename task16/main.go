package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := randSlice(30)
	fmt.Println("Initial slice:", arr)
	arr = quickSort(arr)
	fmt.Println("Sorted slice:", arr)
}

// Функция разбиения: элементы меньше опорного помещаются перед ним, больше - после него
func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high] //Опорный элемент
	i := low           // Позиция, в которой окажется опорный элемент после сортировки данной части массива
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i] // Ставим опорный элемент в свою позицию
	return arr, i
}

func quickSortHelper(arr []int, low, high int) []int {
	if low < high {
		var p int // индекс опорного элемента после прошлой сортировки
		arr, p = partition(arr, low, high)
		// Сортируем части массива слева и справа опорного элемента
		arr = quickSortHelper(arr, low, p-1)
		arr = quickSortHelper(arr, p+1, high)
	}
	return arr
}

// Функция для удобного запуска сортировки
func quickSort(arr []int) []int {
	return quickSortHelper(arr, 0, len(arr)-1)
}

// Создание среза случайных чисел
func randSlice(size int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}
