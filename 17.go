package main

import "fmt"

func main() {
	arr := fillArr()

	if value, ok := binarySearch(arr, 437); ok {
		fmt.Println("Элемент успешно найден: ", value)
	} else {
		fmt.Println("Элемент не найден!")
	}
}

func binarySearch(array []int, lookingFor int) (int, bool) {
	low := 0               // левая граница
	high := len(array) - 1 // правая граница

	for low <= high { // выполняем пока индекс левой границы не стал больше правой
		mid := (low + high) / 2 // находим индекс среднего элемента

		if array[mid] == lookingFor { // если средний элемент оказался искомым, возвращаем его и true
			return array[mid], true
		}

		// Усли средний элемент оказался больше искомого,
		// то выбираем первую половину слайса как следующий диапазон поиска, сдвигая правую границу.
		// Иначе выбираем вторую половину слайса, сдвигая левую границу.
		if array[mid] > lookingFor {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return 0, false // если не нашли, возвращаем 0 и false
}

// заполнение слайса отсортированными данными
func fillArr() (arr []int) {
	sizeArr := 1000
	arr = make([]int, sizeArr)

	for i := 0; i < sizeArr; i++ {
		arr[i] = i
	}

	return
}
