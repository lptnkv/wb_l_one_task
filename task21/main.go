package main

import "fmt"

func main() {
	var cardReader USB = CardReader{card: MemoryCard{}}
	cardReader.connectWithUsbCable()
}

// Интерфейс для подключения через USB
type USB interface {
	connectWithUsbCable()
}

// Карта памяти
type MemoryCard struct{}

// Метод для вставки карты памяти
func (card *MemoryCard) insert() {
	fmt.Println("Карта памяти вставлена")
}

// Метод для копирования данных на карту
func (card *MemoryCard) copyData() {
	fmt.Println("Данные скопированы")
}

// Кардридер, который можно подключить через USB
// Является адаптером для карты памяти
type CardReader struct {
	card MemoryCard
}

// Реализация метода интерфейса USB
func (reader CardReader) connectWithUsbCable() {
	reader.card.insert()
	reader.card.copyData()
}
