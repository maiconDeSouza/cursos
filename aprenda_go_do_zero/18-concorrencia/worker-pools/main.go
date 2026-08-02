package main

import (
	"fmt"
	"sync"
	"time"
)

func fibonacci(p uint) uint {
	if p <= 1 {
		return p
	}

	return fibonacci(p-2) + fibonacci(p-1)
}

func worker(exit chan<- uint, entry <-chan uint) {
	for en := range entry {
		exit <- fibonacci(en)
	}
}

func main() {
	wg := sync.WaitGroup{}
	v := 40
	w := 6
	start := time.Now()

	entry := make(chan uint, v)
	exit := make(chan uint, v)

	for i := 1; i <= v; i++ {
		p := uint(i)
		entry <- p
	}

	close(entry)

	for i := 1; i <= w; i++ {
		wg.Go(func() {
			worker(exit, entry)
		})
	}

	go func() {
		wg.Wait()
		close(exit)
	}()

	for ex := range exit {
		fmt.Println(ex)
	}

	t := time.Since(start).Seconds()
	fmt.Printf("Demorou %2.fs\n", t)
}
