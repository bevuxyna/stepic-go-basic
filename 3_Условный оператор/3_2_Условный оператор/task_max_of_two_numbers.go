//Напиши программу, которая находит максимальное из двух целых чисел. Если числа равны, выведи любое из них.

//Формат входных данных:
//Два целых числа, каждое из которых не превышает 10000.

//Формат выходных данных:
//Одно число — максимальное из двух исходных чисел.

package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Scan(&num1)
	fmt.Scan(&num2)

	if num1 > num2 {
		fmt.Println(num1)
	} else if num1 == num2 {
		fmt.Println(num2)
	} else {
		fmt.Println(num2)
	}
}
