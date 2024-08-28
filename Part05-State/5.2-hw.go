package main

import (
	"fmt"
	"math"
)

type coordinate struct {
	d, m, s float64
	h       rune
}

type location struct {
	lat, long float64
}
type world struct {
	radius float64
}

func main() {

	// hw1 table 22.1 landing sites on mars'coor translated to location
	spirit := newLocation(coordinate{14, 34, 6.2, 'S'}, coordinate{175, 28, 21.5, 'E'})
	oppo := newLocation(coordinate{1, 56, 46.3, 'S'}, coordinate{354, 28, 24.2, 'E'})
	cur := newLocation(coordinate{4, 35, 22.2, 'S'}, coordinate{137, 26, 30.1, 'E'})
	insght := newLocation(coordinate{4, 30, 0.0, 'N'}, coordinate{135, 54, 0, 'E'})

	fmt.Println(spirit, "\n", oppo, "\n", cur, "\n", insght)
	fmt.Printf("curiosity's lat and long are %.2f, %.2f\n",cur.lat, cur.long)
	// hw2 table 22.1 distance between two locations
	var (
		// mecury = world{2439.7}
		// venus = world{6051.8}
		earth = world{6371.0}
		mars  = world{3389.5}
		// jupiter = world{69911}
		// saturn = world{58232}
		// uranus = world{25362}
		// neptune = world{24622}
		London     = newLocation(coordinate{51, 30, 0.0, 'N'}, coordinate{0, 8, 0.0, 'W'})
		Paris      = newLocation(coordinate{48, 50, 0.0, 'N'}, coordinate{2, 30, 0.0, 'E'})
		mountsharp = newLocation(coordinate{5, 4, 48, 'S'}, coordinate{137, 51, 0, 'E'})
		olymons    = newLocation(coordinate{18, 39, 0, 'N'}, coordinate{226, 12, 0, 'E'})
	)

	dist := []float64{
		mars.distance(spirit, oppo),
		mars.distance(spirit, cur),
		mars.distance(spirit, insght),
		mars.distance(oppo, cur),
		mars.distance(oppo, insght),
		mars.distance(cur, insght),
		earth.distance(London, Paris),
		mars.distance(mountsharp, olymons),
	}

	for i := 0; i < len(dist); i++ {
		fmt.Printf("%.2f km\n", dist[i])
	}
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 's', 'w', 'S', 'W':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func newLocation(lat, long coordinate) location {
	return location{lat.decimal(), long.decimal()}
}
func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)

}
