package main

import (
	"fmt"
	"math"
)

type gps struct {
	location
	world
}

type location struct {
	lat, long float64
}
type world struct {
	radius float64
}

type rover struct {
	gps
}

func main() {

	var (
		mars      = world{3389.5}
		curiosity = rover{gps{location{-4.59, 137.44}, mars}}
	)
	now := gps{location{-4.5895, 137.4417}, mars}
	des := gps{location{4.5, 135.9}, mars}

	dist := curiosity.distance(now.location, des.location)
	curiosity.message(dist)

}

// func (g gps) distance

func (g gps) message(d float64) {
	fmt.Printf("%.2f km remaining to the destination.\n", d)
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func (g gps) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return g.world.radius * math.Acos(s1*s2+c1*c2*clong)

}
