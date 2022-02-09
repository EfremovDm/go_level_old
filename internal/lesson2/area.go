package lesson2

import (
	"fmt"
)

/**
 * 1. Напишите программу для вычисления площади прямоугольника.
 * Длины сторон прямоугольника должны вводиться пользователем с клавиатуры.
 */
func Area() {

	var a,b,s float64

	fmt.Println("1. Вычисление площади прямоугольника")
	fmt.Print("Введите длину стороны a: ")
	fmt.Scan(&a)
	fmt.Print("Введите длину стороны b: ")
	fmt.Scan(&b)
	s = a * b;
	fmt.Print("Площадь прямоугольника: ", s)

	var input string
	fmt.Scanf("%v",&input)
}