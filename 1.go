package main

import "fmt"

// структура Human
type Human struct {
	field1 int
	field2 float32
	field3 string
}

// метод структуры Human
func (h *Human) print() {
	fmt.Printf("Fields: %v, %v, %v", h.field1, h.field2, h.field3)
}

// структура Action
type Action struct {
	Human // анонимное поле структуры Human
}

func main() {
	// инициализация структуры Action
	a := Action{}
	a.field1 = 1
	a.field2 = 3.3
	a.field3 = "field3"

	// вызов метода *Human.print()
	a.print()
}
