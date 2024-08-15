package main

import "fmt"

// capstone 4.6.1 : Conway's Game of Life

//

const (
	width  = 80
	height = 15
)

type universes [][]bool

func new() universes{
	make(universes[][], height, width)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			universes[i][j] = false
		}	
	}
}

func ()Show()  {
	fmt.Println(universes)
}

func main() {

	new()
	
	fmt.Println("")

}
