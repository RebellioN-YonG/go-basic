package main

import "fmt"

func main() {

	//use type to declare new type

	// just like define in C/Cpp
	type f64 float64

	var t f64 = 20

	t += 10

	fmt.Println(t)

	// HW : type c,f,k trans methods
	var k kel = 294.0
	var c cel
	var f fah
	var k1 kel

	c = k2c(k)
	c = k.cel()

	f = c.fah()

	k1 = f.kel()

	fmt.Println(c, f, k1)
}

type cel float64
type fah float64
type kel float64

func k2c(k kel) cel {

	return cel(k - 273.15)
}

func (k kel) cel() cel {

	return cel(k - 273.15)
}

func c2f(c cel) fah {

	return fah((c * 9.0 / 5.0) + 32)

}

func (c cel) fah() fah {

	return fah((c * 9.0 / 5.0) + 32)
}

func f2k(f fah) kel {

	return kel(((f - 32) * 5.0 / 9.0) + 273.15)
}

func (f fah) kel() kel {

	return kel(((f - 32) * 5.0 / 9.0) + 273.15)

}
