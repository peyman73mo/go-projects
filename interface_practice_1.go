package main

import "fmt"

type F interface {
	F1()
}

type S struct {
	info string
}

func (s S) F1() {
	fmt.Println(s.info)
}

type T string

func (t T) print() {
	fmt.Println(t)
}

func main() {
	// var a F = S{"this is S type struct"}

	// fmt.Printf("type : %T\n", a)
	// // *** this is a error  ***
	// // fmt.Println(a.info)
	// a.F1()
	var a T = "Nothing"
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", T("alaki"))
	a.print()

}
