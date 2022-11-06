package main

import (
	"fmt"
	"json/files"
)

type SimpleList struct {
	Items []string
}

func Example2() {
	jsonString := files.ReadFile("examples/002_simple_list.json")
	fmt.Println(jsonString)

	sl := SimpleList{}
	fmt.Println("The simple list is:", sl.Items)
	fmt.Println("Thelength is:", len(sl.Items))
}
