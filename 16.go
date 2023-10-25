package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	arr := createArr()
	emptyArr := make([]int, len(arr))

	copy(emptyArr, arr)
	start := time.Now()
	quicksort(emptyArr)
	duration := time.Since(start)
	fmt.Println("Реализация quicksort - O(n x log n) занимает:", duration)

	copy(emptyArr, arr)
	start = time.Now()
	sort.Ints(emptyArr)
	duration = time.Since(start)
	fmt.Println("Метод из библиотеки sort - O(n x log n) занимает:", duration)
}

func quicksort(arr []int) {
	var sortLH func(arr []int, low, high int)

	sortLH = func(arr []int, low, high int) {
		if low < high {
			partitionIndex := partition(arr, low, high)
			sortLH(arr, low, partitionIndex-1)
			sortLH(arr, partitionIndex+1, high)
		}
	}

	sortLH(arr, 0, len(arr)-1)
}

func partition(arr []int, low, high int) int {
	middle := (low + high) / 2                      // находим индекс числа находящегося в середине слайса
	arr[high], arr[middle] = arr[middle], arr[high] // меняем средний и последний элемент местами

	pivot := arr[high] // ставим пивот на последний элемент
	leftWall := low    // указываем границу слева

	for i := low; i < high; i++ {
		if arr[i] <= pivot { // ищем элемент, который меньше или равен пивоту
			arr[leftWall], arr[i] = arr[i], arr[leftWall] // меняем элементы местами
			leftWall++                                    //двигаем границу
		}
	}
	arr[leftWall], arr[high] = arr[high], arr[leftWall] // меняем элементы местами
	return leftWall
}

// генерация слайса с 100000 элементов
func createArr() (arr []int) {
	sizeArr := 100000
	arr = make([]int, sizeArr)
	time.Now().UnixNano()
	for i := 0; i < sizeArr; i++ {
		arr[i] = rand.Intn(sizeArr)
	}
	return
}
