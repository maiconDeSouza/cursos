package main

import "fmt"

type Person struct {
	name string
	age  uint8
}

type Student1 struct {
	Person
	course string
}

type Student2 struct {
	person Person
	course string
}

func main() {
	st1 := Student1{}
	st2 := Student2{}

	st1.name = "Pedro"
	st1.age = 25
	st1.course = "GO"

	st2.person.name = "Maria"
	st2.person.age = 23
	st2.course = "JS"

	fmt.Println(st1)
	fmt.Println(st2)

}
