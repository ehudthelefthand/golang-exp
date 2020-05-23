package main

import "fmt"

func main() {
	defer cleanup()
	defer first()
	defer second()
	third()
}

func first() {
	fmt.Println("first")
}

func second() {
	fmt.Println("second")
}

func third() {
	fmt.Println("third")
}

func cleanup() {
	fmt.Println("cleanup")
}
