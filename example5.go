package main

import (
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
	fmt.Println("The length is:", len(yt.Items))
	for i, item := range yt.Items {
		fmt.Println(i, item.Snippet.Title)
	}
}
