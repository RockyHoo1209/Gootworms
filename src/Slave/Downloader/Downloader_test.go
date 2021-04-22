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
	links, err := LinksUtil.ExtractUrls(resp,"((http[s]{0,1})://)(([a-zA-Z0-9\\._-]+\\.[a-zA-Z]{2,6})|([0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}))(:[0-9]{1,4})*(/[a-zA-Z0-9\\&%_\\./-~-]*)?")
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
