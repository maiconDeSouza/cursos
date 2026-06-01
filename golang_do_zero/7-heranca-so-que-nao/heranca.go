package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       uint8
}

type student struct {
	person
	registration string
	grade        uint8
}

func main() {
	p := person{firstName: "Dante", lastName: "Kiko", age: 5}
	s := student{person: p, registration: "23456", grade: 3}

	fmt.Println(s)
}
