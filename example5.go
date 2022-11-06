package main

import (
	"encoding/json"
	"fmt"
	"json/files"
)

type Youtube struct {
	Items []Item
}

type Item struct {
	Snippet Snippet
}

type Snippet struct {
	Title string
}

func Example5() {
	jsonString := files.ReadFile("examples/005_youtube.json")
	fmt.Println(jsonString)

	yt := Youtube{}
	json.Unmarshal([]byte(jsonString), &yt)
	fmt.Println("The length is:", len(yt.Items))
	for i, item := range yt.Items {
		fmt.Println(i, item.Snippet.Title)
		//fmt.Println(i, html.UnescapeString(item.Snippet.Title))
	}
}
