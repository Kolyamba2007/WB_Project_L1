package main

import "fmt"

func main() {
	slice := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(unique(slice))
}

// Ограничение comparable позволяет использовать любые типы,
// где значения можно сравнивать с помощью операторов == and !=.
// Учитывая что ключи в мапе обязательно должны быть сопоставимы,
// то необходимо указывать T как comparable для использования его в качестве ключа.
func unique[T comparable](a []T) []T { // используем дженерик, метод возвращает слайс
	set := make([]T, 0)          // инициализация слайса
	hash := make(map[T]struct{}) // инициализация мапы

	// ищем элементы слайса a в хэше
	for _, v := range a {
		if _, ok := hash[v]; !ok { // если не нашли, то записываем элемент в хэш и результирующий слайс
			hash[v] = struct{}{}
			set = append(set, v)
		}
	}

	return set
}
