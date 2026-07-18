package main

import "fmt"

func sumA(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

func counter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	sum := sumA(23)

	fmt.Println(sum(2))  // 25
	fmt.Println(sum(23)) // 46

	myCouter := counter()

	fmt.Println(myCouter()) // 1
	fmt.Println(myCouter()) // 2
	fmt.Println(myCouter()) // 3
	fmt.Println(myCouter()) // 4
}
