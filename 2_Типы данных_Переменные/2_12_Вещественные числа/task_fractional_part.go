//Дано положительное вещественное число x. Выведи его дробную часть.

//Формат входных данных:
//Одно положительное вещественное число x

//Формат выходных данных:
//Одно число — дробная часть числа x

package main

import "fmt"

func main() {
	var num float64
	fmt.Scan(&num)
	integerPart := int(num)
	fmt.Print(num - float64(integerPart))
}
