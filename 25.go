package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	sleep(3)
	duration := time.Since(start)

	fmt.Println(duration)
}

// time.After() возвращает канал, и сообщение будет отправлено на канал по истечении указанного времени
// блокирует поток, пока в канал не придет сообщение через n секунд
func sleep(n int) {
	<-time.After(time.Duration(n) * time.Second)
}
