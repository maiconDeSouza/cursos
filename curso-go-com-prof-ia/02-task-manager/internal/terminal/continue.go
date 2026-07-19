package terminal

import "fmt"

func Continue() {
	fmt.Println()
	fmt.Println("Pressione ENTER para retomar...")
	globalScanner.Scan()
}
