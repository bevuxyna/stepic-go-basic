//Дано трехзначное число. Найди сумму его цифр.

//Формат входных данных:
//Одно трехзначное число.

//Формат выходных данных:
//Одно целое число — сумма цифр введенного числа.

package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	first := num / 100
	second := num / 10 % 10
	third := num % 10

	fmt.Println(first + second + third)
}
