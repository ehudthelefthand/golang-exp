package main

import (
	"fmt"
	"time"
)

type base struct {
	id        uint
	createdAt time.Time
	updatedAt time.Time
}

type product struct {
	base
	name     string
	price    float64
	discount float64
}

func runstruct() {
	fmt.Println("-------- struct ------")
	var p product = product{
		name:     "apple",
		price:    8.0,
		discount: 0,
	}
	printValue(p)
	fmt.Printf("after printValue(p) => %v\n", p)
	printPtr(&p)
	fmt.Printf("after printPtr(&p) => %v\n", p)

	p.printValue()
	fmt.Printf("after p.printValue() => %v\n", p)
	p.printPtr()
	fmt.Printf("after p.printPtr() => %v\n", p)

	fmt.Println("------- loop struct ------")

	var ps []product
	ps = append(ps, product{
		name:     "apple",
		price:    1.0,
		discount: 0,
	})
	ps = append(ps, product{
		name:     "coconut",
		price:    2.0,
		discount: 0.1,
	})

	for _, p := range ps {
		fmt.Printf("%v\n", p)
	}

	// var ps []int= []int{}
	var pMap map[int]product = map[int]product{}
	// var pMap map[int]product = make(map[int]product)
	pMap[1] = product{
		name:     "apple",
		price:    1.0,
		discount: 0,
	}
	pMap[2] = product{
		name:     "coconut",
		price:    1.2,
		discount: 0.2,
	}
	for k, v := range pMap {
		fmt.Printf("%d => %v\n", k, v)
	}

	s := new(product)
	fmt.Printf("s => %#v\n", s)

	t := product{
		base: base{
			id:        1,
			createdAt: time.Now(),
			updatedAt: time.Now(),
		},
	}
	fmt.Printf("t => %#v\n", t)
	fmt.Printf("t.name => %#v\n", t.name)
	fmt.Printf("t.id => %#v\n", t.id)
}

func printValue(p product) {
	p.name = "orange"
	fmt.Printf("p  => %v\n", p)
}

func printPtr(p *product) {
	p.name = "orange"
	fmt.Printf("p  => %v\n", p)
}

func (p product) printValue() {
	p.name = "coconut"
	fmt.Printf("p  => %v\n", p)
}

func (p *product) printPtr() {
	p.name = "coconut"
	fmt.Printf("p  => %v\n", p)
}
