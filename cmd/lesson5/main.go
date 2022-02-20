package main

import (
	"fmt"
	"go_level_1/internal/lesson5"
)

/**
 * 1. Напишите приложение, рекурсивно вычисляющее заданное из стандартного ввода число Фибоначчи.
 */
func main() {

	var number uint

	fmt.Println("Выход из программы: 000")
	fmt.Print("Введите номер числа в последовательности Фибонраччи: ")

	_, _ = fmt.Scan(&number)
	i := 0
	results := map[uint]uint{}

	for i = 0; ; i++ {
		if number != 000 {
			fmt.Println("Результат: ", lesson5.Fibbonachi(number))
			results[number] = lesson5.Fibbonachi(number)

			fmt.Print("Следующее число : ")
			_, _ = fmt.Scan(&number)
		} else if number == 000 {
			fmt.Println("Ниже Вы можете видеть сохраненные результаты Ваших вычислений: ")
			for k, v := range results {
				fmt.Println("Число:", k, "Результат:", v)
			}
			fmt.Println("Конец программы!")
			break
		}
	}
}
