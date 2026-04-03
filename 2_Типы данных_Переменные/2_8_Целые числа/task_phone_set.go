//Напиши программу, которая вычисляет общую стоимость трех комплектов телефона.
//Один комплект включает: само устройство, чехол, зарядку и наушники.

//Формат входных данных:
//Четыре целых числа, каждое на новой строке:
//1. Стоимость устройства
//2. Стоимость чехла
//3. Стоимость зарядки
//4. Стоимость  наушников

//Формат выходных данных:
//Одно целое число — стоимость трёх полных комплектов.

package main

import "fmt"

func main() {
	var phone, phoneCase, charge, headphones int

	fmt.Scan(&phone)
	fmt.Scan(&phoneCase)
	fmt.Scan(&charge)
	fmt.Scan(&headphones)

	sum := phone + phoneCase + charge + headphones

	fmt.Print(sum * 3)
}
