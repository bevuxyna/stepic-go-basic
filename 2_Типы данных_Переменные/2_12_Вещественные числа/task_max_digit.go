//Дано целое трехзначное число. Необходимо найти максимальную цифру этого числа.

//Формат входных данных:
//Одно целое трехзначное число.

//Формат выходных данных:
//Одна цифра - максимальная цифра в составе введенного трехзначного числа.

package main

import (
	"fmt"
	"math"
)

func main() {
	var num int
	fmt.Scan(&num)

	firstDigit := float64(num / 100)
	secondDigit := float64((num % 100) / 10)
	thirdDigit := float64(num % 10)

	maxDigit := math.Max(math.Abs(firstDigit), math.Max(math.Abs(secondDigit), math.Abs(thirdDigit)))

	fmt.Print(maxDigit)
}
