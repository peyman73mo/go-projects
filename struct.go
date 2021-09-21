package main

import (
	"fmt"
)

type t1 struct {
	name string
	id   int
}

func main() {

	var a1 t1 = t1{
		name: "peyman",
		id:   1,
	}
	var a2 = t1{
		"jack",
		2,
	}

	f2(&a1)
	f1(a2)
	a2.id++
	fmt.Println(a1)
	fmt.Println(a2)
}

func f1(u t1) {
	u.id = 0
}
func f2(u *t1) {
	u.id = 0
}
