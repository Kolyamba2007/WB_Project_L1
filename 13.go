package main

import "fmt"

func main() {
	a, b := 0, 1 // инициализация a и b

	a, b = b, a // смена чисел

	fmt.Printf("a: %v, b: %v", a, b)
}
