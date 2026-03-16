package main

import "fmt"

func printslice[T any](s []T)  {
for _,v := range s {
	fmt.Println(v)
}
}
func main() {
	intslice :=[]int{1,2,3,4,5}
	stringslice :=[]string{"mangoes","bananas","oranges"}
	floatslice :=[]float64{34.45,65.78,89.90}

	printslice(intslice)
	printslice(stringslice)
	printslice(floatslice)
} 