package main

import "fmt"

func main() {
	// Отсортированный массив
	arr := []int{2, 15, 27, 44, 100, 191, 250, 1000, 1565, 2023, 3456}
	fmt.Println(arr)
	search := 10[]
	pos := binarySearch(arr, search)
	if pos != -1 {
		fmt.Printf("%d found at index %d\n", search, pos)
	} else {
		fmt.Println("search not found")
	}
}

func binarySearch(a []int, search int) (pos int) {
	// Границы поиска
	l := 0
	r := len(a) - 1
	for l <= r {
		// Находим середину текущего подмассива, в котором производим поиск
		mid := (l + r) / 2
		// Если элемент посередине меньше искомого, то сдвигаем левую границу к середине
		if a[mid] < search {
			l = mid + 1
			// Если элемент посередине больше искомого, то сдвигаем правую границу к середине
		} else if a[mid] > search {
			r = mid - 1
			// Иначе нашли искомый элемент
		} else {
			return mid
		}
	}
	// Если границы сомкнулись, то элемент не найден
	return -1
}
