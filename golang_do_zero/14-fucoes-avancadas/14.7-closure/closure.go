package main

import "fmt"

func sumN1(n1 int8) func(int8) int8 {
	return func(n2 int8) int8 {
		return n1 + n2
	}

}

func main() {
	sum10 := sumN1(10)

	fmt.Println(sum10(5))
	fmt.Println(sum10(12))
	fmt.Println(sum10(-15))
	fmt.Println(sum10(13))

}
