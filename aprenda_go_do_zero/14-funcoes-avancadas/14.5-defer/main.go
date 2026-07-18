package main

import (
	"fmt"
	"math/rand"
	"slices"
	"strings"
	"time"
)

func megaSena() string {
	defer fmt.Println("Os números sorteados são!")
	var s = make(map[int]int)

	for len(s) < 6 {
		n := rand.Intn(60) + 1
		if _, ok := s[n]; ok {
			continue
		}

		s[n] = n
	}
	var r []int

	for _, v := range s {
		r = append(r, v)
	}

	slices.Sort(r)

	var builder strings.Builder

	for i, v := range r {
		if i == len(r)-1 {
			fmt.Fprintf(&builder, "%d", v)
			continue
		}
		fmt.Fprintf(&builder, "%d-", v)
	}

	time.Sleep(time.Second * time.Duration(6))
	return builder.String()
}

func main() {
	fmt.Println(megaSena())
}
