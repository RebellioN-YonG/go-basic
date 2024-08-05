package main

import (
	"fmt"
	// "time"

	"math/rand"
)

func main() {

	s := fS
	fmt.Println(s())

	s = rS
	fmt.Println(s())

	mt(3, fS)

	sex()

	//declare Anonymous func in the main 1
	f := func() {
		fmt.Println("wxcb!")
	}

	f()

	//declare Anonymous func in the main 2
	//decalre and execute immediately 
	func ()  {
		fmt.Println("wxcb!!")
	}()

}

type cel float64

// type kel float64

func fS() cel {

	return cel(rand.Intn(151) + 150)

}

func rS() cel {

	return 0

}

// passing func to other func

func mt(n int, s func() cel) {

	for i := 0; i < n; i++ {

		k := s()
		fmt.Printf("%v K\n", k)
	}

}

func sex() {
	fmt.Println("I want 2 sex!")
}
