package main

import "fmt"

func main() {
	str := "главрыба"
	fmt.Println(str)
	fmt.Println(reverse(str))
}

func reverse(s string) string {
	runes := []rune(s) // инициализация слайса из символов

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i] // переворачиваем элементы
	}

	return string(runes) // вывод символов в виде строки
}
