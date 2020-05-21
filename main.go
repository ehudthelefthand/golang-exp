package main

import (
	"fmt"
)

func main() {
	score := 60

	if score >= 60 {
		fmt.Println("Grade D")
	} else {
		fmt.Println("Don't know yet")
	}

	switch score {
	case 60:
		fmt.Println("Grade D")
	default:
		fmt.Println("Don't know yet")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	j := 0
	for j < 15 {
		fmt.Println(j)
		j++
	}

	fmt.Printf("min is %d\n", min(1, 2))

	fruits := []string{"apple", "banana", "orange"}
	for k := 0; k < len(fruits); k++ {
		fmt.Println(fruits[k])
	}

	for _, b := range fruits {
		fmt.Println("b => ", b)
		// fmt.Println("b => ", b)
	}

	var scoreMap map[string]int = map[string]int{
		"pongneng": 10,
		"nut":      12,
		"phume":    9,
	}

	for k, v := range scoreMap {
		fmt.Println("k => ", k, "v =>", v)
	}

	fmt.Println("-------- add ")

	scoreMap["sub"] = 1000000
	for k, v := range scoreMap {
		fmt.Println("k => ", k, "v =>", v)
	}

	fmt.Println("-------- delete")

	delete(scoreMap, "sub")

	for k, v := range scoreMap {
		fmt.Println("k => ", k, "v =>", v)
	}

	fmt.Println("-------- read")
	_, ok := scoreMap["pongneng"]
	fmt.Println("Is pongneng in map?", ok)
	_, ok = scoreMap["sub"]
	fmt.Println("Is sub in map?", ok)

	fmt.Println("-------- check key")

	presentMap := map[string]bool{
		"sub": false,
	}
	v, ok := presentMap["sub"]
	fmt.Println("Is present? v", v)
	fmt.Println("Is present? ok", ok)

	fmt.Println("-------- loop")
	for k, v := range presentMap {
		fmt.Println("k => ", k, "v =>", v)
	}

	// fmt.Println("-------- struct ------")
	// var p product = product{
	// 	name:     "apple",
	// 	price:    8.0,
	// 	discount: 0,
	// }
	// fmt.Printf("p => %v", p)
	runstruct()
	fmt.Println("------ slice -------")
	runSlice()
	fmt.Println("------ duck -------")
	runDuck()
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
