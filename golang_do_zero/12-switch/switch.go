package main

import "fmt"

func dayOfTheWeek(day int8) string {
	switch day {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sábado"
	default:
		return "Dia errado!"
	}
}

func dayOfTheWeek2(day int8) string {
	var response string
	switch {
	case day == 1:
		response = "Domingo"
	case day == 2:
		response = "Segunda-feira"
	case day == 3:
		response = "Terça-feira"
	case day == 4:
		response = "Quarta-feira"
	case day == 5:
		response = "Quinta-feira"
	case day == 6:
		response = "Sexta-feira"
	case day == 7:
		response = "Sábado"
	default:
		response = "Dia errado!"
	}
	return response
}

func dayOfTheWeek3(day int8) string {
	var response string
	switch {
	case day == 1:
		response = "Domingo"
		fallthrough
	case day == 2:
		response = "Segunda-feira"
	case day == 3:
		response = "Terça-feira"
	case day == 4:
		response = "Quarta-feira"
	case day == 5:
		response = "Quinta-feira"
	case day == 6:
		response = "Sexta-feira"
	case day == 7:
		response = "Sábado"
	default:
		response = "Dia errado!"
	}
	return response
}

func main() {
	r := dayOfTheWeek(3)
	r2 := dayOfTheWeek2(1)
	r3 := dayOfTheWeek3(1)

	fmt.Println(r)
	fmt.Println(r2)
	fmt.Println(r3)

}
