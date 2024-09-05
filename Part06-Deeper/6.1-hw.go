package main

import "fmt"

func main() {

	t := turtle{
		x: 0,
		y: 0,
	}
	// var t turtle
	fmt.Println(t)

	t.right()
	t.down()
	t.left()

	fmt.Println(t)

	t.up()
	t.up()
	t.left()
	fmt.Println(t)
}

type turtle struct {
	x, y int
}

func (t *turtle) right() {
	t.x++
}

func (t *turtle) down() {
	t.y++
}

func (t *turtle) up() {
	t.y--
}

func (t *turtle) left() {
	t.x--
}
