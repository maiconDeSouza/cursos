package main

import "fmt"

type User struct {
	name string
	age  uint8
}

func (u User) salvar() {
	fmt.Printf("Salvando o usuário %s", u.name)
}

func (u User) maiorDeIdade() bool {
	return u.age >= 18
}

func (u *User) aniversario() {
	u.age++
}

func main() {
	u := User{name: "João", age: 17}
	fmt.Println(u)

	fmt.Println("Maior de idade:", u.maiorDeIdade())
	u.aniversario()
	fmt.Println("Maior de idade:", u.maiorDeIdade())
	fmt.Println(u)

}
