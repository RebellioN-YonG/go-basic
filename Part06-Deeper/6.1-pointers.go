package main

import (
	"fmt"
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
	type person struct {
		name, superpower string
		age              int
	}

	timmy := &person{
		name: "Timothy",
		age:  10,
	}

	timmy.superpower = "flying"
	fmt.Printf("%+v\n", timmy)

}
