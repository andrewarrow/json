package main

import (
	"fmt"
	"json/files"
)

func Example5() {
	jsonString := files.ReadFile("examples/005_youtube.json")
	fmt.Println(jsonString)
}
