package main

import (
	"fmt"
)

type test struct {
	a int
}

func main() {
	var pointer *test = nil
	fmt.Println(pointer.a)
}
