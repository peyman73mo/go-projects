package main

import (
	"fmt"
)

type T1 struct {
	name string
	id   int
}

// method syntax :
// func (receiver) func_name (parameters) return_type
func (t1 T1) show() {
	fmt.Println("name :", t1.name)
	fmt.Println("id :", t1.id)
}
func show() {
	fmt.Println("this is just a function")
}
func (t1 *T1) changer() {
	t1.name = "blank"
	t1.id = -1
}

func main() {

	var a1 T1 = T1{
		name: "peyman",
		id:   1,
	}
	a2 := T1{
		"jack",
		2,
	}

	f2(&a1)
	fmt.Println(a1)

	f1(a2)
	// a2.id++
	a2.changer()
	a2.show()

	// show()

}

func f1(u T1) {
	u.id = 0
}
func f2(u *T1) {
	u.id = 0
}
