package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		sum int            // сумма
		wg  sync.WaitGroup // Объявление группы ожидания
	)

	// Инициализируем слайс
	mas := []int{2, 4, 6, 8, 10}
	ch := make(chan int, 4)
	end := make(chan struct{}, 1)

	// Вызов горутин
	for i := 0; i < len(mas); i++ {
		wg.Add(1) // добавление ожидания одной горутины

		go func(e *int, wg *sync.WaitGroup) {
			defer wg.Done() // уменьшает значение счетчика ожидания на 1
			ch <- *e * *e   // записываем квадрат числа
		}(&mas[i], &wg)
	}

	go func() {
		for {
			if mes, ok := <-ch; ok {
				sum += mes // пока есть сообщения, считаем сумму
			} else {
				fmt.Println(sum)  // вывод суммы
				end <- struct{}{} // записываем и выходим из главной горутины
				return            // если канал закрыт и из него нечего читать, то выходим из горутины
			}
		}
	}()

	wg.Wait() // этот метод блокирует выполнение кода до того момента, пока счетчик ожидания будет 0
	close(ch) // закрытие канала
	<-end     // блокировка горутины до следующей записи в канал
}
