package main

import (
	"fmt"
	"json/files"
)

type Message struct {
	Message string
}

func Example1() {
	jsonString := files.ReadFile("examples/001_say_hi.json")
	fmt.Println(jsonString)

	m := Message{}
	fmt.Println("The message is:", m.Message)
	fmt.Println("The message length is:", len(m.Message))
}
