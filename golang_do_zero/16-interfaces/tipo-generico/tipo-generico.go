package main

import "fmt"

func generica(g interface{}) {
	fmt.Println(g)
}
func main() {
	generica(23)
	generica("Filme")
	generica(true)
}
