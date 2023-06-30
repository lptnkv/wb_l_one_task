package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr = deleteWithoutOrder(arr, -1)
	fmt.Println(arr)
	arr = deleteWithOrder(arr, 5)
	fmt.Println(arr)
	arr = deleteWithoutOrder(arr, 4)
	fmt.Println(arr)
	arr = deleteWithOrder(arr, 100)
	fmt.Println(arr)
}

// Если порядок не важен
// Работает быстрее
func deleteWithoutOrder(slice []int, i int) []int {
	// Проверка, что индекс находится в границах среза
	if i < 0 || i >= len(slice) {
		return slice
	}
	// На место удаляемого элемента ставим последний элемент среза и возвращаем срез до последнего элемента исходного среза
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

// Если порядок важен
func deleteWithOrder(slice []int, i int) []int {
	// Проверка на нахождение индекса в границах среза
	if i < 0 || i >= len(slice) {
		return slice
	}
	// Элементы сдвигаются влево на один после i-го
	return append(slice[:i], slice[i+1:]...)
}
