package main

import (
	"fmt"
	"math/big"
	// "math"
	// "math/rand"
)

func main() {

	// a := big.NewInt(198964)

	// How to trans int64 to big,int?
	// #1
	// distance := new(big.Int)
	// distance.SetString("2400000000000000000000",10)
	// fmt.Println(distance)
	// #2

	// HOMEWORK: LIGHTYEAR 2 KILOMETER

	lightspeed := big.NewInt(299792)

	distance1 := new(big.Int)
	distance1.SetString("236000000000000000", 10)

	seconds := new(big.Int)
	seconds.Div(distance1, lightspeed)

	fmt.Println(seconds)

	secondpd := big.NewInt(86400)

	days := new(big.Int)
	days.Div(seconds, secondpd)

	dayspy := big.NewInt(365)

	lightyear := new(big.Int)
	lightyear.Div(days, dayspy)

	fmt.Println(days)

	fmt.Println(lightyear)

