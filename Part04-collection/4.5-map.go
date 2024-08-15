package main

import (
	"fmt"
	"math"
	"sort"
	// "math/rand"
)

func main() {

	// like dictionay in python , map has a similar format

	m := map[string]int{
		"a":    1,
		"b":    3,
		"moon": 0,
	}

	fmt.Println(m)

	c := m["C"]
	a := m["a"]
	moon := m["moon"]
	fmt.Println(c, a, moon)

	// "," and "ok":
	if moon, ok := m["moon"]; ok {
		fmt.Printf("moon is %v\n", moon)
	} else {
		fmt.Println("where is the 'moon'?")
	}

	tmp := map[float64]int{
		1.1: 1,
		2.2: 2,
		3.3: 3,
	}

	fmt.Println(tmp)

	//  Preallocating maps with make
	num := make(map[float64]int, 8)

	// maps that created with make's len is 0
	fmt.Println(len(num))

	// create a counter with map
	cnt := []float64{
		23.4, 34.9, -48.2, 21.3, 23.4, -43.1, 22.0, 22.0,
	}

	for _, t := range cnt {
		num[t]++
	}

	for t, fre := range num {
		fmt.Printf("%+.2f occurs %d  times\n", t, fre)
	}

	scores := []float64{
		60.5, 72.8, 88.6, 87.8, 50.5, 76.9, 70.6, 60.4, 59.2, 67.4, 60.5, 79.7,
	}

	grp := make(map[float64][]float64)

	for _, t := range scores {
		g := math.Trunc(t/10) * 10
		grp[g] = append(grp[g], t)
	}

	for g, scores := range grp {
		fmt.Printf("%v: %v\n", g, scores)
	}

	// using map as replacement of set
	set := make(map[float64]bool)

	for _, t := range scores {
		set[t] = true
	}

	if set[32.0] {

		fmt.Println("set member")
	}

	fmt.Println(set)

	// to sort the key value in set
	unique := make([]float64, 0, len(set))
	for t := range set {
		unique = append(unique, t)
	}

	sort.Float64s(unique)
	fmt.Println(unique)

}

// for i := 0 ; i < 10; i++{
// 	fmt.Printf("%.1f, ",(rand.Float64() + float64(rand.Intn(50)) + 50))
// }
