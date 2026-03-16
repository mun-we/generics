package main

import "fmt"

type numbers interface {
	int | int16 | int32 | int64 | float32 | float64
}

func large[T numbers](s []T) T {
	max := s[0]

	for _, num := range s {
		if num > max {
			max = num
		}
	}
	return max
}
func main() {
	ints := []int{1, 7, 8, 9, 34, 6, 5}
	floats := []float64{58.76, 67.85, 56.89}

	k := large(ints)
	fmt.Println(k)
	l := large(floats)
	fmt.Println(l)

}
