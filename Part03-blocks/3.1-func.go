package main

import (
	"fmt"
)

func K2C(k float64) float64 {
	k -= 273.15
	return k
}

func C2F(c float64) float64 {
	f := (c * 9.0 / 5.0) + 32
	return f
}

func K2F(k float64) float64 {
	f := ((k - 273.15) * 9.0 / 5.0) + 32

	return f
}

func main() {

	k1 := 303.0
	k0 := 0.0

	c1 := K2C(k1)

	f1 := C2F(c1)

	f2 := K2F(k0)

	fmt.Println(f1)
	fmt.Printf("%5.2f\n", f2)

}
