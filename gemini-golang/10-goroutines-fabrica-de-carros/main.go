package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var colors = [3]string{"red", "blue", "black"}

type Car struct {
	id    uint
	name  string
	color string
}

func CreateCar(name string, r *Report) {
	for i := 1; i <= 200; i++ {
		c := Car{id: uint(i), name: name}
		r.addCars(&c)
	}
}

type Report struct {
	red      []*Car
	blue     []*Car
	black    []*Car
	defect   []*Car
	total    []*Car
	painting sync.RWMutex
}

func (r *Report) addCars(car *Car) {
	r.total = append(r.total, car)
}

func (r *Report) paintedCar(car *Car) {
	r.painting.Lock()
	defer r.painting.Unlock()
	switch car.color {
	case "red":
		r.red = append(r.red, car)
	case "blue":
		r.blue = append(r.blue, car)
	case "black":
		r.black = append(r.black, car)
	default:
		r.defect = append(r.defect, car)
	}
}

func (r *Report) report() string {
	return fmt.Sprintf(
		"Total de Carros vermelho: %d\n"+
			"Total de Carros azul: %d\n"+
			"Total de Carros preto: %d\n"+
			"Total de Carros com defeito: %d\n"+
			"Total de carros: %d\n", len(r.red), len(r.blue), len(r.black), len(r.defect), len(r.total))
}

func randColor() string {
	i := rand.Intn(len(colors))

	return colors[i]
}

func robotWorkers(chRobot chan *Car, r *Report) {
	for car := range chRobot {
		color := randColor()
		time.Sleep(time.Duration(10) * time.Second)
		car.color = color
		r.paintedCar(car)
	}
}

func main() {
	chProductionLine := make(chan *Car)
	nameCar := "Celta"
	r := Report{}

	CreateCar(nameCar, &r)

	wg.Go(func() {
		time.Sleep(time.Duration(3) * time.Second)
		robotWorkers(chProductionLine, &r)
	})

	wg.Go(func() {
		time.Sleep(time.Duration(4) * time.Second)
		robotWorkers(chProductionLine, &r)
	})

	wg.Go(func() {
		time.Sleep(time.Duration(5) * time.Second)
		robotWorkers(chProductionLine, &r)
	})

	wg.Go(func() {
		time.Sleep(time.Duration(1) * time.Second)
		robotWorkers(chProductionLine, &r)
	})

	for _, car := range r.total {
		chProductionLine <- car
	}
	close(chProductionLine)
	wg.Wait()

	fmt.Println(r.report())

}
