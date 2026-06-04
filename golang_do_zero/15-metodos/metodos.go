package main

import "fmt"

type User struct {
	Name string
	Age  uint
}

func (u User) nameAndAge() string {
	return fmt.Sprintf("%s tem %d anos\n", u.Name, u.Age)
}

func (u *User) birthday() {
	u.Age++
}

func main() {
	u1 := User{
		Name: "Dante",
		Age:  4,
	}
	fmt.Print(u1.nameAndAge())
	u1.birthday()
	fmt.Print(u1.nameAndAge())
}
