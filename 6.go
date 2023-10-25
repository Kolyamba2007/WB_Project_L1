package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//1 Способ завершения горутины через канал
	ch := make(chan int)        // инициализация канала
	quit := make(chan struct{}) // инициализация канала
	go func() {
		for {
			// конструкция select-case
			select {
			case ch <- 1:
			case <-quit:
				fmt.Println("1-ая горутина завершила работу")
				close(ch)   // закрытие канала
				close(quit) // закрытие канала
				return
			default:
				// …
			}
		}
	}()
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	quit <- struct{}{} // триггер закрытия канала
	time.Sleep(100 * time.Millisecond)

	//2 Способ завершения горутины через канал, который также используется для чтения
	number := func() chan int {
		ch := make(chan int) // инициализация канала
		go func() {
			n := 1
			for {
				// конструкция select-case
				select {
				case ch <- n:
					n++
				case <-ch:
					fmt.Println("2-ая горутина завершила работу")
					close(ch) // закрытие канала
					return
				default:
					// …
				}
			}
		}()
		return ch
	}()
	fmt.Println(<-number)
	fmt.Println(<-number)
	number <- 0 // триггер закрытия канала
	time.Sleep(100 * time.Millisecond)

	//3 Способ завершения горутины через контекст
	forever := make(chan string)                            // инициализация канала
	ctx, cancel := context.WithCancel(context.Background()) // инициализация контекста, с помощью которого будет завершаться горутина

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done(): // если сработал cancel()
				forever <- "3-я горутина завершила работу"
				return
			default:
				fmt.Println("for loop")
			}

			time.Sleep(500 * time.Millisecond)
		}
	}(ctx)

	go func() {
		time.Sleep(1 * time.Second)
		cancel() // вызов контекста
	}()

	fmt.Println(<-forever)
}
