package main

import (
	"fmt"
	"math/rand"
)

type kel float64

type sen func() kel

func rs() kel {
	return 0
}

func cali(s sen, offset kel) sen {

	return func() kel {

		return s() + offset

	}

}

func fs() kel {

	return kel(rand.Intn(151) + 150)

}

func main() {

	// 3.3 HW

	var offset kel = 5
	se := cali(fs, offset)

	for i := 0; i < 10; i ++{
		fmt.Println(se())
	}
	

}
