//На плоскости заданы координаты двух точек - A(x1,y1) и B(x2,y2)
//Напиши программу, которая вычисляет расстояние между ними.
//
//Формат входных данных:
//Четыре целых числа x1, y1, x2, y2 - координаты точек A и B, каждое на новой строке.
//
//Формат выходных данных:
//Одно число — расстояние между точками A и B.

package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2 int
	fmt.Scan(&x1)
	fmt.Scan(&y1)
	fmt.Scan(&x2)
	fmt.Scan(&y2)

	distance1 := float64(x1 - x2)
	distance2 := float64(y1 - y2)
	distance := math.Sqrt(distance1*distance1 + distance2*distance2)
	fmt.Print(distance)
}
