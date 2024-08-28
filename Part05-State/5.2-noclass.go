package main

import (
	"fmt"
	"math"
)

type coordinate struct{
	d, m, s float64
	h rune
}

type lo2 struct {
	lat, long float64

}

type world struct {
	radius float64
}

func main() {

	lat := coordinate{4, 35, 22.2, 'S'}

	long := coordinate{137, 26, 30.12, 'E'}

	fmt.Println(lat.decimal(), long.decimal())
	
	cur := newLocation(coordinate{4, 35, 22.2, 'S'}, coordinate{137, 26, 30.12, 'E'})

	fmt.Println(cur)
	
	var mars = world{radius: 3389.5}

	spirit := lo2{-14.5684, 175.472636}
	
	oppo := lo2{-1.9462, 354.4734}

	dis := mars.distance(spirit, oppo)

	fmt.Printf("%.2f km\n", dis)

}

func (c coordinate) decimal() float64  {
	sign := 1.0
	switch c.h {
	case 's', 'w', 'S', 'W':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func newLocation(lat, long coordinate) lo2  {
	return lo2{lat.decimal(), long.decimal()}
}
func rad (deg float64) float64{
	return deg * math.Pi / 180
}

func (w world) distance (p1, p2 lo2) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2 + c1*c2*clong)

}
