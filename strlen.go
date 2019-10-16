package piscine

func StrLen(str string) int {
	C := 0
	A := []rune(str)
	for P := range A {
		C = P + 1
	}
	return C
}
