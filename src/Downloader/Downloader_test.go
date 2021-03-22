package Downloader

import (
	"fmt"
	"testing"
)

func TestDownloader(*testing.T) {
	links, items, err := Crawl("https://www.dandanzan.cc/", "")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("links", links)
	fmt.Println("items", items)
}
