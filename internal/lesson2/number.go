package lesson2

import (
	"fmt"
	"math"
)

/**
 * 3. С клавиатуры вводится трехзначное число.
 *    Выведите цифры, соответствующие количество сотен, десятков и единиц в этом числе.
 */
func Number() {

	var number, s, d, e int64

	fmt.Println("3. С клавиатуры вводится трехзначное число.")

	fmt.Print("Введите число: ")
	fmt.Scan(&number)

	s = int64(math.Floor(float64(number) / 100))
	d = int64(math.Floor(float64(number - s * 100) / 10))
	e = number - s * 100 - d * 10;

	fmt.Println("Сотни: ",  s)
	fmt.Println("Десятки: ", d)
	fmt.Println("Единицы: ", e)
}