package main

import (
	"fmt"
	"log"

	helloExport "example.com/greetings"
)

func main() {
	
    log.SetPrefix("check if error is present: ")
    // log.SetFlags(0) disable timestamp, date & source file. 0 if you want to disable all of them


    // message, err := helloExport.Hello("Joe")
    messages, err := helloExport.HelloMultiple([]string{"Joe", "John", "Jane"})

    if (err != nil) {
        log.Fatal(err) // GO does not support one liner if statement
    }

    fmt.Println(messages)
}