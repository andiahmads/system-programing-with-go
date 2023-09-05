package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("please specify a path")
		return
	}

	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(string(b))
}
