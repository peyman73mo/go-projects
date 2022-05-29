package main

import (
	"errors"
	"fmt"
)

func main() {

	err := f1(0)
	// fmt.Println(err)
	if err != nil {
		fmt.Println(err)
	}
}

func f1(a int) error {
	if a == 0 {
		return errors.New("* error! a == 0")
	} else {
		fmt.Println("a == ", a)
	}
	return nil
}
