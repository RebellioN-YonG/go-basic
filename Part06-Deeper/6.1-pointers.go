package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	a := "wxcb"
	addr := &a

	fmt.Println(a, &a)
	fmt.Println(*addr, addr)
	// fmt.Printf("addr is a %T\n",addr)
	// fmt.Printf("a is a %T\n",a)

	var home *string
	// fmt.Printf("%T\n",home)
	cn := "CN"
	home = &cn
	fmt.Println(home, *home)

	var admin *string
	scolese := "CS"
	admin = &scolese
	fmt.Println(admin, *admin)

	bolden := "CB"
	admin = &bolden
	fmt.Println(admin, *admin)
	// admin points to bolden,
	// whatever bolden changes, admin(address of boldern) won't change
	bolden = "CFB"
	admin = &bolden
	fmt.Println(admin, *admin)

	// if *admin changes, that means bolden get changed
	// so admin or &bolden(address) won't change
	*admin = "MGCB"
	fmt.Println(&bolden, bolden)

	// if another var: major points to admin,
	// then major and admin both point to the same address
	major := admin
	*major = "MGCB"
	fmt.Println(major, *major)
	fmt.Println(admin == major)

	// if admin points to other var,
	// that means admin get changed
	// it no longer equals to major or bolden's address
	light := "RL"
	admin = &light
	fmt.Println(admin, *admin)
	fmt.Println(admin == major, admin == &bolden)

	// with admin pointed to light,
	// if a new var pointed to major
	// then the new var will have the same value as major
	// but their address are different
	nv := *major
	fmt.Println(&nv, nv)
	fmt.Println(major, *major)

	// change bolden's value to MGCB
	// nv's value will same as bolden's value
	// but their address are different
	bolden = "MGCB"
	fmt.Println(&bolden, bolden)
	fmt.Println(nv == bolden)
	fmt.Println(&bolden == &nv)

	// pointers that point to a struct

	timmy := &person{
		name: "Timothy",
		age:  10,
	}

	timmy.superpower = "flying"
	fmt.Printf("%+v\n", timmy)

	// pointer to an array or slice
	// the array automatically dereferences when slicing or indexing
	sps := &[3]string{"flight", "speed", "strength"}
	fmt.Println(sps[0])
	fmt.Println(sps[:2])

	// pointers as parameters

	birthday(timmy) // var timmy is a pointer
	fmt.Printf("%+v\n", timmy)

	rebel := &person{
		name:       "Rebellion",
		superpower: "revenge",
		age:        23,
	}
	fmt.Printf("%+v\n", rebel)
	birthday(rebel) // var rebel is a pointer
	fmt.Printf("%+v\n", rebel)

	fmt.Printf("%+v\n", rebel)
	rebel.bd()
	fmt.Printf("%+v\n", rebel)

	day := time.Now()
	tommorow := day.Add(24 * time.Hour)

	fmt.Println(day.Format(layout))
	fmt.Println(tommorow.Format(layout))

	// methed call without a pointer
	wennie := person{
		name: "Wennie",
		age:  70,
	}
	wennie.bd()
	fmt.Printf("%+v\n", wennie)

	player := character{
		name: "Player",
	}

	// &player.stats provides a interior pointer
	// which points to the structure
	fmt.Printf("%+v\n", player.stats)
	LvUp(&player.stats)
	fmt.Printf("%+v\n", player.stats)
	LvUp(&player.stats)
	fmt.Printf("%+v\n", player.stats)

	// using pointers to mutating arrays
	var board [8][8]rune
	reset(&board)
	fmt.Printf("%c\n", board[0][0])

	planets := []string{

		"Mercury", "Venus", "Earth",
		"Mars", "Jupiter", "Saturn",
		"Uranus", "Neptune", "Pluto",
	}

	reclassify(&planets)
	fmt.Println(planets)

	pl2 := reclassify2(&planets)
	fmt.Println(pl2)

	// both martian and pointer to martian
	// satisfy the talker interface
	shout(martian{})
	shout(&martian{})

	// if the methods'recievers are pointers,
	// then that will be different as shown below
	// &pew satisfies the talker interface while pew doesn't
	//  as it's a value, shout(pew) won't work
	pew := laser(2)
	shout(&pew)

}

// pointers and interfaces
type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type martian struct{}

func (m martian) talk() string {
	return "neck neck"
}

type laser int

func (l *laser) talk() string {
	return strings.Repeat("pew", int(*l))
}

// slices point at the arrays
func reclassify(planets *[]string) {
	*planets = (*planets)[0:8]
}

// this func returns a new slice
func reclassify2(planets *[]string) []string {
	pl2 := (*planets)[0:8]
	return pl2
}

func reset(board *[8][8]rune) {
	board[0][0] = 'r'

}

type stats struct {
	level    int
	endu, hp int
}

// address operator in go could get the
// address of the struct's field
func LvUp(s *stats) {
	s.level++
	s.endu = 42 + (14 * s.level)
	s.hp = 5 * s.endu
}

type character struct {
	name  string
	stats stats
}

const layout = "Mon. Jan 2, 2006"

type person struct {
	name, superpower string
	age              int
}

func birthday(p *person) {
	p.age++
}

// This func is a method of
// the person pointer reciever

func (p *person) bd() {
	p.age++
}
