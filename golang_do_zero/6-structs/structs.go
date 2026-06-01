package main

import "fmt"

type user struct {
	name string
	age  uint8
}

func main() {
	var u user
	u.name = "Dante"
	u.age = 5

	u2 := user{"Dona Maia", 13}

	u3 := user{age: 10}

	fmt.Println(u)
	fmt.Println(u2)
	fmt.Println(u3)

}
