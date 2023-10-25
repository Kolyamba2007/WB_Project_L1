package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"
	fmt.Println(str)
	fmt.Println(reverseWords(str))
}

func reverseWords(s string) string {
	words := strings.Split(s, " ") // разделение строки на элементы слайса, где " " знак для разделения

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i] // переворачиваем элементы
	}

	return strings.Join(words, " ") // соединение строк слайса в одну с разделителем " "
}
