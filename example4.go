package main

import (
	"encoding/json"
	"fmt"
	"json/files"
)

type ObjectsInObjects struct {
	Items []Item
}

type Item struct {
	Extra Message
}

func Example4() {
	jsonString := files.ReadFile("examples/004_objects_in_objects.json")
	fmt.Println(jsonString)

	oio := ObjectsInObjects{}
	json.Unmarshal([]byte(jsonString), &oio)
	fmt.Println("The list is:", oio.Items)
	fmt.Println("The length is:", len(oio.Items))
}
