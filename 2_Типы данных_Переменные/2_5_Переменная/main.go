package main

import "fmt"

func main() {
	var a int = 5
	b := 10

	a = b - 2

	var c = a * b
	fmt.Println(c)

	name := "Олег"
	age := 29
	weight := 80.9

	fmt.Println("Имя:", name, "Возраст:", age, "Вес:", weight)
}
