package main

import "fmt"

func inverterSinal(v int) int {
	r := v * -1
	return r
}

func inverterSinalPonteiro(v *int) {
	*v = *v * -1
}

func main() {
	a := 42
	b := -23

	r := inverterSinal(a)

	fmt.Println("O valor de a", a)            // O valor de a 42
	fmt.Println("O valor do retorno de a", r) // O valor do retorno de a -42
	fmt.Println("Agora com Ponteiros")        // Agora com Ponteiros
	inverterSinalPonteiro(&b)
	fmt.Println("O valor de b", b) // O valor de b 23

}
