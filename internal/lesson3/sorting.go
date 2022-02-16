package lesson3

import "fmt"

/**
 * 1. Познакомьтесь с алгоритмом сортировки вставками. Напишите приложение, которое принимает на вход набор
 * целых чисел и отдаёт на выходе его же в отсортированном виде. Сохраните код, он понадобится нам в дальнейших уроках.
 */
func Sorting() {
	var n int

	slice := []int{}
	fmt.Println("Введите целые числа до 1000 шт:")
	fmt.Scan(&n)

	// заполняем динамический массив
	for i := 0; i < 1001; i++ {
		if n != 1000 {
			fmt.Scan(&n)
			slice = append(slice, n)
		} else {
			var result = slice[0 : len(slice)-1]
			fmt.Println(result)
			break
		}
	}

	// сортировка вставками по возрастанию
	for i := 0; i < len(slice); i++ {
		x := slice[i]
		j := i
		for ; j >= 1 && slice[j-1] > x; j-- {
			slice[j] = slice[j-1]
		}
		slice[j] = x
	}

	sort1 := slice[0 : len(slice)-1]
	fmt.Println(sort1)
}