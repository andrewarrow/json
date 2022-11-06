package main

import (
	"fmt"
	"json/files"
)

type ObjectsInObjects struct {
	Items []Extra
}

type Extra struct {
	Message Message
}

func Example4() {
	jsonString := files.ReadFile("examples/004_objects_in_objects.json")
	fmt.Println(jsonString)

	oio := ObjectsInObjects{}
	fmt.Println("The list is:", oio.Items)
	fmt.Println("The length is:", len(oio.Items))
}
