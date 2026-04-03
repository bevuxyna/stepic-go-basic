package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Привет! Как тебя зовут?")

	var name string
	fmt.Scan(&name)

	fmt.Println("Привет,", name)

	// countWords, error := fmt.Scan(&s1, &s2, &s3)
	// countWords - Количество успешно считанных значений

	var lastName, firstName string
	countWords, errors := fmt.Scan(&lastName, &firstName)
	fmt.Println("Привет,", firstName, lastName)
	fmt.Println("Количество введенных строк:", countWords)
	fmt.Println("Ошибки:", errors)

	scanner()
}

func scanner() {
	scanner := bufio.NewScanner(os.Stdin) // создаем экземпляр структуры bufio.Scanner для чтения из консоли
	fmt.Println("Введите строку с пробелами:")
	scanner.Scan()
	text := scanner.Text()
	fmt.Println("Вы ввели:", text)
}
