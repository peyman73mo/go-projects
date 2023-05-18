package interfaces

import (
	"fmt"
	"io"
	"log"
	"os"
)

func Copy(input io.ReadSeeker, output io.Writer) error {
	// writer variable that write to output and stdout
	writer := io.MultiWriter(output, os.Stdout)

	// direct copy from input to writer (destination, source)
	if _, err := io.Copy(writer, input); err != nil {
		log.Fatal(err)
		return err
	}

	input.Seek(0, 0)

	// copy from input to buffer and then to writer
	//  1 - create buffer
	buff := make([]byte, 64)
	//  2 - copy from input to buffer
	if _, err := io.CopyBuffer(writer, input, buff); err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("\nCopy is done")

	return nil
}
