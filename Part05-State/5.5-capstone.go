package main

import (
	"fmt"
	"math/rand"
)

func main() {
	animals := []animal{
		dog{name: "Hond"},
		rabit{name: "Konijn"},
	}

	var hour, day int
	for {
		fmt.Printf("%v:00 ", hour)
		if hour <= sunrise || hour >= sunset {
			fmt.Printf("The animals are sleeping.\n")
		} else {
			i := rand.Intn(len(animals))
			action(animals[i])
		}

		hour++
		if hour >= 24 {
			hour = 0
			day++
			if day >= 3 {
				break
			}
		}
	}
}

type animal interface {
	move() string
	eat() string
}

const sunrise, sunset = 8, 18

type dog struct {
	name string
}

type rabit struct {
	name string
}

func (d dog) String() string {
	return d.name
}

func (r rabit) String() string {
	return r.name
}

func (d dog) move() string {
	switch rand.Intn(2) {
	case 0:
		return "running"
	case 1:
		return "playing"
	default:
		return "unknown"
	}
}

func (r rabit) move() string {
	switch rand.Intn(2) {
	case 0:
		return "jumping"
	case 1:
		return "walking"
	default:
		return "unknown"
	}
}

func (d dog) eat() string {
	switch rand.Intn(2) {
	case 0:
		return "dog food"
	case 1:
		return "meat"
	default:
		return "unknown"
	}
}

func (r rabit) eat() string {
	switch rand.Intn(3) {
	case 0:
		return "vegetables"
	case 1:
		return "fruits"
	case 2:
		return "grain"
	default:
		return "unknown"

	}
}

func action(a animal) {
	switch rand.Intn(2) {
	case 0:
		fmt.Printf("%v is %v.\n", a, a.move())
	default:
		fmt.Printf("%v is eating %v.\n", a, a.eat())
	}
}
