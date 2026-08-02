package main

import (
	"fmt"
	"sync"
)

func sequence(start, end int) <-chan int {
	s := make(chan int)

	go func() {
		for i := start; i <= end; i++ {
			s <- i
		}
		close(s)
	}()

	return s
}

func multiplexar(channels ...<-chan int) <-chan int {
	out := make(chan int)
	wg := sync.WaitGroup{}

	for _, ch := range channels {
		wg.Go(func() {
			for n := range ch {
				out <- n
			}
		})
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	s1 := sequence(1, 5)
	s2 := sequence(10, 15)
	s3 := sequence(100, 105)

	out := multiplexar(s1, s2, s3)

	for n := range out {
		fmt.Println(n)
	}
}
