package main

import "fmt"

func invertSign(n *int) {
	*n = *n * -1
}

func arrInvertSign(arr *[5]int) {
	for i, value := range *arr {
		invertSign(&value)
		arr[i] = value
	}
}

func main() {
	a := -5
	b := 23
	c := 10
	d := [5]int{-1, -2, -3, -4, -5}

	invertSign(&a)
	invertSign(&b)
	invertSign(&c)
	arrInvertSign(&d)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
