//По заданному номеру месяца определи количество дней в этом месяце.

//Формат входных данных:
//Одно целое число от 1 до 12 — номер месяца.

//Формат выходных данных:
//Одно число — количество дней в заданном месяце.

package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	switch num {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println(31)
	case 2:
		fmt.Println(29)
	default:
		fmt.Println(30)
	}
}
