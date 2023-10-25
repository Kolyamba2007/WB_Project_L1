package main

import (
	"fmt"
	"math"
)

// структура точки с инкапсулированными x,y
type Point struct {
	x float64
	y float64
}

// конструктор
func createPoint(x, y float64) *Point {
	p := new(Point)
	p.x = x
	p.y = y
	return p
}

// метод получения расстояния между точками
func getDistance(p1, p2 *Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
}

func main() {
	p1 := createPoint(-3.24, -2)
	p2 := createPoint(4.4, 6.12)

	d := getDistance(p1, p2)
	fmt.Println(d)
}
