package main

import (
	"fmt"
	"strings"
	"time"
)

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type crater struct{}

type ship struct {
	laser
}

func main() {

	var t interface {
		talk() string
	}

	t = martian{}

	fmt.Println(t.talk())

	t = laser(5)

	fmt.Println(t.talk())

	shout(martian{})
	shout(laser(2))
	// shout(crater{}) 	//This will cause a compile-time error.

	s := ship{laser(3)}

	fmt.Println(s.talk())
	shout(s)

	day := time.Date(2012, 8, 6, 5, 17, 0, 0, time.UTC)
	fmt.Printf("%.1f Curiosity has landed\n", date(day))

	fmt.Printf("%.1f Curiosity has landed\n", stardate(day))

	s2 := sol(1422)

	fmt.Printf("%.1f Happy birthday\n", stardate(s2))

	cur := location{-4.5985, 137.4417}

	fmt.Println(cur)

}

type str interface {
	String() string
}

type location struct {
	lat, long float64
}

func (l location) String() string {
	return fmt.Sprintf("%v, %v", l.lat, l.long)
}

func date(t time.Time) float64 {
	day := float64(t.YearDay())
	h := float64(t.Hour()) / 24.0
	return 1000 + day + h
}

type dater interface {
	YearDay() int
	Hour() int
}

func stardate(t dater) float64 {
	day := float64(t.YearDay())
	h := float64(t.Hour()) / 24.0
	return 1000 + day + h
}

type sol int

func (s sol) YearDay() int {
	return int(s % 668)
}

func (s sol) Hour() int {
	return 0
}
