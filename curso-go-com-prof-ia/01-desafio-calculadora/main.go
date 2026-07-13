package main

import (
	"errors"
	"fmt"
)

func operar(a, b float64, operador string) (float64, error) {
	var result float64
	switch operador {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "/":
		if b <= 0 {
			return 0, errors.New("divisão por zero")
		}
		result = a / b
	case "*":
		result = a * b
	default:
		return 0, errors.New("Operador não aceito")
	}

	return result, nil
}

func media(notas ...float64) (float64, error) {
	var sum float64

	for _, n := range notas {
		if n < 0 || n > 10 {
			return 0, errors.New("A nota deve ser entre 0 e 10")
		}
		sum += n
	}

	result := sum / float64(len(notas))

	return result, nil
}

func result(res float64, err error) {
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}

	fmt.Printf("%.2f\n", res)
}

func main() {
	result(operar(23, 5, "+"))  // 28.00
	result(operar(23, 5, "-"))  // 18.00
	result(operar(23, 0, "/"))  // divisão por zero
	result(operar(23, 3, "/"))  // 7.67
	result(operar(23, 0, "+"))  // 23.00
	result(operar(23, 23, "*")) // 529.00
	result(operar(23, 23, "#")) // Operador não aceito

	result(media(2, 4, 6, 10, -1))   // A nota deve ser entre 0 e 10
	result(media(2, 4, 6, 10, 10.1)) // A nota deve ser entre 0 e 10
	result(media(2, 4, 6, 10, 10))   // 6.40

}
