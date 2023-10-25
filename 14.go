package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(determineType(1))
	fmt.Println(determineType("string"))
	fmt.Println(determineType(true))
	fmt.Println(determineType(make(chan int)))
}

// Ограничение на используемые типы any
// any - аналог interface{} (синтаксический сахар). Это ключевое слово можно использовать в любом месте.
func determineType[T any](a T) reflect.Type {
	return reflect.TypeOf(a) // reflect.TypeOf() определяет тип
}
