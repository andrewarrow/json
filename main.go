package main

import (
	"math/rand"
	"os"
	"template/util"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) == 1 {
		PrintHelp()
		return
	}
	command := os.Args[1]
	argMap := util.ArgsToMap()

	if command == "example1" {
	} else if command == "example2" {
	} else if command == "help" {
		PrintHelp()
	}
}
