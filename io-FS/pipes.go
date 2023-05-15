package interfaces

import (
	"fmt"
	"io"
	"os"
)

func PipeExample() error {
	// create a pipe
	reader, writer := io.Pipe()

	go func() {
		// this is a goroutine that writes to the 'writer'
		data := []byte("Hello, Pipe!")

		_, err := writer.Write(data)
		if err != nil {
			fmt.Println("Error writing to pipe: ", err)
		}
		writer.Close()
	}()

	//  read from the 'reader' to stdout
	if _, err := io.Copy(os.Stdout, reader); err != nil {
		fmt.Println("Error reading from pipe: ", err)
	}

	return nil
}
