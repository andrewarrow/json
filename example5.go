package main

import (
	"encoding/json"
	"fmt"
	"html"
	"json/files"
)

type Youtube struct {
	Items []Item
}

type Item struct {
	Snippet Snippet
}

type Snippet struct {
	Title      string
	Thumbnails Thumbnails
}

type Thumbnails struct {
	Default Thumbnail
	Medium  Thumbnail
	High    Thumbnail
}

type Thumbnail struct {
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

func Example5() {
	jsonString := files.ReadFile("examples/005_youtube.json")

	yt := Youtube{}
	json.Unmarshal([]byte(jsonString), &yt)
	fmt.Println("The length is:", len(yt.Items))
	for i, item := range yt.Items {
		fmt.Println(i, html.UnescapeString(item.Snippet.Title))
		fmt.Print(item.Snippet.Thumbnails.Default.Width)
		fmt.Print("x")
		fmt.Println(item.Snippet.Thumbnails.Default.Height)
		fmt.Println("")
	}
}
