package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup

type Register struct {
	id     uint
	weight float64
}

func (r *Register) GenerateID() uint {
	r.id++
	return r.id
}

func (r Register) GenerateWeight() float64 {
	min := 0.5
	max := 20.0
	r.weight = min + rand.Float64()*(max-min)
	return r.weight
}

type Package struct {
	id     uint
	weight float64
}

func (p Package) NewPackage(r *Register) Package {
	p.id = r.GenerateID()
	p.weight = r.GenerateWeight()

	return p
}

type Painel struct {
	pk     []Package
	pac    []Package
	sedex  []Package
	painel sync.RWMutex
}

func (p *Painel) AddPortage(pk *Package) {
	p.painel.Lock()
	defer p.painel.Unlock()

	if pk.weight > 10.01 {
		p.pac = append(p.pac, *pk)
		return
	}

	p.sedex = append(p.sedex, *pk)
}

func (p *Painel) AddPackage(pk Package) {
	p.pk = append(p.pk, pk)
}

func (p *Painel) Total() string {
	tPac := len(p.pac)
	tSedex := len(p.sedex)

	return fmt.Sprintf("Total de Pac:%d | Total de Sedex:%d", tPac, tSedex)
}

func GeneratePK(p *Painel, r Register, pk Package) {
	for i := 1; i <= 100; i++ {
		newPK := pk.NewPackage(&r)
		p.AddPackage(newPK)
	}

}

func Workers(chanPK chan Package, p *Painel) {
	for cPK := range chanPK {
		p.AddPortage(&cPK)
	}
}

func main() {
	p := Painel{}
	r := Register{}
	pk := Package{}

	GeneratePK(&p, r, pk)

	chPK := make(chan Package)

	for i := 1; i <= 5; i++ {
		wg.Go(func() {
			Workers(chPK, &p)
		})
	}

	for _, pk := range p.pk {
		chPK <- pk
	}

	close(chPK)
	wg.Wait()

	fmt.Println(p.Total())

}
