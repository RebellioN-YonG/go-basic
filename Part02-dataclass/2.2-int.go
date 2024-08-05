package main

import (

	"fmt"
	// "math"
	"math/rand"
)


func main()  {
	

	var r, g, b uint8 = 0, 141, 213
	
	fmt.Printf("%x %x %x\n", r, g, b)

	fmt.Printf("color: #%02x%02x%02x\n", r, g, b)

	var color uint8 = 3

	for color <= 10{
		fmt.Printf("%08b\n", color)
		color ++
	}

	// HOMEWORK: PIGGY 2

	const (
		c1 = 5
		c2 = 10
		c3 = 25
	)

	total := 0

	for total < 2000 {

		switch rand.Intn(3) {

		case 0:
			total += c1
		case 1:
			total += c2
		case 2:
			total += c3
		}

		dollar := total / 100
		cent := total % 100

		fmt.Printf("%d.%02d\n", dollar, cent)

	}

}



}