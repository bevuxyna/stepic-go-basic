//Напиши программу, которая считывает одну строку-разделитель и три строки, а затем выводит их на одной строке, разделяя заданным разделителем.
//
//
//Формат входных данных:
//Четыре строки:
//Первая — разделитель;
//Затем — три строки с произвольным текстом.

//Формат выходных данных:
//Одна строка: три введённых строки, разделённых заданным разделителем.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	separator := scanner.Text()

	scanner.Scan()
	line1 := scanner.Text()

	scanner.Scan()
	line2 := scanner.Text()

	scanner.Scan()
	line3 := scanner.Text()

	fmt.Print(line1, separator, line2, separator, line3)
}
