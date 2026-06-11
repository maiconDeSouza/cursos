package main

import (
	"fmt"
	"math/rand"
	"tickets/db"
)

var names = db.Names
var lastName = db.LastName

func generateName() string {
	iName := rand.Intn(100)
	iLastName := rand.Intn(100)

	return fmt.Sprintf("%s %s", names[iName], lastName[iLastName])

}

type Fan struct {
	id   int
	name string
}

type TicketOffice struct {
	tickets int
	buyers  []*Fan
}

func CreateFan(id int, name string, f *Fan, t *TicketOffice) {
	f.id = id
	f.name = name

	t.buyers = append(t.buyers, f)
}

func main() {
	for i := 1; i <= 1000; i++ {
		fmt.Println(generateName())
	}
}
