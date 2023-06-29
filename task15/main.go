package main

import (
	"fmt"
	"math/rand"
	"unicode/utf8"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var runes = []rune("Helo世界привет😀😉")

// глобальная переменная, лучше определить в функции main или someFunc, если не используется в других функциях
// var justString string

func main() {
	someFunc()
}

// В этой функции в переменную justString срезаются 100 байт
// Взятие среза не копирует лежащий под ним массив, так что он остается, пока gc не уничтожит его (то есть пока на него есть ссылки)
func someFunc() {
	// Определил переменную justString в этой функции, чтобы память освободилась после ее завершения
	var justString string

	// Создаем строку размером 1024 байт
	v := createHugeByteString(1 << 10)
	// Берем слайс размером 100 байт
	justString = v[:100]
	fmt.Println("string:", justString)

	v2 := createHugeRuneString(1 << 10)
	// Берем слайс размером 100 байт!!!
	justString = v2[:100]
	// Так как руны могут занимать больше 1 байта в памяти, то наша строка может обрезаться (����)
	fmt.Println("rune string:", justString)
	fmt.Println("length in bytes and runes: ", len(justString), utf8.RuneCountInString(justString))

	// Берем слайс размером 100 рун
	justString = firstNRunes(v2, 100)
	fmt.Println("better rune string:", justString)
	fmt.Println("length in bytes and runes: ", len(justString), utf8.RuneCountInString(justString))

}

// Создаем строку размером 1024 байт из символов ASCII
// Строки неизменяемые, так что у нас в памяти будут храниться все 1024 байт
func createHugeByteString(size int) string {
	b := make([]byte, size)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// Создаем строку размером 1024 рун из символов utf8
func createHugeRuneString(size int) string {
	r := make([]rune, size)
	for i := range r {
		randIndex := rand.Intn(len(runes))
		r[i] = runes[randIndex]
	}
	return string(r)
}

// Функция для создания слайса из первых n рун
func firstNRunes(str string, n int) string {
	v := []rune(str)
	if n >= len(v) {
		return str
	}
	return string(v[:n])
}

func firstNBytes(src string, n int) string {
	res := make([]byte, len(src))
	copy(res, src)
	return string(res)
}
