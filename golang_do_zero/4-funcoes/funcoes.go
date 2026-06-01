package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func operacoesMatematicas(n1, n2 float32) (int8, int8, float32, float32) {
	v1 := n1 + n2
	v2 := n1 - n2
	v3 := n1 * n2
	v4 := n1 / n2

	return int8(v1), int8(v2), v3, v4
}

func main() {
	result := somar(8, 9)
	fmt.Println(result)

	v1, v2, v3, v4 := operacoesMatematicas(8, 9)
	fmt.Println(v1, v2, v3, v4)

}
