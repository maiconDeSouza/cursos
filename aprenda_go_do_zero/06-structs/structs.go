package main

import "fmt"

type User struct {
	name string
	age  uint8
	Address
}

type Address struct {
	street string
	number uint16
}

func main() {
	var u1 User
	var u2 = User{
		name: "Dante",
		age:  5,
		Address: Address{
			street: "Rua A",
			number: 123,
		},
	}
	var u3 = User{}
	u4 := User{
		name: "Dona Maia",
		age:  12,
	}
	u5 := User{}

	u1.name = "Delinha"
	u1.age = 8
	u1.street = "Rua O"
	u1.number = 123

	u5.Address.street = "Rua U"
	u5.Address.number = 123

	fmt.Println(u1)
	fmt.Println(u2)
	fmt.Println(u3)
	fmt.Println(u4)
	fmt.Println(u5)

}
