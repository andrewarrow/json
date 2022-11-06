package main

import (
	"fmt"
	"json/files"
)

func Example1() {
	jsonString := files.ReadFile("examples/001_say_hi.json")
	fmt.Println(jsonString)
}
