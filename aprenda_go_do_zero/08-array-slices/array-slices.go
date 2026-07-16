package main

import "fmt"

func main() {
	var arr1 [5]int
	var arr2 = [5]int{1, 2, 3, 4, 5}
	arr3 := [...]string{"a", "b", "c", "d", "e"}
	fmt.Println(arr1) // [0 0 0 0 0]
	fmt.Println(arr2) // [1 2 3 4 5]
	fmt.Println(arr3) // [a b c d e]

	fmt.Println("Slices")
	slc := []int{1, 2, 3}
	fmt.Printf("O slice %v e tem o tamanho %d com a capacidade %d\n", slc, len(slc), cap(slc))
	// O slice [1 2 3] e tem o tamanho 3 com a capacidade 3

	slc = append(slc, 4)
	fmt.Printf("O slice %v e tem o tamanho %d com a capacidade %d\n", slc, len(slc), cap(slc))
	// O slice [1 2 3 4] e tem o tamanho 4 com a capacidade 6

	slc2 := arr3[1:3]
	fmt.Println(slc2) // [b c]

	slc2 = append(slc2, "ç")
	fmt.Println(slc2) // [b c ç]

	arr3[1] = "Alterado"
	fmt.Println(slc2) // [Alterado c ç]

	slc3 := slc[len(slc)-1:]
	fmt.Println(slc3) //[4]

	slc4 := slc
	slc5 := make([]int, len(slc))
	copy(slc5, slc)
	slc[0] = 23

	slc[len(slc)-1] = 88
	fmt.Println(slc3) //[88]
	fmt.Println(slc4) //[23 2 3 88]
	fmt.Println(slc5) // [1 2 3 4]

	slc4 = append(slc3, 255)
	slc4[0] = 77
	fmt.Println(slc)  //[23 2 3 77]
	fmt.Println(slc4) //[77 255]
	fmt.Println(slc5) //[77 255]

}
