package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Counter struct {
	i int32 // инкрементируемое значение
}

func main() {
	var counter Counter   // инициализация структуры счетчика
	var wg sync.WaitGroup // объявление группы ожидания

	ctx, cancel := context.WithCancel(context.Background()) // инициализация контекста, с помощью которого будут завершаться все горутины

	wg.Add(10) // добавление 10 воркеров в группу ожидания

	for i := 0; i < 10; i++ {
		go func(counter *Counter, wg *sync.WaitGroup) {
			defer wg.Done() // уменьшает значение счетчика ожидания на 1
			for {
				atomic.AddInt32(&counter.i, 1) // атомики позволят конкурентно читать и писать данные без блокировок
				select {
				case <-ctx.Done(): // если сработал cancel(), то выходим из горутины
					return
				default: // если контекст выше не выполнен, продолжаем цикл
					time.Sleep(100 * time.Millisecond)
					continue
				}
			}
		}(&counter, &wg)
	}

	// инкрементируем в течение какого-то времени до вызова контекста
	go func() {
		time.Sleep(4 * time.Second)
		cancel() // вызов контекста
	}()

	wg.Wait() // этот метод блокирует выполнение кода до того момента, пока счетчик ожидания будет 0

	fmt.Println(counter.i) // вывод инкрементируемого значения
}
