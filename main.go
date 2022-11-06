package main

import (
	"fmt"
	"json/files"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) == 1 {
		PrintHelp()
		return
	}
	command := os.Args[1]

	if command == "example1" {
		jsonString := files.ReadFile("examples/001_say_hi.json")
		fmt.Println(jsonString)
	} else if command == "example2" {
	} else if command == "help" {
		PrintHelp()
	}
}
