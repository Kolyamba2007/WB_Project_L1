package main

import (
	"fmt"
	"sync"
)

func main() {
	mas := [5]int{2, 4, 6, 8, 10} // инициализация массива

	ch1 := make(chan int) // инициализация первого канала
	ch2 := make(chan int) // инициализация второго канала

	var wg sync.WaitGroup // объявление группы ожидания
	wg.Add(3)             // добавление ожидания трех горутин

	go func() {
		defer wg.Done() // уменьшает значение счетчика ожидания на 1
		for i := 0; i < len(mas); i++ {
			ch1 <- mas[i] // запись элемента массив в 1-ый канал
		}
	}()

	go func() {
		defer wg.Done() // уменьшает значение счетчика ожидания на 1
		for i := 0; i < len(mas); i++ {
			mes := <-ch1 // чтение элемента массива из 1-ого канала
			mes *= 2     // элемент, умноженный на 2
			ch2 <- mes   // запись модифицированного элемента во 2-ой канал
		}
	}()

	go func() {
		defer wg.Done() // уменьшает значение счетчика ожидания на 1
		for i := 0; i < len(mas); i++ {
			fmt.Println(<-ch2) // чтение из 3-го канала
		}
	}()

	wg.Wait() // этот метод блокирует выполнение кода до того момента, пока счетчик ожидания будет 0
}
