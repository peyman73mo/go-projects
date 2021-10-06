package main

import (
	"fmt"
)

// ----------------------------------------------------------------
type F interface {
	F1()
	foo()
}

// ----------------------------------------------------------------
// S has to have both foo() and F1 for being F interface
type S struct {
	info string
}

func (s S) foo() {
}
func (s S) F1() {
	fmt.Println(s.info)
}

// ----------------------------------------------------------------
type T string

func (t T) print() {
	fmt.Println(t)
}

// ----------------------------------------------------------------
// type T1 = strings

func main() {
	// var a F = S{"this is S type struct"}
	// fmt.Printf("type : %T\n", a)
	// ----------------------------------------------------------------
	// // *** this is a error  ***
	// // fmt.Println(a.info)

	// a.F1()
	// ----------------------------------------------------------------
	a := T("Nothing")
	fmt.Printf("%T\n", a)
	a.print()
	// ----------------------------------------------------------------
	// fmt.Printf("%T\n", T("alaki"))

}
