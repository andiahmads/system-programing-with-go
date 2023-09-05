package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("please speficy file")
		return
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer f.Close()

	var (
		b = make([]byte, 16) //create buffer
	)

	for n := 0; err == nil; {
		n, err = f.Read(b)
		if err == nil {
			fmt.Print(b[:n])
		}
	}

	if err != nil && err != io.EOF { //only print what's been read
		fmt.Println("\n\nError:", err)
	}

}
