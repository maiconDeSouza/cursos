package main

import "fmt"

func main() {
	user := map[string]string{
		"name":  "Dante",
		"breed": "Pug",
	}

	v, ok := user["breed"]
	if !ok {
		fmt.Println("Valor não existe!")
		return
	}

	fmt.Println(v)
}
