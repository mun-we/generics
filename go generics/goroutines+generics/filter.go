package main

import (
	"fmt"
	"sync"
)

type number interface {
	int | int16 | int32 | int64 | float32 | float64
}

func sum[T number](num []T, w *sync.WaitGroup) {
	defer w.Done()
	var total T
	for _, numbers := range num {
		total += numbers
	}
	fmt.Println(total)
}
func main() {
	var w sync.WaitGroup
	ints := []int{1, 2, 3, 4, 5, 7, 6, 7, 8, 9}
	floats := []float64{2.45, 45.56, 78.32, 2.56}
	w.Add(1)
	go sum(ints, &w)
	w.Add(1)
	go sum(floats, &w)
	w.Wait()

}
