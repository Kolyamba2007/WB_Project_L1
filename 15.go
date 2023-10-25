package main

import (
	"fmt"
	"os"
)

const fileName string = "huge_str.txt"

// Здесь приведен пример работы с очень большими строками,
// которые могут не поместиться в памяти, работа с ними так же будет не быстрая.
// Такую строку можно, например, записывать в файл и читать по мере необходимости.
func main() {
	someFunc1()
	someFunc2()
}

func someFunc1() {
	f, err := os.Create(fileName) // создаем файл
	if err != nil {
		panic(err)
	}
	defer f.Close()

	createHugeString(1<<10, f) // Передаем файл в качестве io.Writer, чтобы записать строку в него; (1<<16 = 2^16)
}

func someFunc2() {
	f, err := os.Open(fileName) // открытие файла
	if err != nil {
		panic(err)
	}
	defer f.Close() // закрытие файла в конце стека

	str := make([]byte, 100) // Создаем слайс с размером 100

	f.Read(str) // Считываем файл в слайс

	fmt.Println(string(str)) // Вывод результата
}

// метод, в котором записывается строка в файл
func createHugeString(c int, f *os.File) {
	for i := 0; i < c; i++ {
		_, err := f.WriteString("R")
		if err != nil {
			panic(err)
		}
	}
}
