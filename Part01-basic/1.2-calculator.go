package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Print("Jag heter Rebel,")
	fmt.Print("Jag komer fran CSUST\n")
	fmt.Printf("jag heter %v \njag komer fran %v \n", "Rebellion", "CSUST")

	var num = rand.Intn(100) + 1
	fmt.Println(num)

	num = rand.Intn(100) + 1
	fmt.Println(num)

	var distance = 56000000
	fmt.Printf("The speed at least is %v\n",distance/(28*24))
}
