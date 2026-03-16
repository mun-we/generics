package main

import "fmt"

func mapslice[T any, R any](input []T, fn func(T) R) []R {
	var reslt []R
	for _, v := range input {
		reslt = append(reslt, fn(v))
	}
	return reslt
}
func main() {
	nums := []int{1, 2, 3, 4, 5}
	strings := mapslice(nums, func(n int) string {

		return fmt.Sprintf("number %d", n)
	})
	fmt.Println(strings)

}
