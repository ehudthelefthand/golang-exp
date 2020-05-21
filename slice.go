package main

import "fmt"

func runSlice() {
	var a [10]int
	// printSlice(a)
	_ = a
	// var b []int
	// printSlice(b)

	c := a[2:4]
	printSlice(c)
	c[0] = 1
	fmt.Println(a)

	d := a[2:5]
	printSlice(d)

	e := make([]int, 5, 10)
	printSlice(e)

	vals := make([]int, 5)
	for i := 0; i < 5; i++ {
		vals = append(vals, i)
	}
	fmt.Println(vals)
}

func printSlice(s []int) {
	fmt.Printf("len = %d, cap = %d, s = %v, ptr = %p\n", len(s), cap(s), s, s)
}
