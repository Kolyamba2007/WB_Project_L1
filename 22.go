package main

import (
	"fmt"
	"math/big"
)

// использование библиотеки big; в обычном случае при больших числах out of range
func main() {
	a := big.NewInt(1 << 62)
	b := big.NewInt(1 << 50)

	fmt.Println(multiplication(a, b))
	fmt.Println(division(a, b))
	fmt.Println(substraction(a, b))
	fmt.Println(addition(a, b))
}

func multiplication(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b) // перемножение
}

func division(a, b *big.Int) *big.Int {
	return new(big.Int).Div(a, b) // деление
}

func substraction(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b) // разность
}

func addition(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b) // сумма
}
