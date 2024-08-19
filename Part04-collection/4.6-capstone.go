package main

import (
	"fmt"
	"math/rand"

)

// capstone 4.6.1 : Conway's Game of Life

// A new universe

const (
	width  = 8
	height = 8
)

type universe[][] bool

func new() universe {
	u := make(universe, height)
	for i := range u {
		u[i] = make([]bool, width)
	}
	return u
}

func (n universe) Show() {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
				fmt.Printf("%v ", n[i][j])			
		}
		fmt.Println()
	}
}

func (n universe) Seed() {
	
	cnt := height * width / 4
	for i := 0; i <= cnt; i++ {
		x := rand.Intn(height) 
		y := rand.Intn(width) 
		n[x][y] = true
	}	
}

func (n universe) Alive(x, y int) bool {
		
	if y - 1 < 0  {
		y = height - 1
	}

	if y + 1 > height - 1 {
		y = 0
	}

	if x - 1 < 0 {
		x = width - 1		
	}

	if x + 1 > width - 1 {
		x = 0
	}		

	x := rand.Intn(height) 
	y := rand.Intn(width) 
	n[x][y] = true

	return n
}

func main() {

	n := new()
	n.Seed()
	n.Show()

}
