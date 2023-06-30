package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	p := Point{x, y}
	return p
}

func GetDistance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
}

func main() {
	p1 := NewPoint(2, -5)
	p2 := NewPoint(-4, 3)
	fmt.Println("Distance between points: ", GetDistance(p1, p2))
}
