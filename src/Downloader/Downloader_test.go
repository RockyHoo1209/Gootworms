package Downloader

import (
	"fmt"
	"main/src/Utils/LinksUtil"
	"testing"
)

func TestDownloader(*testing.T) {
	resp, err := Download("https://www.bilibili.com/")
	if err != nil {
		fmt.Println(err)
	}
	links, err := LinksUtil.ExtractUrls(resp)
	if err != nil {
		fmt.Println(err.Error())
	}
	items, err := LinksUtil.ExtractItems(resp, "")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("links", links)
	fmt.Println("items", items)
}
