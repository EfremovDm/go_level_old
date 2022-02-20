package lesson5

func Fibbonachi(n uint) uint {

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return Fibbonachi(n-1) + Fibbonachi(n-2)
}
