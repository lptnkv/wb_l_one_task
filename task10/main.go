package main

import "fmt"

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	temps := make(map[int][]float64)
	for _, val := range arr {
		key := floor(val)
		_, ok := temps[key]
		if ok {
			temps[key] = append(temps[key], val)
		} else {
			temps[key] = []float64{val}
		}
	}
	fmt.Println(temps)

}

// Окруляет до ближайшего десятка вниз по модулю
func floor(a float64) int {
	return int(a) / 10 * 10
}
