package main

import "fmt"

func main() {
	var a int64

	fmt.Print("Введите целочисленное значение: ")
	fmt.Scan(&a)

	fmt.Println(fmt.Sprintf("Число в двоичном виде: %064b", uint64(a)))

	fmt.Printf("Введите номер бита [0-63], который следует заменить: ")

	var b int8
	fmt.Scan(&b)

	a ^= 1 << b

	fmt.Println(fmt.Sprintf("Число в двоичном виде: %064b", uint64(a)))
	fmt.Printf("Результат: %v", a)
}
