package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var N int // секунды

	fmt.Print("Введите N секунд работы программы: ")
	fmt.Scan(&N) // ввод времени работы программы в секундах

	c := make(chan int) // инициализация канала

	go func() {
		time.Sleep(time.Duration(N) * time.Second) // отсчет времени
		os.Exit(0)                                 // выход из программы
	}()

	go func() {
		for {
			fmt.Println(<-c) // чтение канала
		}
	}()

	// последовательное отправление данных в канал
	var count int
	for {
		c <- count
		count++
		time.Sleep(100 * time.Millisecond)
	}
}
