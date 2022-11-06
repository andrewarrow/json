package main

import (
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
		Example1()
	} else if command == "example2" {
	} else if command == "help" {
		PrintHelp()
	}
}
