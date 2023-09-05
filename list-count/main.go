package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 { // ensure path spesific
		fmt.Println("please spesific a path.")
		return
	}
	root, err := filepath.Abs(os.Args[1]) // get absolut path
	if err != nil {
		fmt.Println("cannot get absolute path:", err)
		return
	}

	fmt.Println("listing files in", root)

	var c struct {
		files int
		dirs  int
	}

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		// walk the tree to count files and folder
		if info.IsDir() {
			c.dirs++
		} else {
			c.files++
		}
		fmt.Println("-", path)
		return nil
	})

	fmt.Printf("Total: %d file in %d directories", c.files, c.dirs)
}
