package main

import "fmt"

var arr1 [5]int8
var slice1 []int8

func main() {
	arr1[0] = 1
	slice1 = append(slice1, 23)

	arr2 := [...]string{"Go", "Python", "JS"}
	slice2 := []string{"Lua", "Rust"}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(slice1)
	fmt.Println(slice2)

	slice3 := make([]string, 0, 4)
	slice3 = append(slice3, "Dante", "Dona Maia")
	fmt.Println(slice3, len(slice3), cap(slice3))
	slice3 = append(slice3, "Delinha", "Taluzinha", "Tachorrin")
	fmt.Println(slice3, len(slice3), cap(slice3))

}
