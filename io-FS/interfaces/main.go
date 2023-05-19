package main

import (
	"bytes"
	"fmt"
	"interfaces"
)

func main() {

	// func NewReader(b []byte) *Reader
	// *type Reader defines read() and seek()
	in := bytes.NewReader([]byte("sample input text"))
	// type Buffer defines write()
	out := &bytes.Buffer{} 

	fmt.Println("[debug] calling interfaces.Copy \n[debug] first copy input to output variable and then copy to stdout")
	fmt.Print("[stdout] ")
	if err := interfaces.Copy(in,out); err != nil {
		panic(err)
	}

	fmt.Println("out variable : ", out.String())
	
	fmt.Println("[debug] calling interfaces.Pipe\n[debug] using io.Pipe()")
	fmt.Print("stdout on PipeExample = ")
	if err := interfaces.PipeExample(); err != nil {
		panic(err)
	}

}
