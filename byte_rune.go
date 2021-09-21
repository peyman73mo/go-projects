package main

import (
	"fmt"
	"reflect" // for printing type of variable
)

func main() {
	var chr1 = 'A' //rune type (by default)
	// byte type :
	chr2 := byte('a')
	// chr2 := []byte("a")

	fmt.Printf("type: %T value: %v\n", chr1, chr1)
	fmt.Println("type: ", reflect.TypeOf(chr2), "value: ", chr2)
	if chr2 == 97 {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	var chr3 = []byte("something")
	for _, i := range chr3 {
		fmt.Println(string(i))
	}

}
