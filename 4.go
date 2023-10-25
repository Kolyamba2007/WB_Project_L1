package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	var workerPoolSize int // объявление величины пула воркеров

	fmt.Print("Введите кол-во воркеров: ")
	fmt.Scan(&workerPoolSize) // ввод величины пула воркеров

	c := make(chan int)                                     // инициализация основного канала
	ctx, cancel := context.WithCancel(context.Background()) // инициализация контекста, с помощью которого будут завершаться все горутины

	var wg sync.WaitGroup  // объявление группы ожидания
	wg.Add(workerPoolSize) // добавление кол-ва воркеров в группу ожидания

	for i := 0; i < workerPoolSize; i++ {
		// анонимная горутина
		go func(c *chan int, wg *sync.WaitGroup) {
			defer wg.Done() // уменьшает значение счетчика ожидания на 1
			for {
				if mes, ok := <-*c; ok {
					fmt.Println(mes) // если есть что прочитать, выводим сообщение
				} else {
					return // если канал закрыт и из него нечего читать, то выходим из горутины
				}
			}
		}(&c, &wg)
	}

	go func(wg *sync.WaitGroup) {
		osCh := make(chan os.Signal, 1)   // инициализация канала для сигнала закрытия программы
		signal.Notify(osCh, os.Interrupt) //оповещение, если сработает выход из процесса
		// обработка выхода из процесса
		for sig := range osCh {
			cancel()  // вызов контекста
			wg.Wait() // этот метод блокирует выполнение кода до того момента, пока счетчик ожидания будет 0
			fmt.Println(sig)
			os.Exit(0) // выход из программы
		}
	}(&wg)

	// постоянная запись данных в канал
	var count int
	for {
		c <- count // запись
		count++
		select {
		case <-ctx.Done(): // если сработал cancel()
			close(c)  // закрытие канала
			select {} // блокировка текущей горутины навсегда
		default: // если контекст выше не выполнен, продолжаем цикл
			time.Sleep(100 * time.Millisecond)
			continue
		}
	}
}
