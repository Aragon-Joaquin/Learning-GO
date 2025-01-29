package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	fmt.Println(reverse.String("This string will be reversed"), reverse.Int(123))
	// reverse.Int() is not a native function from the example/hello/reverse package, the Go tutorial helps you implement it by modifying the reverse.go package
}