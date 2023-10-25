package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6}

	s = removeIndex(s, 2)
	fmt.Println(s)
}

func removeIndex(s []int, index int) []int {
	ret := make([]int, 0)              // инициализируем слайс
	ret = append(ret, s[:index]...)    // добавляем в ret из s элементы с начала до индекса (сам элемент с индексом не берем)
	return append(ret, s[index+1:]...) // добавляем в ret из s элементы с элемента следующего за индексом и до конца, возвращаем слайс
}
