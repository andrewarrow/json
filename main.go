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
		Example2()
	} else if command == "example3" {
		Example3()
	} else if command == "example4" {
		Example4()
	} else if command == "help" {
		PrintHelp()
	}
}
