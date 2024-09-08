package main

import "fmt"

type role struct {
	name string
	lefthand *item
}

type item struct {
	name string

}

func (r *role) pick(i *item) {
	if r == nil || i == nil {
		return
	}
	r.lefthand = i
	fmt.Printf("%v picked %v.\n",r.name, i.name)
}

func (r *role) give(t *role) {
	if r == nil || t == nil {
		return
	}
	if r.lefthand == nil {
		fmt.Printf("%v has no item to give.\n", r.name)
		return
	}

	if t.lefthand != nil {

		fmt.Printf("%v already has an item.\n", t.name)
		return
	}


	t.lefthand = r.lefthand
	r.lefthand = nil
	fmt.Printf("%v gave a %v to %v.\n", r.name, t.lefthand.name, t.name)
}

func (r *role) String() string {
	if r.lefthand == nil {
		return fmt.Sprintf("%v has no item.d", r.name)
	}
	return fmt.Sprintf("%v has a %v.", r.name, r.lefthand.name)
}

func main() {
	
	arthur := &role{name: "Arthur"}
	sword := &item{name: "sword"}

	arthur.pick(sword)

	knight := &role{name: "Knight"}
	arthur.give(knight)

	fmt.Println(arthur)
	fmt.Println(knight)

}