//Напиши программу, которая считывает три строки по очереди, а затем выводит их в той же последовательности, каждую на отдельной строке.
//
//
//Формат входных данных:
//Три строки с произвольным текстом.
//
//Формат выходных данных:
//Те же три строки, в том же порядке.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	line1 := scanner.Text()

	scanner.Scan()
	line2 := scanner.Text()

	scanner.Scan()
	line3 := scanner.Text()

	fmt.Println(line1)
	fmt.Println(line2)
	fmt.Println(line3)
}
