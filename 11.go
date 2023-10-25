package main

import "fmt"

func main() {
	slice1 := []int8{2, 4, 6, 8}
	slice2 := []int8{6, 4, 8, 9, 0}
	fmt.Println(intersection(slice1, slice2))
}

// Ограничение comparable позволяет использовать любые типы,
// где значения можно сравнивать с помощью операторов == and !=.
// Учитывая что ключи в мапе обязательно должны быть сопоставимы,
// то необходимо указывать T как comparable для использования его в качестве ключа.
func intersection[T comparable](a []T, b []T) []T { // используем дженерик, метод возвращает слайс
	set := make([]T, 0)          // инициализация слайса
	hash := make(map[T]struct{}) // инициализация мапы

	// записываем элементы слайса a в хэш
	for _, v := range a {
		hash[v] = struct{}{}
	}

	// ищем элементы слайса b в хэше
	for _, v := range b {
		if _, ok := hash[v]; ok {
			set = append(set, v) // если нашли, то записываем элемент в результирующий слайс
		}
	}

	return set
}
