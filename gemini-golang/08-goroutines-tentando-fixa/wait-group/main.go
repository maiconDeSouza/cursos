// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func Falar(pessoa string) {
// 	fmt.Println(pessoa, "começou a falar...")
// 	time.Sleep(time.Second)
// 	fmt.Println(pessoa, "terminou!")
// }

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	go func() {
// 		defer wg.Done()
// 		Falar("Maria")
// 	}()
// 	go func() {
// 		defer wg.Done()
// 		Falar("João")
// 	}()

// 	wg.Wait()
// }

package main

import (
	"fmt"
	"sync"
	"time"
)

func Falar(pessoa string) {
	fmt.Println(pessoa, "começou a falar...")
	time.Sleep(time.Second)
	fmt.Println(pessoa, "terminou!")
}

func main() {
	var wg sync.WaitGroup
	wg.Go(func() {
		Falar("Maria")
	})

	wg.Go(func() {
		Falar("João")
	})

	wg.Wait()
}
