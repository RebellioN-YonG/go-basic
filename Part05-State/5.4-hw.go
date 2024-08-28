package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type coordinate struct {
	d, m, s float64
	h       rune
}

func (c coordinate) String() string {
	return fmt.Sprintf("%vº%v'%.1f\" %c", c.d, c.m, c.s, c.h)
}

type location struct {
	Name string     `json:"name"`
	Lat  coordinate `json:"latitude"`
	Long coordinate `json:"longitude"`
}

func (l location) String() string {
	return fmt.Sprintf("%v %v", l.Lat, l.Long)
}

func main() {
	ely := location{
		Name: "Elysium",
		Lat:  coordinate{4.0, 30.0, 0.0, 'N'},
		Long: coordinate{135.0, 53.0, 0.0, 'E'},
	}
	fmt.Println("Elysium Planitia is at", ely)

	fmt.Println(ely.Lat.decimal(), ely.Long.decimal())

	bytes, err := json.MarshalIndent(ely, "", "  ")
	exit0nError(err)
	fmt.Println(string(bytes))

}

func exit0nError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
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

func (c coordinate) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Dec float64 `json:"decimal"`
		Dms string  `json:"dms"`
		D   float64 `json:"degree"`
		M   float64 `json:"minute"`
		S   float64 `json:"second"`
		H   string  `json:"hemisphere"`
	}{
		Dec: c.decimal(),
		Dms: c.String(),
		D:   c.d,
		M:   c.m,
		S:   c.s,
		H:   string(c.h),
	})
}
