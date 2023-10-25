package main

import (
	"fmt"
	"unicode"
)

func main() {
	strings := []string{"abcd", "abCdefAaf", "aabcd"}

	for _, str := range strings {
		fmt.Println(checkCharUnique(str))
	}
}

func checkCharUnique(str string) (string, bool) {
	hash := make(map[rune]struct{}) // инициализация мапы

	for _, v := range str { // проходимся по символам
		v = unicode.ToLower(v) // меняем на нижний регистр

		if _, ok := hash[v]; !ok { // ищем символ в хэше
			hash[v] = struct{}{} // если не нашли, записываем в хэш
		} else {
			return str, false // если нашли, это говорит о повторении символов, значит возвращаем false
		}
	}

	return str, true // если ни разу не нашли повторяющийся символ, возвращаем true
}
