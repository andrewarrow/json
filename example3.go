package main

import (
	"fmt"
	"json/files"
)

type ListWithObjects struct {
	Items []Message
}

func Example3() {
	jsonString := files.ReadFile("examples/003_list_with_objects.json")
	fmt.Println(jsonString)

	lwo := ListWithObjects{}
	fmt.Println("The list with objects is:", lwo.Items)
	fmt.Println("The length is:", len(lwo.Items))
}
