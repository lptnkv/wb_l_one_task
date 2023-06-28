package main

import (
	"flag"
	"fmt"
)

func main() {
	// Флаги
	nBitFlag := flag.Int("i", 0, "position of bit to set")
	isZeroFlag := flag.Bool("zero", false, "set bit to zero")
	flag.Parse()
	pos := *nBitFlag
	isZero := *isZeroFlag

	// ... 0101 0011 1101 0101
	var a int64 = 21461
	fmt.Println("Initial value: ", a)
	if isZero {
		mask := ^(1 << pos)
		a &= int64(mask)
	} else {
		a |= (1 << pos)
	}
	fmt.Println("New value: ", a)
}
