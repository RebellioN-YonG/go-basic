package main

import (
	"fmt"
	// "math"
	"math/rand"
)

func main() {

	// third := 10.0 / 3
	// fmt.Printf("%v\n", third)
	// fmt.Printf("%f\n", third)
	// fmt.Printf("%.2f\n", third)
	// fmt.Printf("%.5f\n", third)
	// fmt.Printf("%5.3f\n", third)

	// HOMEWORK: PIGGY

	const (
		c1 = 0.05
		c2 = 0.10
		c3 = 0.20
	)

	total := 0.00

	for total < 20.00 {

		switch rand.Intn(3) {

		case 0:
			total += c1
		case 1:
			total += c2
		case 2:
			total += c3
		}

		fmt.Printf("%4.2f\n", total)

	}

}
