package main

import "fmt"

// целевой интерфейс - AnimalAdapter
type AnimalAdapter interface {
	Reaction()
}

// класс собака
type Dog struct{}

// реакция собаки
func (dog *Dog) WoofWoof() {
	fmt.Println("Гав-Гав")
}

// адаптер для собаки
type DogAdapter struct {
	*Dog
}

// реакция собаки
func (adapter *DogAdapter) Reaction() {
	adapter.WoofWoof()
}

// конструктор адаптера для собаки
func newDogAdapter(dog *Dog) AnimalAdapter {
	return &DogAdapter{dog}
}

// класс кошка
type Cat struct{}

// реакция кошки, если ее позвать
func (dog *Cat) MeowMeow(isCall bool) {
	if isCall {
		fmt.Println("Мяу-Мяу")
	}
}

// адаптер для кошки
type CatAdapter struct {
	*Cat
}

// реакция кошки
func (adapter *CatAdapter) Reaction() {
	adapter.MeowMeow(true)
}

// конструктор адаптера для кота
func newCatAdapter(cat *Cat) AnimalAdapter {
	return &CatAdapter{cat}
}

func main() {
	dogAdapter := newDogAdapter(&Dog{})
	dogAdapter.Reaction()

	catAdapter := newCatAdapter(&Cat{})
	catAdapter.Reaction()
}
