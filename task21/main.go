package main

import "fmt"

func main() {
	var cardReader USB = CardReader{card: MemoryCard{}}
	cardReader.connectWithUsbCable()
}

type USB interface {
	connectWithUsbCable()
}

type MemoryCard struct{}

func (card *MemoryCard) insert() {
	fmt.Println("Карта памяти вставлена")
}

func (card *MemoryCard) copyData() {
	fmt.Println("Данные скопированы")
}

type CardReader struct {
	card MemoryCard
}

func (reader CardReader) connectWithUsbCable() {
	reader.card.insert()
	reader.card.copyData()
}
