/*
 * Filename: f:\work\DistributeScrawer\src\Utils\LinksUtil\LinksUtil.go
 * Path: f:\work\DistributeScrawer\src\Utils
 * Created Date: Wednesday, March 10th 2021, 8:26:20 pm
 * Author: Rocky Hoo
 *
 * Copyright (c) 2021 XiaoPeng Studio
 */

package LinksUtil

import (
	"fmt"
	"log"
	"main/src/Utils/ChromeDriverUtil"
	"regexp"
	"strings"
)

func Extract(url string) ([]string, string, error) {
	doc_html, err := ChromeDriverUtil.StartChrome(url)
	if err != nil {
		log.Println(err)
		return nil, "", err
	}

	doc_html = strings.Replace(doc_html, " ", "", -1)
	doc_html = strings.Replace(doc_html, "\n", "", -1)
	fmt.Println(doc_html)
	//匹配所有链接形式
	reg, err := regexp.Compile("<ahref=\"([^\"]+)\"[^>]*>[^<]+</a>")
	if err != nil {
		log.Printf("reg parse error:%s\n", err.Error())
		return nil, doc_html, err
	}
	links := reg.FindAllString(doc_html, -1)
	return links, doc_html, nil
}
