package main

import "fmt"

func main() {
	var a, b = 0, 1
	f1(a, b)
	f2(b)
}

func f1(a, b int) {
	c := a + b
	f2(c)
}

func f2(c int) {
	fmt.Println(c)
}
