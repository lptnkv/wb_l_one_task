package main

import (
	"fmt"
	"math/big"
	"math/rand"
)

func main() {
	// Создаем два больших числа значением от 2^61 до 2^62
	a := rand.Int63n(1<<62-1<<61) + 1<<61
	b := rand.Int63n(1<<62-1<<61) + 1<<61
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("a + b: ", a+b) // Возможно переполнение сетки
	fmt.Println("a - b: ", a-b)
	fmt.Println("a * b: ", a*b) // Возможно переполнение сетки
	fmt.Println("a / b: ", float64(a)/float64(b))

	// С использованием math/big
	bigA := big.NewInt(a)
	bigB := big.NewInt(b)
	fmt.Println("a + b: ", new(big.Int).Add(bigA, bigB))
	fmt.Println("a - b: ", new(big.Int).Sub(bigA, bigB))
	fmt.Println("a * b: ", new(big.Int).Mul(bigA, bigB))
	fmt.Println("a / b: ", big.NewRat(a, b).FloatString(10)) // float с точностью до 10 знаков
}
