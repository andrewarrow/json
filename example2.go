package main

import (
	"fmt"
	"json/files"
)

func Example2() {
	jsonString := files.ReadFile("examples/002_simple_list.json")
	fmt.Println(jsonString)
}
