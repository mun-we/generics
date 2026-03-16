package main

import "fmt"

func filter[T comparable](s []T, key T) bool {
	for _, value := range s {
		if key == value {
			return true
		}
	}
	return false
}
func main() {
	strings := []string{"peter", "mary", "kelvin", "james"}
	ints := []int{21, 2, 3, 45, 6, 78}
	j := filter(strings, "kelvin")
	fmt.Println(j)
	l := filter(ints, 33)
	fmt.Println(l)
}
