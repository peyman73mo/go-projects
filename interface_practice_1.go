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

func main() {
	var a F = S{"this is S type struct"}
	fmt.Printf("type : %T\n", a)
	// *** this is a error  ***
	// fmt.Println(a.info)
	a.F1()

}
