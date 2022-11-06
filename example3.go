package main

import (
	"fmt"
	"json/files"
)

func Example3() {
	jsonString := files.ReadFile("examples/003_list_with_objects.json")
	fmt.Println(jsonString)
}
