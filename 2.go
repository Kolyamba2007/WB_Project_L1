package main

import (
	"fmt"
	"sync"
)

func main() {
	// Инициализируем слайс
	mas := []int{2, 4, 6, 8, 10}

	// Объявление группы ожидания
	var wg sync.WaitGroup

	// Вызов горутин
	for i := 0; i < len(mas); i++ {
		wg.Add(1) // добавление ожидания одной горутины

		go func(e *int, wg *sync.WaitGroup) {
			defer wg.Done() // уменьшает значение счетчика ожидания на 1
			*e *= *e        // подсчет квадрата
		}(&mas[i], &wg)
	}

	// Этот метод блокирует выполнение кода до того момента, пока счетчик ожидания будет 0
	wg.Wait()

	// Вывод слайса квадратов чисел
	fmt.Println(mas)
}