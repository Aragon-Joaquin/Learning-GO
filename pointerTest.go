package main

import (
	"fmt"
)

//! found this in https://stackoverflow.com/questions/38172661/what-is-the-meaning-of-and
func main() {
	var a = 5
	var pointer = &a // pointer holds variable a's memory address, for example, 0x7ffe5367e044
	fmt.Printf("1. Address of var a: %p\n", pointer)
	fmt.Printf("2. Value of var a: %v\n", *pointer)

	// Let's change a value (using the initial variable or the pointer)
	*pointer = 3 // using pointer

	fmt.Printf("3. Address of var a: %p\n", pointer)
	fmt.Printf("4. Value of var a: %v\n", *pointer)
}

/*
	Basically, the & operator returns where the variable is stored in memory.
	And the * operator resolves that address and gets the value stored.
*/