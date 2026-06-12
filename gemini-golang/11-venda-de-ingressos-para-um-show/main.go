package main

import (
	"fmt"
	"math/rand"
	"sync"
	"tickets/db"
	"time"
)

var names = db.Names
var lastName = db.LastName
var wg sync.WaitGroup
var wg2 sync.WaitGroup

func generateName() string {
	iName := rand.Intn(100)
	iLastName := rand.Intn(100)

	return fmt.Sprintf("%s %s", names[iName], lastName[iLastName])

}

type Fan struct {
	id   int
	name string
}

type Online struct {
	onLine []*Fan
}

type TicketOffice struct {
	tickets   int
	buyers    []*Fan
	turnstile sync.RWMutex
}

func (t *TicketOffice) buy(f *Fan) {
	t.turnstile.Lock()
	defer t.turnstile.Unlock()
	if len(t.buyers) >= t.tickets {
		return
	}
	t.buyers = append(t.buyers, f)
}

func CreateFan(id int, name string, f *Fan, o *Online) {
	f.id = id
	f.name = name

	o.onLine = append(o.onLine, f)
}

func queue(chanFan chan *Fan, t *TicketOffice) {
	for fan := range chanFan {
		t.buy(fan)
	}
}

func main() {
	t := TicketOffice{tickets: 50}
	o := Online{}

	for i := 1; i <= 1000; i++ {
		f := Fan{}
		name := generateName()
		CreateFan(i, name, &f, &o)
	}

	chanFan := make(chan *Fan)

	for _, fan := range o.onLine {
		fanT := fan
		purchaseTime := rand.Intn(60) + 1
		wg.Go(func() {
			time.Sleep(time.Duration(purchaseTime) * time.Second)
			chanFan <- fanT
		})
	}

	wg2.Go(func() {
		queue(chanFan, &t)
	})

	wg.Wait()
	close(chanFan)
	wg2.Wait()

	for i, fan := range t.buyers {
		if i == 0 {
			fmt.Printf("%d -> %s - Parabéns foi o primeiro(a) a Comprar!!!!\n", i+1, fan.name)
			continue
		}
		if i == len(t.buyers)-1 {
			fmt.Printf("%d -> %s - Que sorte ficou com o último ingresso!\n", i+1, fan.name)
			continue
		}
		fmt.Printf("%d -> %s\n", i+1, fan.name)
	}

}
