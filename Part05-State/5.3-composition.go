package main

import "fmt"

// type report struct {
// 	sol      int
// 	tmp      tmp
// 	location location
// }

type report struct {
	sol int
	tmp
	location
}

type tmp struct {
	high, low cel
}

type location struct {
	lat, long float64
}

type report2 struct {
	sol
	location
	tmp
}

type cel float64
type sol int

func main() {

	// bradbury := location{-4.5895, 137.4417}
	// t := tmp{-1.0, -18.0}
	rp := report{
		sol:      15,
		tmp:      tmp{-1.0, -18.0},
		location: location{-4.5895, 137.4417},
	}

	fmt.Printf("%+v\n", rp)

	fmt.Printf("%v\n", rp.tmp.high)



	rp2 := report2{sol: 15}

	fmt.Println(rp2.sol.days(1446))
	
	fmt.Println(rp2.days(1446))

}

func (t tmp) average() cel {
	return (t.high + t.low) / 2.0

}

func (s sol) days(s2 sol) int {
	days := int(s2 - s)
	if days < 0 {
		days = -days
	}
	return days
}

func (r report2) days(s2 sol) int {
	return r.sol.days(s2)
}

// func (l location) days (l2 location) int {

// 	return 5
// }
