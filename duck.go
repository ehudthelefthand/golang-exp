package main

import "fmt"

type duck struct {
	name string
}

func (d *duck) fly() {
	fmt.Printf("%s is a duck. so, it fly like duck\n", d.name)
}

type electronicDuck struct {
	name            string
	batteryIncluded bool
}

func (d *electronicDuck) fly() {
	fmt.Printf("%s is a electronic duck. so, can not fly\n", d.name)
}

type flyable interface {
	fly()
}

type dog struct {
}

func (d *dog) fly() {
	fmt.Println("what!?")
}

func runDuck() {
	d := duck{name: "John"}
	makeItFly(&d)

	ed := electronicDuck{name: "Sub"}
	makeItFly(&ed)

	makeItFly(&dog{})

}

func makeItFly(d flyable) {
	d.fly()
}
